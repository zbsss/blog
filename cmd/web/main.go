package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type config struct {
	addr string
}

type app struct {
	config *config
	logger *slog.Logger
}

func main() {
	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":8080", "server address")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := &app{
		config: &cfg,
		logger: logger,
	}

	srv := http.Server{
		Addr:         cfg.addr,
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.logger.Info("starting server", slog.String("addr", cfg.addr))

	err := srv.ListenAndServe()
	app.logger.Error(err.Error())
	os.Exit(1)
}
