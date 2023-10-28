package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/harmoniemand/stopmotion/internal/configuration"
	"github.com/harmoniemand/stopmotion/internal/handlers"
	"github.com/harmoniemand/stopmotion/internal/middlewares"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not found."))
}

func main() {
	log.Info("Starting stopmotion API server")

	config := configuration.Config{
		Port:     8080,
		LogLevel: "debug",
		BasePath: "/api/v1",
	}

	if strings.ToLower(config.LogLevel) == "debug" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	router := mux.NewRouter()

	imagesHandler := handlers.NewImagesHandler()
	router.HandleFunc(fmt.Sprintf("%v/images", config.BasePath), imagesHandler.GetImages).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	n := negroni.New()
	n.Use(&middlewares.LoggerMiddleware{Config: config})
	n.UseHandler(router)

	log.Infof("starting server on port %v", config.Port)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", config.Port), n)
}
