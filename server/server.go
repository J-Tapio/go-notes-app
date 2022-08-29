package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go-notes-app/handler"
)

func Start() {
	server := &http.Server{
		Handler: handler.AppRouter,
		Addr:    ":" + os.Getenv("PORT"),
		// From mux docs; avoid Slowloris attacks by implementing timeouts.
		// Slowloris - partial HTTP requests.
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Second * 60,
	}

	// Gracefully shutdown server
	// Referencing https://pkg.go.dev/net/http#Server.Shutdown
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// Received an interrupt signal, shut down.
		if err := server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
