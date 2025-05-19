package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/atomhudson/go_lang/student-api/internal/config"
	"github.com/atomhudson/go_lang/student-api/internal/http/handlers/student"
	"github.com/atomhudson/go_lang/student-api/internal/storage/sqlite"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// database setup

	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatalf("failed to create storage: %v", err.Error())
	}

	slog.Info("Storage initialized successfully", slog.String("env", cfg.Env), slog.String("storage_path", cfg.StoragePath))

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.NewStudent(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetStudents(storage))
	router.HandleFunc("GET /api/students", student.GetStudentsList(storage))
	router.HandleFunc("DELETE /api/students/{id}", student.DeleteStudent(storage))
	router.HandleFunc("PUT /api/students/{id}", student.UpdateStudent(storage))

	// http server setup
	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	slog.Info("starting server...", slog.String("address", cfg.HTTPServer.Addr))
	fmt.Println("starting server on", cfg.HTTPServer.Addr)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatalf("failed to start server: %v", err.Error())
		}
	}()

	<-done // wait for signal

	slog.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server stopped successfully")
}
