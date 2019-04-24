package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"wei/k8s-service-sample/handlers"
	"wei/k8s-service-sample/version"
)

func main() {
	log.Printf("Starting the service...\ncommit: %s\nrelease: %s\nbuild: %s\n", version.Commit, version.Release, version.BuildTime)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	router := handlers.Router(version.BuildTime, version.Commit, version.BuildTime)

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Print("The service is ready to listen and serve.")
	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	killSig := <-interrupt
	switch killSig {
	case os.Interrupt:
		log.Println("Got SIGINT...")
	case syscall.SIGTERM:
		log.Println("Got SIGTERM...")
	}

	log.Println("The service is shutting down...")
	server.Shutdown(context.Background())
	log.Print("Done")
}
