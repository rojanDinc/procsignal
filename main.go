package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var port = envOrDefault("PORT", "8888")

func main() {
	log.Println("starting application...")
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	go func() {
		s := <-signalChannel
		log.Printf("shutting down got signal: %s\n", s.String())
		os.Exit(0)
	}()

	http.HandleFunc("/ping", handlePing)
	log.Println("listening on port: ", port)
	log.Println(http.ListenAndServe("localhost:"+port, nil).Error())
}

func handlePing(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("pong"))
}

func envOrDefault(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return defaultVal
}
