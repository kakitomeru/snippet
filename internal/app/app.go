package app

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/kakitomeru/shared/env"
	"github.com/kakitomeru/shared/interceptor"
	"github.com/kakitomeru/shared/logger"
	"github.com/kakitomeru/shared/telemetry"
	"github.com/kakitomeru/snippet/internal/api"
	"github.com/kakitomeru/snippet/internal/config"
	"github.com/kakitomeru/snippet/internal/repository"
	"github.com/kakitomeru/snippet/internal/service"
	pb "github.com/kakitomeru/snippet/pkg/pb/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type App struct {
	db     *gorm.DB
	cfg    *config.Config
	port   string
	server *grpc.Server
}

func NewApp(db *gorm.DB, cfg *config.Config) *App {
	return &App{
		db:   db,
		cfg:  cfg,
		port: env.GetSnippetPort(),
	}
}

func (a *App) Start(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	shutdownTracer, err := telemetry.NewTracerProvider(ctx, a.cfg.Name, env.GetOtelCollector())
	if err != nil {
		logger.Error(ctx, "failed to create tracer provider", err)
		os.Exit(1)
	}
	defer func() {
		if err := shutdownTracer(ctx); err != nil {
			logger.Error(ctx, "failed to shutdown tracer provider", err)
			os.Exit(1)
		}
	}()

	shutdownMeter, err := telemetry.NewMeterProvider(ctx, a.cfg.Name, env.GetOtelCollector())
	if err != nil {
		logger.Error(ctx, "failed to create meter provider", err)
		os.Exit(1)
	}
	defer func() {
		if err := shutdownMeter(ctx); err != nil {
			logger.Error(ctx, "failed to shutdown meter provider", err)
			os.Exit(1)
		}
	}()
	// metric.Init()

	statsHandler, customInterceptor := telemetry.NewGRPCServerHandlers()
	authInterceptor := interceptor.AuthUnaryServerInterceptor()

	a.server = grpc.NewServer(
		statsHandler,
		grpc.ChainUnaryInterceptor(
			customInterceptor,
			authInterceptor,
		),
	)

	repo := repository.NewRepository(a.db)
	service := service.NewService(repo)

	snippetHandler := api.NewSnippetServiceHandler(service)
	pb.RegisterSnippetServiceServer(a.server, snippetHandler)

	reflection.Register(a.server)

	go func() {
		lis, err := net.Listen("tcp", ":"+a.port)
		if err != nil {
			logger.Error(ctx, "failed to listen", err)
			os.Exit(1)
		}

		log.Printf("Starting gRPC server on port %s", a.port)
		if err := a.server.Serve(lis); err != nil {
			logger.Error(ctx, "failed to serve", err)
			os.Exit(1)
		}
	}()

	a.GracefulShutdown(ctx)
}

func (a *App) GracefulShutdown(ctx context.Context) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		logger.Debug(ctx, "Shutdown requested via context")
	case <-quit:
		logger.Debug(ctx, "Shutdown requested via signal")
	}

	logger.Debug(ctx, "Shutting down gRPC server...")
	a.server.GracefulStop()
	logger.Info(ctx, "Server gracefully stopped")
}
