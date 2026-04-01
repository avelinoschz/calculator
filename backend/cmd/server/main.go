package main

import (
	"context"
	"errors"
	"log/slog"
	"math"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/avelinoschz/calculator/backend/internal/calculator"
	"github.com/avelinoschz/calculator/backend/internal/handler"
)

var version = "dev" // set via -ldflags at build time

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	slog.Info("starting", "version", version)

	min := parseEnvFloat("CALC_MIN", math.Inf(-1))
	max := parseEnvFloat("CALC_MAX", math.Inf(1))

	svc, err := calculator.NewService(min, max)
	if err != nil {
		slog.Error("invalid service configuration", "error", err)
		os.Exit(1)
	}

	if !math.IsInf(min, -1) || !math.IsInf(max, 1) {
		slog.Info("operand limits configured", "min", min, "max", max)
	}

	h := handler.New(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("POST /api/v1/calculations", h.Calculate)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		slog.Info("server starting", "addr", srv.Addr)
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server failed", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("shutdown error", "error", err)
		os.Exit(1)
	}

	slog.Info("server stopped")
}

func parseEnvFloat(key string, fallback float64) float64 {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		slog.Warn("invalid env var value, using default", "key", key, "value", val, "default", fallback)
		return fallback
	}
	return f
}
