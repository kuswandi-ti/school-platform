package app

import (
	"context"
	"log/slog"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"

	identityv1 "school-platform/packages/proto/gen/go/identity/v1"
	"school-platform/services/identity-service/internal/config"
	"school-platform/services/identity-service/internal/repository"
	"school-platform/services/identity-service/internal/token"
	grpctransport "school-platform/services/identity-service/internal/transport/grpc"
	"school-platform/services/identity-service/internal/usecase"
)

type App struct {
	config config.Config
	logger *slog.Logger
}

func New(cfg config.Config, logger *slog.Logger) *App {
	return &App{config: cfg, logger: logger}
}

func (a *App) Run(ctx context.Context) error {
	pool, err := pgxpool.New(ctx, a.config.DatabaseURL)
	if err != nil {
		return err
	}
	defer pool.Close()
	if err := pool.Ping(ctx); err != nil {
		return err
	}

	privateKey, err := token.LoadPrivateKey(a.config.JWTPrivateKeyPath)
	if err != nil {
		return err
	}
	tokenIssuer, err := token.NewIssuer(
		privateKey,
		a.config.JWTIssuer,
		a.config.JWTAudience,
		a.config.AccessTokenTTL,
		a.config.RefreshTokenTTL,
	)
	if err != nil {
		return err
	}

	users := repository.NewUserRepository(pool)
	sessions := repository.NewSessionRepository(pool)
	authorization := repository.NewAuthorizationRepository(pool)
	login, err := usecase.NewLogin(
		users,
		sessions,
		authorization,
		tokenIssuer,
	)
	if err != nil {
		return err
	}
	refresh := usecase.NewRefresh(users, authorization, sessions, tokenIssuer)
	logout := usecase.NewLogout(sessions, tokenIssuer)
	server := grpc.NewServer()
	identityv1.RegisterIdentityServiceServer(server, grpctransport.NewIdentityServer(login, refresh, logout))

	listener, err := net.Listen("tcp", a.config.GRPCAddr)
	if err != nil {
		return err
	}
	a.logger.Info("starting identity gRPC server", slog.String("grpc_addr", a.config.GRPCAddr))

	errCh := make(chan error, 1)
	go func() { errCh <- server.Serve(listener) }()
	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		server.GracefulStop()
		return nil
	}
}
