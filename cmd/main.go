package main

import (
	"net/http"
	"webPractice1/netHttp/handlers"
	"webPractice1/server"
)

func main() {
	handler := handlers.NewHandlerAssetsResponse()
	http.HandleFunc("/Abuseip/", handler.TaskHandler)
	server.StartServer()
}
