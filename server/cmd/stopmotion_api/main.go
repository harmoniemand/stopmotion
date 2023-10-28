package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/harmoniemand/stopmotion/internal/configuration"
	"github.com/harmoniemand/stopmotion/internal/handlers"
	"github.com/harmoniemand/stopmotion/internal/middlewares"
	"github.com/harmoniemand/stopmotion/internal/stores"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not found."))
}

func options(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// w.Header().Set("Content-Type", "text")
	w.Write(nil)
}

var ErrEnvVarEmpty = errors.New("getenv: environment variable empty")

func getenvStr(key string, default_value string) string {
	v := os.Getenv(key)
	if v == "" {
		return default_value
	}
	return v
}
func getenvInt(key string, default_value int) int {
	s := getenvStr(key, "")
	v, err := strconv.Atoi(s)
	if err != nil {
		return default_value
	}
	return v
}

func main() {
	log.Info("Starting stopmotion API server")

	config := configuration.Config{
		Port:         getenvInt("PORT", 8080),
		LogLevel:     getenvStr("LOGLEVEL", "debug"),
		BasePath:     getenvStr("BASEPATH", "/api/v1"),
		DBConnection: getenvStr("MONGO_CONNECTION", ""),
		DatabaseName: getenvStr("DBNAME", "stopmotion"),
	}

	if strings.ToLower(config.LogLevel) == "debug" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	projectStore := stores.NewProjectStore(&config)

	router := mux.NewRouter()

	projectsHandler := handlers.NewProjectsHandler(projectStore)

	router.Path(fmt.Sprintf("%v/projects/{id}", config.BasePath)).HandlerFunc(options).Methods("OPTIONS")
	router.Path(fmt.Sprintf("%v/projects/{id}", config.BasePath)).HandlerFunc(projectsHandler.GetProject).Methods("GET")

	router.Path(fmt.Sprintf("%v/projects/{id}/image", config.BasePath)).HandlerFunc(options).Methods("OPTIONS")
	router.Path(fmt.Sprintf("%v/projects/{id}/image", config.BasePath)).HandlerFunc(projectsHandler.PostImageToProject).Methods("POST")

	router.Path(fmt.Sprintf("%v/projects", config.BasePath)).HandlerFunc(options).Methods("OPTIONS")
	router.Path(fmt.Sprintf("%v/projects", config.BasePath)).HandlerFunc(projectsHandler.GetProjects).Methods("GET")
	router.Path(fmt.Sprintf("%v/projects", config.BasePath)).HandlerFunc(projectsHandler.PostProject).Methods("POST")

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	n := negroni.New()
	n.Use(&middlewares.CorsMiddleware{})
	n.Use(&middlewares.LoggerMiddleware{Config: config})
	n.UseHandler(router)

	log.Infof("starting server on port %v", config.Port)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", config.Port), n)
}
