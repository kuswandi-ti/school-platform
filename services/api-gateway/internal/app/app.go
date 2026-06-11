package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"school-platform/services/api-gateway/internal/config"
	httptransport "school-platform/services/api-gateway/internal/transport/http"
)

type App struct {
	config config.Config
	logger *slog.Logger
}

func New(cfg config.Config, logger *slog.Logger) *App {
	return &App{
		config: cfg,
		logger: logger,
	}
}

func (a *App) Run() error {
	router := httptransport.NewRouter(a.config, a.logger)

	server := &http.Server{
		Addr:              a.config.HTTPAddr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	a.logger.Info(
		"starting api gateway",
		slog.String("service", a.config.ServiceName),
		slog.String("environment", a.config.AppEnv),
		slog.String("http_addr", a.config.HTTPAddr),
	)

	errCh := make(chan error, 1)
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
			return
		}
		errCh <- nil
	}()

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(stopCh)

	select {
	case err := <-errCh:
		return err
	case signal := <-stopCh:
		a.logger.Info("shutdown signal received", slog.String("signal", signal.String()))
	}

	ctx, cancel := context.WithTimeout(context.Background(), a.config.ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		return err
	}

	a.logger.Info("api gateway shut down gracefully")
	return nil
}
