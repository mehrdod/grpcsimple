package app

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"github.com/mehrdod/grpcsimple/internal/config"
	grpcdelivery "github.com/mehrdod/grpcsimple/internal/delivery/grpc"
	delivery "github.com/mehrdod/grpcsimple/internal/delivery/http"
	"github.com/mehrdod/grpcsimple/internal/server"
	"github.com/mehrdod/grpcsimple/internal/service"
	"github.com/mehrdod/grpcsimple/pkg/cache"
	"github.com/mehrdod/grpcsimple/pkg/logger"
)

func Run(configPath string) {
	//Initialize configs
	cfg, err := config.Init(configPath, "dev")
	log := logger.NewLogrusLogger()
	if err != nil {
		log.Errorf("failed to init config %v", err)
		return
	}

	redis := cache.NewRedisCache(cfg.Redis.Host, cfg.Redis.Port)

	services := service.NewServices(redis, cfg.Redis.CacheLimit, log)

	handlers := delivery.NewHandler(services, log)
	grpcHandlers := grpcdelivery.NewHandler(services)

	//GRPC Server
	lis, err := net.Listen("tcp", cfg.GRPC.Port)
	if err != nil {
		log.Errorf("failed to listen: %v", err)
	}
	grpcSrv := grpc.NewServer()
	grpcdelivery.RegisterFibonacciServer(grpcSrv, grpcHandlers)

	go func() {
		if err := grpcSrv.Serve(lis); err != nil {
			log.Errorf("failed to serve: %v", err)
		}
	}()

	// HTTP Server
	srv := server.NewServer(cfg, handlers.Init())

	go func() {
		if err := srv.Start(); !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	log.Info("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Errorf("failed to stop server: %v", err)
	}

	grpcSrv.Stop()

	log.Info("Server stopped")
}
