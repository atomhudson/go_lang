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
)

func main() {
	// load config
	cfg := config.MustLoad()

	// database setup
	// setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to the student api"))
	})
	// http server setup
	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	slog.Info("starting server...", slog.String("address", cfg.HTTPServer.Addr))
	fmt.Println("starting server on", cfg.HTTPServer.Addr)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, os.Kill)

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
