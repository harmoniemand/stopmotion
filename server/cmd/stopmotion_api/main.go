package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/harmoniemand/stopmotion/internal/configuration"
	"github.com/harmoniemand/stopmotion/internal/handlers"
	"github.com/harmoniemand/stopmotion/internal/middlewares"
	"github.com/harmoniemand/stopmotion/internal/stores"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func notFoundHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte("Not found."))
	if err != nil {
		log.Errorf("Error writing response: %v", err)
	}
}

func options(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	// w.Header().Set("Content-Type", "text")
	_, err := w.Write(nil)
	if err != nil {
		log.Errorf("Error writing response: %v", err)
	}
}

var ErrEnvVarEmpty = errors.New("getenv: environment variable empty")

func getenvStr(key string, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}
func getenvInt(key string, defaultValue int) int {
	s := getenvStr(key, "")
	v, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}
	return v
}

func main() {
	err := server()

	if err != nil {
		log.Errorf("Error starting server: %v", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func server() error {
	log.Info("Starting stopmotion API server")

	config := configuration.Config{
		Port:         getenvInt("PORT", 8080), //nolint:gomnd // 8080 is the default port
		LogLevel:     getenvStr("LOGLEVEL", "debug"),
		BasePath:     getenvStr("BASEPATH", "/api/v1"),
		DBConnection: getenvStr("MONGO_CONNECTION", "mongodb://root:example@localhost:27017"),
		DatabaseName: getenvStr("DBNAME", "stopmotion"),
	}

	log.Debugf("Config: %v", config)

	if strings.ToLower(config.LogLevel) == "debug" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	projectStore, err := stores.NewProjectStore(&config)
	if err != nil {
		return err
	}

	router := mux.NewRouter()

	projectsHandler := handlers.NewProjectsHandler(projectStore)

	router.Path(fmt.Sprintf("%v/projects/{id}", config.BasePath)).HandlerFunc(options).Methods("OPTIONS")
	router.Path(fmt.Sprintf("%v/projects/{id}", config.BasePath)).HandlerFunc(projectsHandler.GetProject).Methods("GET")
	router.Path(fmt.Sprintf("%v/projects", config.BasePath)).HandlerFunc(options).Methods("OPTIONS")
	router.Path(fmt.Sprintf("%v/projects", config.BasePath)).HandlerFunc(projectsHandler.GetProjects).Methods("GET")
	router.Path(fmt.Sprintf("%v/projects", config.BasePath)).HandlerFunc(projectsHandler.PostProject).Methods("POST")

	imageHandler := handlers.NewImagesHandler(&config)
	router.Path(fmt.Sprintf("%v/projects/{project_id}/images", config.BasePath)).HandlerFunc(options).Methods("OPTIONS")
	router.Path(fmt.Sprintf("%v/projects/{project_id}/images", config.BasePath)).HandlerFunc(imageHandler.PostImage).Methods("POST")
	router.Path(fmt.Sprintf("%v/projects/{project_id}/images/{img_id}", config.BasePath)).HandlerFunc(options).Methods("OPTIONS")
	router.Path(fmt.Sprintf("%v/projects/{project_id}/images/{img_id}", config.BasePath)).HandlerFunc(imageHandler.GetImageAsFile).Methods("GET")
	router.Path(fmt.Sprintf("%v/projects/{project_id}/images/{img_id}", config.BasePath)).HandlerFunc(imageHandler.DeleteImage).Methods("DELETE")

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	n := negroni.New()
	n.Use(&middlewares.CorsMiddleware{})
	n.Use(&middlewares.LoggerMiddleware{Config: config})
	n.UseHandler(router)

	server := &http.Server{
		Addr:              fmt.Sprintf("0.0.0.0:%v", config.Port),
		ReadHeaderTimeout: 3 * time.Second, //nolint:gomnd // 3 seconds is a reasonable timeout
		Handler:           n,
	}

	log.Infof("starting server on port %v", config.Port)
	err = server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
