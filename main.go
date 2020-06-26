package main

import (
	"context"
	"log"
	"mongodb-test/controllers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type RoutesConfig struct {
	Path        string
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
	Method      []string
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime)

	router := mux.NewRouter()

	routes := []RoutesConfig{{
		Path:        "/login",
		HandlerFunc: controllers.LogiController,
		Method:      []string{http.MethodGet},
	}}

	for _, r := range routes {
		log.Printf("mounting %s - %s", r.Path, r.Method)
		go router.HandleFunc(
			r.Path,
			r.HandlerFunc,
		).Methods(r.Method...)
	}

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Println("starting the application...")
		log.Printf("server running: %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalln(err.Error())
		}
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	<-exit

	ctx, cancel := context.WithTimeout(
		context.Background(),
		15*time.Second)
	defer cancel()

	server.Shutdown(ctx)
	log.Println("shuting down...")
	os.Exit(0)
}
