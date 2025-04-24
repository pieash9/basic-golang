package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pieash9/basic-golang/students-api/internal/config"
	"github.com/pieash9/basic-golang/students-api/internal/http/handlers/student"
	"github.com/pieash9/basic-golang/students-api/internal/storage/sqlite"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// database setup
	storage, err := sqlite.New(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage initalized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/student", student.New(storage))

	// setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server started on", slog.String("address", cfg.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done

	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to Shutdown", slog.String("error", err.Error()))
	}

	slog.Info("Server shut down successfully!")
}
