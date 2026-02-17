package main

// @title User Service API
// @version 1.0
// @description User service example API
// @BasePath /

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"user_service/internal/core"
)

func main() {
	server := core.GetServer()

	log.Printf("Server starting on: %s", server.Addr)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("Server is shutting down...")

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server shutdown successfully")
}
