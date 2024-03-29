package main

import (
	"net/http"
	"time"

	"github.com/br0-space/musicbot/container"
	"github.com/br0-space/musicbot/pkg/config"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/pflag"
)

const (
	readTimeout    = 15 * time.Second
	writeTimeout   = 15 * time.Second
	maxHeaderBytes = 4096
)

func main() {
	config.Init()
	pflag.Parse()

	logger := container.ProvideLogger()
	cfg := container.ProvideConfig()

	logger.Info("Starting HTTP server listening on", cfg.Server.ListenAddr)

	r := mux.NewRouter()
	r.HandleFunc("/webhook", container.ProvideTelegramWebhookHandler().ServeHTTP)
	r.Handle("/metrics", promhttp.Handler())
	r.NotFoundHandler = http.HandlerFunc(notFound)
	http.Handle("/", r)

	srv := &http.Server{ //nolint:exhaustruct
		Addr:           cfg.Server.ListenAddr,
		Handler:        r,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		IdleTimeout:    0,
		MaxHeaderBytes: maxHeaderBytes,
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}

func notFound(_ http.ResponseWriter, req *http.Request) {
	logger := container.ProvideLogger()

	logger.Debugf("%s %s %s from %s", req.Method, req.URL, req.Proto, req.RemoteAddr)
	logger.Error("not found")
}
