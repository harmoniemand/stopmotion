package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harmoniemand/stopmotion/internal/configuration"
	"github.com/harmoniemand/stopmotion/internal/models"
	"github.com/harmoniemand/stopmotion/internal/stores"

	log "github.com/sirupsen/logrus"
)

type ProjectsHandler struct {
	Config       *configuration.Config
	ProjectStore *stores.ProjectStore
}

func NewProjectsHandler(projectStore *stores.ProjectStore) *ProjectsHandler {
	return &ProjectsHandler{
		ProjectStore: projectStore,
	}
}

func (h *ProjectsHandler) GetProjects(w http.ResponseWriter, r *http.Request) {

	models, error := h.ProjectStore.GetAllProjects(r.Context())
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	json, err := json.Marshal(models)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func (h *ProjectsHandler) GetProject(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	log.Debugf("Getting project with id: %v", id)

	model, error := h.ProjectStore.GetProject(r.Context(), id)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	log.Debugf("Found project")

	json, err := json.Marshal(model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func (h *ProjectsHandler) PostProject(w http.ResponseWriter, r *http.Request) {

	var model models.Project
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}

	new_model, err := h.ProjectStore.InsertProject(r.Context(), model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	json, err := json.Marshal(new_model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func (h *ProjectsHandler) PostImageToProject(w http.ResponseWriter, r *http.Request) {

	var model models.Image
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}

	id := mux.Vars(r)["id"]

	err = h.ProjectStore.AddImageToProject(r.Context(), id, model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("image saved"))
}
