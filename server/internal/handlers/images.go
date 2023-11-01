package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/harmoniemand/stopmotion/internal/configuration"
	"github.com/harmoniemand/stopmotion/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	log "github.com/sirupsen/logrus"
)

type ImagesHandler struct {
	config *configuration.Config
}

func NewImagesHandler(c *configuration.Config) *ImagesHandler {
	return &ImagesHandler{
		config: c,
	}
}

func (h *ImagesHandler) GetImageAsFile(w http.ResponseWriter, r *http.Request) {
	log.Debug("The GetImages handler is executing!")

	projectID := mux.Vars(r)["project_id"]
	if projectID == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, e := w.Write([]byte("project_id is required"))
		if e != nil {
			log.Errorf("Error writing response: %v", e)
		}
		return
	}

	imgID := mux.Vars(r)["img_id"]
	if imgID == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, e := w.Write([]byte("id is required"))
		if e != nil {
			log.Errorf("Error writing response: %v", e)
		}
	}

	imgDir := path.Join("_images/", projectID)
	imgPath := path.Join(imgDir, imgID+".png")

	http.ServeFile(w, r, imgPath)
}

func (h *ImagesHandler) PostImage(w http.ResponseWriter, r *http.Request) {
	log.Debug("The PostImage handler is executing!")

	//nolint: gomnd // 32 << 20 is 32MB - defines filesize limit
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Errorf("Error parsing multipart form: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		_, e := w.Write([]byte("Bad request"))
		if e != nil {
			log.Errorf("Error writing response: %v", e)
		}
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, e := w.Write([]byte("Bad request"))
		if e != nil {
			log.Errorf("Error writing response: %v", e)
		}
		return
	}

	defer file.Close()

	projectID := mux.Vars(r)["project_id"]
	imgDir := path.Join("_images/", projectID)
	imgID := uuid.New().String()

	img := models.Image{
		ID:        imgID,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		Filename:  imgID + ".png",
	}

	// save image to disk

	err = os.MkdirAll(imgDir, os.ModePerm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, e := w.Write([]byte("Internal server error"))
		if e != nil {
			log.Errorf("Error writing response: %v", e)
		}
		return
	}

	f, err := os.OpenFile(path.Join(imgDir, img.Filename), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, e := w.Write([]byte("Internal server error"))
		if e != nil {
			log.Errorf("Error writing response: %v", e)
		}
		return
	}

	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, e := w.Write([]byte("Internal server error"))
		if e != nil {
			log.Errorf("Error writing response: %v", e)
		}
		return
	}

	// save image to db

	log.Debugf("db connection: %v", h.config.DBConnection)
	client, err := mongo.Connect(r.Context(), options.Client().ApplyURI(h.config.DBConnection))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, e := w.Write([]byte("Internal server error"))
		if e != nil {
			log.Errorf("Error writing response: %v", e)
		}
		return
	}

	collection := client.Database(h.config.DatabaseName).Collection("projects")
	_, err = collection.UpdateOne(r.Context(), bson.M{"id": projectID}, bson.M{"$push": bson.M{"images": img}})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, e := w.Write([]byte("Internal server error"))
		if e != nil {
			log.Errorf("Error writing response: %v", e)
		}
		return
	}

	err = client.Disconnect(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, e := w.Write([]byte("Internal server error"))
		if e != nil {
			log.Errorf("Error writing response: %v", e)
		}
		return
	}

	json, err := json.Marshal(img)
	if err != nil {
		log.Errorf("Error marshalling json: %v", err)
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

func (h *ImagesHandler) DeleteImage(w http.ResponseWriter, r *http.Request) {
	log.Debug("The DeleteImage handler is executing!")

	projectID := mux.Vars(r)["project_id"]
	imgID := mux.Vars(r)["img_id"]

	imgDir := path.Join("_images/", projectID)
	imgPath := path.Join(imgDir, imgID+".png")

	err := os.Remove(imgPath)
	if err != nil {
		log.Errorf("Error removing image: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, e := w.Write([]byte("Internal server error"))
		if e != nil {
			log.Errorf("Error writing response: %v", e)
		}
	}

	log.Debugf("db connection: %v", h.config.DBConnection)
	client, err := mongo.Connect(r.Context(), options.Client().ApplyURI(h.config.DBConnection))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, e := w.Write([]byte("Internal server error"))
		if e != nil {
			log.Errorf("Error writing response: %v", e)
		}
		return
	}

	collection := client.Database(h.config.DatabaseName).Collection("projects")
	_, err = collection.UpdateOne(r.Context(), bson.M{"id": projectID}, bson.M{"$pull": bson.M{"images": bson.M{"id": imgID}}})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, e := w.Write([]byte("Internal server error"))
		if e != nil {
			log.Errorf("Error writing response: %v", e)
		}
		return
	}

}
