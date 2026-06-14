package app

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	schoolcorev1 "school-platform/packages/proto/gen/go/schoolcore/v1"
	"school-platform/services/school-core-service/internal/config"
	"school-platform/services/school-core-service/internal/repository"
	grpctransport "school-platform/services/school-core-service/internal/transport/grpc"
	"school-platform/services/school-core-service/internal/usecase"
)

type App struct {
	config config.Config
	logger *slog.Logger
}

func New(c config.Config, l *slog.Logger) *App { return &App{c, l} }
func (a *App) Run(ctx context.Context) error {
	pool, err := pgxpool.New(ctx, a.config.DatabaseURL)
	if err != nil {
		return err
	}
	defer pool.Close()
	if err = pool.Ping(ctx); err != nil {
		return err
	}
	svc := usecase.NewService(repository.NewSchoolRepository(pool))
	server := grpc.NewServer()
	schoolcorev1.RegisterSchoolCoreServiceServer(server, grpctransport.NewSchoolCoreServer(svc))
	listener, err := net.Listen("tcp", a.config.GRPCAddr)
	if err != nil {
		return err
	}
	a.logger.Info("starting school core gRPC server", "grpc_addr", a.config.GRPCAddr)
	ch := make(chan error, 1)
	go func() { ch <- server.Serve(listener) }()
	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		server.GracefulStop()
		return nil
	}
}
