package controller

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/qinyul/messaging-api/helpers"
	"github.com/qinyul/messaging-api/models"
	"github.com/qinyul/messaging-api/service"
	"github.com/qinyul/messaging-api/utils"
)

type MessageController struct {
	service service.MessageService
}

func NewMessageController(service service.MessageService) *MessageController {
	return &MessageController{service: service}
}

func (c *MessageController) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message models.Message
	slog.Info("message-controller CreateMessage starting")
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		slog.Error("message-controller error decode")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if message.Body == "" {
		http.Error(w, "Message body cannot be empty", http.StatusBadRequest)
		return
	}
	slog.Info("message-controller calling message service")
	createdMessage, err := c.service.CreateMessage(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = helpers.NewResponseToJson(w, http.StatusCreated, createdMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	slog.Info("message-controller CreateMessage finish")
}

func (c *MessageController) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	var message models.Message
	slog.Info("message-controller UpdateMessage starting")
	vars := mux.Vars(r)
	id := vars["id"]

	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		slog.Error("message-controller error decode")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if message.Body == "" {
		http.Error(w, "Message body cannot be empty", http.StatusBadRequest)
		return
	}
	slog.Info("message-controller calling message service")
	updatedMessage, err := c.service.UpdateMessage(utils.UUIDParser(id), message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = helpers.NewResponseToJson(w, http.StatusCreated, updatedMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	slog.Info("message-controller CreateMessage finish")
}

func (c *MessageController) GetMessages(w http.ResponseWriter, r *http.Request) {
	slog.Info("message-controller GetMessages starting")
	messages, err := c.service.GetMessages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = helpers.NewResponseToJson(w, http.StatusOK, messages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	slog.Info("message-controller GetMessages finish")
}

func (c *MessageController) GetMessageById(w http.ResponseWriter, r *http.Request) {
	slog.Info("message-controller GetMessageById starting")
	vars := mux.Vars(r)
	id := vars["id"]

	message, err := c.service.GetMessageById(utils.UUIDParser(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = helpers.NewResponseToJson(w, http.StatusOK, message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	slog.Info("message-controller GetMessageById finish")
}

func (c *MessageController) DeleteMessageById(w http.ResponseWriter, r *http.Request) {
	slog.Info("message-controller DeleteMessageById starting")
	vars := mux.Vars(r)
	id := vars["id"]

	err := c.service.DeleteMessageById(utils.UUIDParser(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = helpers.NewResponseToJson(w, http.StatusAccepted, models.BaseResponse{
		Code:    fmt.Sprintf("%d", http.StatusAccepted),
		Message: "Message sucessfully deleted",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	slog.Info("message-controller DeleteMessageById finish")
}
