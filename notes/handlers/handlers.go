package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RamiroCuenca/go-rest-notesApi/common/handler"
	"github.com/RamiroCuenca/go-rest-notesApi/common/logger"
	"github.com/RamiroCuenca/go-rest-notesApi/notes/models"
)

type Storage interface {
	Create(*models.Note) error
	Update(*models.Note) error
	GetAll() ([]models.Note, error)
	GetById(id int) (models.Note, error)
	Delete(id int) error
}

type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{s}
}

// This handler is responsible of creating a new Note
func (s *Service) Create(w http.ResponseWriter, r *http.Request) {
	// 1° Decode the json received on a Note object
	data := models.Note{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		logger.Log().Infof("Error decoding: %v", err)
		handler.SendError(w, http.StatusBadRequest)
		return
	}

	// 2° Send the model to the controller
	// err = controllers
	// db.CreateNoteController(data)

	// ° Send response
	handler.SendResponse(w, http.StatusOK, []byte("Notes handler working"))
	return
}
