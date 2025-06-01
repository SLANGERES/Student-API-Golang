package main

import (
	"context"

	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/slangeres/Student-api-go/internal/config"
	"github.com/slangeres/Student-api-go/internal/http/handler/student"
)

func main() {
	cfg := config.MustDone()

	router := http.NewServeMux()

	router.HandleFunc("GET /api/", student.Home())
	router.HandleFunc("POST /api/students", student.PostStudent())

	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	// Channel to listen for interrupt signals
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine so it doesn't block
	go func() {
		slog.Info("Server is starting...", "addr", cfg.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed", "error", err)
			os.Exit(1)
		}
	}()

	// Block until we receive a termination signal
	<-done
	slog.Info("Shutting down the server...")

	// Give the server up to 5 seconds to shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shut down server gracefully", "error", err)
	} else {
		slog.Info("Server shut down gracefully")
	}
}
