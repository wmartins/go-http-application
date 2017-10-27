package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const version = "1.0.0"

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go HTTP Application - Version %s\n", version)
}

func main() {
	log.Println("Starting application")

	mux := http.NewServeMux()
	mux.HandleFunc("/version", versionHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	addr := net.JoinHostPort("", port)

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Attempting to get server up and runing on %s", server.Addr)
		log.Fatal(server.ListenAndServe())
	}()

	receivedSignal := <-signalChan

	log.Println(fmt.Sprintf("Captured %v. Exiting...", receivedSignal))
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	server.Shutdown(ctx)
	os.Exit(0)
}
