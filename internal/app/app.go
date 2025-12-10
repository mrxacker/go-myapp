package app

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mrxacker/go-myapp/internal/config"
	"github.com/mrxacker/go-myapp/internal/repository/postgres"
	"github.com/mrxacker/go-myapp/internal/server"
	"github.com/mrxacker/go-myapp/internal/service"
	"github.com/mrxacker/go-myapp/pkg/logger"
	"github.com/mrxacker/go-myapp/pkg/logger/zap"
)

func Run(ctx context.Context) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	db, err := config.ConnectDB(cfg)
	if err != nil {
		return err
	}

	zapLogger, err := zap.NewZapLogger(cfg)
	if err != nil {
		return err
	}

	logger.Initialize(zapLogger)
	err = zapLogger.Sync()
	if err != nil {
		return fmt.Errorf("failed to sync logger: %w", err)
	}

	logger.Get().Info("Starting application")
	if err := runServers(ctx, cfg, db); err != nil {
		logger.Get().Error("Failed to run servers", logger.Error(err))
		return err
	}

	return nil
}

func runServers(ctx context.Context, cfg *config.Config, db *sql.DB) error {
	userRepository := postgres.NewPostgresUserRepository(db)
	userService := service.NewUserService(userRepository)

	httpServer := server.New(cfg, userService)
	go func() {
		logger.Get().Info("HTTP server starting", logger.Int("port", cfg.HTTPPort))
		if err := httpServer.Start(); err != nil {
			logger.Get().Fatal("HTTP server failed", logger.Error(err))
		}
	}()
	return nil
}
