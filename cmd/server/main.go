package main

import (
	"net/http"
	"time"

	"github.com/delivery-much/dm-go/logger"
	"github.com/g1stavo/taco-api/api"
	"github.com/g1stavo/taco-api/config"
)

func main() {
	http.Handle("/", api.NewRouter())
	startServer()
}

func startServer() {
	sv := &http.Server{
		Addr:    config.ServerPort,
		Handler: http.TimeoutHandler(http.DefaultServeMux, 10*time.Second, "Timeout"),
	}
	logger.Infof("Server started: %s", sv.Addr)
	if err := sv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Error: %s", err.Error())
	}
}
