package main

import (
	"fmt"
	"log"
	"net/http"

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
	fmt.Println("starting server on", cfg.HTTPServer.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start server: %v", err.Error())
	}

}
