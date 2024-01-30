package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/charmbracelet/log"
	"github.com/go-process-monitor/pkg/monitor"
)

const (
	ProcessListInterval = 10 * time.Millisecond
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	events := make(chan monitor.Event)
	go monitor.TrackProcesses(ctx, events)

	for event := range events {
		b, err := json.Marshal(event)
		if err != nil {
			log.Fatalf("Failed to marshal event: %v", err)
		}
		fmt.Println(string(b))
	}
}
