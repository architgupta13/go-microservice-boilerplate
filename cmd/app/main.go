package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"go-microservice-boilerplate/pkg/utl/config"
	stdlog "log"
	"os"
)

func main() {
	fmt.Println("Can y'all see this too?")

	// Set up logger interface
	var logger log.Logger
	logger = log.NewJSONLogger(log.NewSyncWriter(os.Stderr))
	stdlog.SetOutput(log.NewStdlibAdapter(logger))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "loc", log.DefaultCaller)

	_ = config.GetConfig()

}
