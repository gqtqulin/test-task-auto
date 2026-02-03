package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
	"log"
	"os"
	"os/signal"
	"syscall"
	"test-task-auto/internal/config"
	"test-task-auto/internal/handler"
	"test-task-auto/internal/server"
	"test-task-auto/internal/service"
	"test-task-auto/internal/storage"
)

func main() {
	// TODO: перевести на viper
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: сделать без dsn
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
	log.Println("dsn:", dsn)

	connCfg, err := pgx.ParseDSN(dsn)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: можно на pgxpool
	conn, err := pgx.Connect(connCfg)
	if err != nil {
		log.Fatal(err)
	}

	store := storage.NewStorage(conn)
	s := service.NewService(store)
	h := handler.NewHandler(s)

	srv := server.Server{}
	go func() {
		if err := srv.Run(cfg.ServerPort, h.InitRoutes()); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal(err)
	}
}
