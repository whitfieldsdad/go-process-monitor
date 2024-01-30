package monitor

import (
	"context"
	"time"

	"github.com/charmbracelet/log"
	lru "github.com/hashicorp/golang-lru"
	ps "github.com/shirou/gopsutil/v3/process"
)

const (
	ProcessListInterval = 10 * time.Millisecond
)

func TrackProcesses(ctx context.Context, events chan Event) {
	seen, err := lru.New(10000)
	if err != nil {
		log.Fatalf("Failed to create LRU cache: %v", err)
		return
	}

	processes, err := ps.Processes()
	if err != nil {
		log.Fatalf("Failed to list processes: %v", err)
		return
	}
	for _, p := range processes {
		ppid, err := p.PpidWithContext(ctx)
		if err != nil {
			continue
		}
		h := calculateProcessId(p.Pid, ppid)
		seen.Add(h, nil)
	}

	ids, err := ListProcessIdentities()
	if err != nil {
		log.Fatalf("Failed to list processes: %v", err)
		return
	}
	for _, id := range ids {
		seen.Add(id.Hash(), id)
	}

	ticker := time.NewTicker(ProcessListInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			processes, err := ps.Processes()
			if err != nil {
				log.Fatalf("Failed to list processes: %v", err)
				continue
			}
			for _, p := range processes {
				ppid, err := p.PpidWithContext(ctx)
				if err != nil {
					continue
				}
				h := calculateProcessId(p.Pid, ppid)
				_, ok := seen.Get(h)
				if !ok {
					seen.Add(h, nil)

					process, err := GetProcess(p.Pid, nil)
					if err != nil {
						log.Errorf("Failed to get process: %v", err)
						continue
					}
					log.Debugf("New process detected (PID: %d, PPID: %d, name: %s)", process.PID, process.PPID, process.Name)
					events <- NewProcessStartEvent(*process)
				}
			}

		case <-ctx.Done():
			return
		}
	}
}
