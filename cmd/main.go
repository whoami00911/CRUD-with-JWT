package main

import (
	//"net/http"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"webPractice1/internal/server"
	"webPractice1/internal/transport/handlers"
	"webPractice1/pkg/errorPrinter"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	handler := handlers.NewHandlerAssetsResponse()
	//http.HandleFunc("/Abuseip/", handler.TaskHandler)
	abuseipGroup := router.Group("/Abuseip")
	{
		abuseipGroup.POST("/", handler.CreateHandler)
		abuseipGroup.PUT("/", handler.CreateHandler)
		abuseipGroup.GET("/", handler.GetAllHandler)
		abuseipGroup.DELETE("/", handler.DeleteAllHandler)

		// Обработка маршрутов с IP
		abuseipGroup.GET("/:ip", handler.GetHandler)
		abuseipGroup.DELETE("/:ip", handler.DeleteHandler)
	}

	srv, err := server.StartServer(router)
	if err != nil {
		errorPrinter.PrintCallerFunctionName(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
