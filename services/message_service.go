package services

import (
	"helloworld/domain"
	"helloworld/models"
	"time"
)

func CreateMessage(message *models.Message) (*models.Message, error) {
	message.CreatedAt = time.Now()
	message, err := domain.Create(message)
	if err != nil {
		return nil, err
	}
	return message, nil
}