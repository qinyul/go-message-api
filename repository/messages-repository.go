package repository

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/google/uuid"
	"github.com/qinyul/messaging-api/models"
	"gorm.io/gorm"
)

type MessageRepositoryInterface interface {
	CreateMessage(message models.Message) (models.Message, error)
	UpdateMessage(uuid uuid.UUID, message models.Message) (models.Message, error)
	GetMessages([]models.Message) ([]models.Message, error)
	GetMessageById(uuid uuid.UUID) (*models.Message, error)
}

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	if db == nil {
		log.Fatal("DB not initialized")
	}
	return &MessageRepository{db: db}
}

func (r *MessageRepository) CreateMessage(message models.Message) (models.Message, error) {

	newUser := models.Message{
		ID:             uuid.New(),
		ConversationID: uuid.New(),
		SenderId:       message.ID,
		Body:           message.Body,
		IsSoftDeleted:  false,
	}

	slog.Info("Inserting data to database")
	query := r.db.Create(&newUser)

	if query.Error != nil {
		fmt.Printf("CreateMessage:: error when inserting message to DB %v", query.Error)
		return models.Message{}, query.Error
	}

	return newUser, nil
}

func (r *MessageRepository) UpdateMessage(id uuid.UUID, message models.Message) (*models.Message, error) {

	var updatedMessage models.Message

	if err := r.db.First(&updatedMessage, id).Error; err != nil {
		return nil, err
	}

	slog.Info("Updating data to database")

	updatedMessage.Body = message.Body
	query := r.db.Save(&updatedMessage)
	if query.Error != nil {
		fmt.Printf("UpdateMessage:: error when updating message to DB %v", query.Error)
		return nil, query.Error
	}

	return &updatedMessage, nil
}

func (r *MessageRepository) GetMessages() ([]models.Message, error) {
	var messages []models.Message
	slog.Info("GetMessages:: getting messages from database")
	query := r.db.Find(&messages)

	if query.Error != nil {
		fmt.Printf("CreateMessage:: error when getting messages from DB %v", query.Error)
		return messages, query.Error
	}
	return messages, nil
}

func (r *MessageRepository) GetMessageById(uuid uuid.UUID) (*models.Message, error) {
	var message models.Message
	query := r.db.First(&message, uuid)

	if query.Error != nil {
		fmt.Printf("GetMessageById:: error when getting message by id from DB %v\n", query.Error)
		return nil, query.Error
	}
	return &message, nil
}

func (r *MessageRepository) DeleteMessageById(uuid uuid.UUID) error {
	var message models.Message
	query := r.db.Delete(&message, uuid)

	if query.Error != nil {
		fmt.Printf("DeleteMessageById:: error when deleting message by id from DB %v\n", query.Error)
		return query.Error
	}

	return nil
}
