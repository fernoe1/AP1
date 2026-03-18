package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fernoe1/Assignment2_TemirlanSapargali/internal/model"
	"github.com/fernoe1/Assignment2_TemirlanSapargali/internal/server"
	"github.com/fernoe1/Assignment2_TemirlanSapargali/internal/store"
	"github.com/fernoe1/Assignment2_TemirlanSapargali/internal/worker"
)

func main() {
	str := store.New[string, string]()
	stats := model.NewStats()
	stopCh := make(chan struct{})
	srv := server.New(":8080", str, stats)

	go worker.StartStatsWorker(stats, stopCh)

	go func() {
		fmt.Println("Server listening on :8080")
		err := srv.ListenAndServe()
		if err != nil {
			fmt.Println("Error starting server:", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	<-sigCh
	fmt.Println("Shutdown signal received")

	close(stopCh)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Error shutting down:", err)
	}

	fmt.Println("Server gracefully shutdown")
}
