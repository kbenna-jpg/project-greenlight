package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Declares application version number
const version = "1.0.0"

// Holds all configuration settings for application
type config struct {
	port int
	env  string
}

// Application struct to hold dependencies for HTTP handlers, helpers
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	// read value of port and env command0line flags into config struct
	/// default to port 4000 and env development if no flags provided
	flag.IntVar(&cfg.port, "port", 8000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// initialize logger which writes messages to standard out stream, prefixed with date and time
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Declare instance of application struct containing config struct and logger
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Use httprputer instance returned by app.rputes() as server handler
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}


	// start HTTP server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
