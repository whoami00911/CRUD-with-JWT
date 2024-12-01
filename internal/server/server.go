package server

import (
	"net/http"
	"webPractice1/pkg/errorPrinter"
)

func StartServer() error {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		errorPrinter.PrintCallerFunctionName(err)
	}
	return nil
}
