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

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	identityv1 "school-platform/packages/proto/gen/go/identity/v1"

	"school-platform/services/api-gateway/internal/config"
	"school-platform/services/api-gateway/internal/middleware"
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
	identityConnection, err := grpc.NewClient(
		a.config.GRPCTargets.Identity,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}
	defer identityConnection.Close()

	publicKey, err := middleware.LoadPublicKey(a.config.JWTPublicKeyPath)
	if err != nil {
		return err
	}
	authValidator, err := middleware.NewJWTValidator(publicKey, a.config.JWTIssuer, a.config.JWTAudience)
	if err != nil {
		return err
	}

	router := httptransport.NewRouter(a.config, a.logger, identityv1.NewIdentityServiceClient(identityConnection), authValidator)

	server := &http.Server{
		Addr:              a.config.HTTPAddr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	a.logger.Info(
		"starting api gateway",
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
