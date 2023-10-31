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
	models, err := h.ProjectStore.GetAllProjects(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, e := w.Write([]byte("Internal server error"))
		if e != nil {
			log.Errorf("Error writing response: %v", e)
		}
		return
	}

	json, err := json.Marshal(models)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, e := w.Write([]byte("Internal server error"))
		if e != nil {
			log.Errorf("Error writing response: %v", e)
		}

		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(json)
	if err != nil {
		log.Errorf("Error writing response: %v", err)
	}
}

func (h *ProjectsHandler) GetProject(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	log.Debugf("Getting project with id: %v", id)

	model, err := h.ProjectStore.GetProject(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, e := w.Write([]byte("Internal server error"))
		if e != nil {
			log.Errorf("Error writing response: %v", e)
		}
		return
	}

	log.Debugf("Found project")

	json, err := json.Marshal(model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, e := w.Write([]byte("Internal server error"))
		if e != nil {
			log.Errorf("Error writing response: %v", e)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(json)
	if err != nil {
		log.Errorf("Error writing response: %v", err)
	}
}

func (h *ProjectsHandler) PostProject(w http.ResponseWriter, r *http.Request) {
	var model models.Project
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, e := w.Write([]byte("Bad request"))
		if e != nil {
			log.Errorf("Error writing response: %v", e)
		}
		return
	}

	insertedModel, err := h.ProjectStore.InsertProject(r.Context(), model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, e := w.Write([]byte("Internal server error"))
		if e != nil {
			log.Errorf("Error writing response: %v", e)
		}
		return
	}

	json, err := json.Marshal(insertedModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, e := w.Write([]byte("Internal server error"))
		if e != nil {
			log.Errorf("Error writing response: %v", e)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, e := w.Write(json)
	if e != nil {
		log.Errorf("Error writing response: %v", e)
	}
}
