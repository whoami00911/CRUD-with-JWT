package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func StartServer(router *gin.Engine) (*http.Server, error) {
	s := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go s.ListenAndServe()
	return s, nil
}
