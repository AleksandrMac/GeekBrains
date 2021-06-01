package main

import (
	"net/http"
	"time"

	. "github.com/AleksandrMac/GeekBrains/Go/BackLevel1/http-server/pkg/handlers"
)

func main() {
	uploadHandler := &UploadHandler{
		UploadDir: "../../data",
	}
	http.Handle("/upload", uploadHandler)
	dirToServe := http.Dir(uploadHandler.UploadDir)

	fs := &http.Server{
		Addr:         ":8080",
		Handler:      http.FileServer(dirToServe),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fs.ListenAndServe()
}
