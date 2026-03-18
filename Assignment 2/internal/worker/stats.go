package worker

import (
	"fmt"
	"time"

	"github.com/fernoe1/Assignment2_TemirlanSapargali/internal/model"
)

func StartStatsWorker(stats *model.Stats, stopCh <-chan struct{}) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			snapshot := stats.Snapshot()
			fmt.Print(snapshot, "\n")
		case <-stopCh:
			fmt.Println("Stats worker shutdown")
			return
		}
	}
}
