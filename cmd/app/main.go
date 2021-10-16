package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"go-microservice-boilerplate/pkg/app/health"
	"go-microservice-boilerplate/pkg/utl/config"
	stdlog "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Can y'all see this too?")

	// Set up logger interface
	var logger log.Logger
	logger = log.NewJSONLogger(log.NewSyncWriter(os.Stderr))
	stdlog.SetOutput(log.NewStdlibAdapter(logger))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "loc", log.DefaultCaller)

	configs := config.GetConfig()

	// Register sub-routers
	var httpLogger = log.With(logger, "component", "http")
	r := mux.NewRouter().StrictSlash(false)
	r.PathPrefix("/api").Handler(health.MakeHandler(httpLogger))

	shutdown := make(chan error)

	// Start an HTTP server
	var addr = ":" + configs.HttpServer.Port
	go func() {
		logger.Log("transport", "http", "address", addr, "msg", "listening on "+configs.HttpServer.Port)
		shutdown <- http.ListenAndServe(addr, r)
	}()

	// Gracefully terminate HTTP server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	logger.Log("terminated", <-c)
}
