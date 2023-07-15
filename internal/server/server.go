package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

func InitServer(ctx context.Context, routes http.Handler) *http.Server {
	server := http.Server{
		Addr:         "0.0.0.0:8000",
		Handler:      routes,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  50 * time.Second,
	}

	// run the server in goroutine so that it don't block the main process
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Println(err.Error())
		}
	}()
	log.Print("server started in port :8000")
	return &server
}
