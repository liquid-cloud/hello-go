package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func healthCheck(w http.ResponseWriter, req *http.Request) {
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func cleanup() {
	fmt.Println("Stopping container")
}

func setupMux(port int) error {
	allEndpoints := http.HandlerFunc(helloServer)
	healthcheckEndpoint := http.HandlerFunc(healthCheck)

	mux := http.NewServeMux()
	mux.Handle("/healthcheck", healthcheckEndpoint)
	mux.Handle("/", allEndpoints)

	fmt.Println(fmt.Sprintf("Listening on port %v...", port))
	return http.ListenAndServe(fmt.Sprintf(":%v", port), mux)
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(0)
	}()

	setupMux(8080)
}
