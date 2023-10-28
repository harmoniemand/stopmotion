package handlers

import (
	"net/http"

	"github.com/harmoniemand/stopmotion/internal/configuration"
)

type ImagesHandler struct {
	Config *configuration.Config
}

func NewImagesHandler() *ImagesHandler {
	return &ImagesHandler{}
}

func (h *ImagesHandler) GetImages(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world!"))
}
