package main

import (
	"context"
	"fmt"
	"github.com/gqtqulin/test-task-auto/internal/config"
	"github.com/gqtqulin/test-task-auto/internal/handler"
	"github.com/gqtqulin/test-task-auto/internal/server"
	"github.com/gqtqulin/test-task-auto/internal/service"
	"github.com/gqtqulin/test-task-auto/internal/storage"
	"github.com/jackc/pgx"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// TODO: расширить уровни логирования
	log := initLogger()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Info("init config err", "error", err)
		os.Exit(1)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	connCfg, err := pgx.ParseDSN(dsn)
	if err != nil {
		log.Info("failed to parse dsn")
		os.Exit(1)
	}

	// TODO: при многопоточном соединении переписать на pgxpool
	conn, err := pgx.Connect(connCfg)
	if err != nil {
		log.Info("failed to connect to database")
		os.Exit(1)
	}

	carStorage := storage.NewCarStorage(conn)

	carService := service.NewCarService(carStorage)

	hand := handler.NewHandler(carService, log)

	srv := server.Server{}
	go func() {
		if err := srv.Run(cfg.ServerPort, hand.InitRoutes()); err != nil {
			log.Error("failed to start server", "error", err)
		}
	}()

	log.Info("server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Info("failed to shutdown server", "error", err)
		os.Exit(1)
	}

	if err := conn.Close(); err != nil {
		log.Info("failed to close connection", "error", err)
		os.Exit(1)
	}
}

func initLogger() *slog.Logger {
	slogHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	})

	return slog.New(slogHandler)
}
