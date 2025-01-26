package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "webPractice1/docs"
	"webPractice1/internal/repository"
	"webPractice1/internal/server"
	"webPractice1/internal/service"
	"webPractice1/internal/transport/handlers"
	"webPractice1/pkg/hasher"
	"webPractice1/pkg/logger"

	"github.com/spf13/viper"
)

// @title           rest with swagger
// @version         1.0
// @description     project4

// @host      localhost:8080
// @BasePath  /
func init() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.ReadInConfig()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка при чтении конфигурации: %v", err)
	}
}
func main() { ///////изменить логику обработки ошибок, добавить error в методы БД, чтобы в хендлерах выдавать 500
	logger := logger.GetLogger()
	hash := hasher.NewHashInit(viper.GetString("hashphrase"))
	repo := repository.NewRepository(repository.PostgresqlConnect(), logger)
	service := service.NewService(repo, hash, logger)
	handler := handlers.NewHandlerAssetsResponse(logger, service)

	srv := new(server.Server)
	go func() {
		err := srv.StartServer(handler.InitRoutes(), viper.GetString("server.port"))
		if err != nil {
			logger.Error(fmt.Sprintf("Server dont start: %s", err))
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("Shutdown error: %s", err))
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 1 seconds.")
	}
	log.Println("Server exiting")
}
