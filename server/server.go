package server

import (
	"net/http"
	"webPractice1/cmd/errorPrinter"
)

func StartServer() error {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		errorPrinter.PrintCallerFunctionName(err)
	}
	return nil
}
