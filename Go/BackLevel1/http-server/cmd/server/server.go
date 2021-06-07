package main

import (
	"log"
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
	downloadHandler := &DownloadHandler{
		Root: dirToServe,
	}
	http.Handle("/download/", downloadHandler)

	fs := &http.Server{
		Addr: ":8080",
		// Handler:      http.FileServer(dirToServe),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	if err := fs.ListenAndServe(); err != nil {
		log.Panic(err)
	}

}
