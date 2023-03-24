package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/kevinsudut/tech-curriculum-workshops/api/routes"
	"github.com/kevinsudut/tech-curriculum-workshops/config"
)

func main() {
	addr := "0.0.0.0:8000"
	if value, ok := os.LookupEnv("BIND_ADDR"); ok {
		addr = value
	}

	muxRouter := mux.NewRouter()

	err := routes.Init(muxRouter)
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: config.HTTP_WRITE_TIMEOUT,
		ReadTimeout:  config.HTTP_READ_TIMEOUT,
		IdleTimeout:  config.HTTP_IDLE_TIMEOUT,
		Handler:      muxRouter,
	}

	fmt.Println("Service started on", addr)

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("shutting down")

	os.Exit(0)
}
