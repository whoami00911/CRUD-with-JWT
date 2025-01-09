package main

import (
	//"net/http"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "webPractice1/docs"
	"webPractice1/internal/server"
	"webPractice1/internal/transport/handlers"
	"webPractice1/pkg/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           rest with swagger
// @version         1.0
// @description     project4

// @host      localhost:8080
// @BasePath  /

func main() {
	logger := logger.GetLogger()
	router := gin.Default()
	handler := handlers.NewHandlerAssetsResponse(logger)
	//http.HandleFunc("/Abuseip/", handler.TaskHandler)
	abuseipGroup := router.Group("/Abuseip")
	{
		abuseipGroup.POST("/", handler.CreateHandler)
		abuseipGroup.PUT("/", handler.UpdateHandler)
		abuseipGroup.GET("/", handler.GetAllHandler)
		abuseipGroup.DELETE("/", handler.DeleteAllHandler)

		// Обработка маршрутов с IP
		abuseipGroup.GET("/:ip", handler.GetHandler)
		abuseipGroup.DELETE("/:ip", handler.DeleteHandler)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	srv, err := server.StartServer(router)
	if err != nil {
		logger.Error(fmt.Sprintf("Server dont start: %s", err))
	}

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
