package main

import (
	"net/http"
	"webPractice1/internal/netHttp/handlers"
	"webPractice1/internal/server"
)

func main() {
	handler := handlers.NewHandlerAssetsResponse()
	http.HandleFunc("/Abuseip/", handler.TaskHandler)
	server.StartServer()
}
