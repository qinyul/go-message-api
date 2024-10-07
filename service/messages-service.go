package service

import (
	"github.com/google/uuid"
	"github.com/qinyul/messaging-api/models"
	"github.com/qinyul/messaging-api/repository"
)

type MessageService interface {
	CreateMessage(message models.Message) (models.Message, error)
	UpdateMessage(id uuid.UUID, message models.Message) (*models.Message, error)
	GetMessages() ([]models.Message, error)
	GetMessageById(uuid uuid.UUID) (*models.Message, error)
	DeleteMessageById(uuid uuid.UUID) error
}

type messageService struct {
	repo *repository.MessageRepository
}

func NewMessageService(repo *repository.MessageRepository) *messageService {
	return &messageService{repo: repo}
}

func (s *messageService) CreateMessage(message models.Message) (models.Message, error) {
	return s.repo.CreateMessage(message)
}

func (s *messageService) UpdateMessage(id uuid.UUID, message models.Message) (*models.Message, error) {
	return s.repo.UpdateMessage(id, message)
}

func (s *messageService) GetMessages() ([]models.Message, error) {
	return s.repo.GetMessages()
}

func (s *messageService) GetMessageById(uuid uuid.UUID) (*models.Message, error) {
	return s.repo.GetMessageById(uuid)
}

func (s *messageService) DeleteMessageById(uuid uuid.UUID) error {
	return s.repo.DeleteMessageById(uuid)
}
