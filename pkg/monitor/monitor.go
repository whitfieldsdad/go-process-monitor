package monitor

import (
	"context"
	"time"

	"github.com/charmbracelet/log"
	lru "github.com/hashicorp/golang-lru"
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
	listOpts := &ProcessOptions{
		IncludeHashes: false,
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
			ids, err := ListProcessIdentities()
			if err != nil {
				log.Fatalf("Failed to list processes: %v", err)
				return
			}
			for _, id := range ids {
				h := id.Hash()
				_, ok := seen.Get(h)
				if !ok {
					seen.Add(h, id)

					var process *Process
					process, err = GetProcess(id.PID, listOpts)
					if err != nil {
						log.Warnf("A new process was detected, but we weren't fast enough to get its details: %v (PID: %d, PPID: %d)", err, id.PID, id.PPID)
						process = &Process{
							PID:  id.PID,
							PPID: id.PPID,
						}
					}
					details := ProcessStartEventData{
						Process: *process,
					}
					events <- NewEvent(ObjectTypeProcess, EventTypeStarted, details)
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
