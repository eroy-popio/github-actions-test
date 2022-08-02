package services

import (
	"helloworld/domain"
	"helloworld/models"
	"time"
)

func CreateMessage(message *models.Message)  error {
	message.CreatedAt = time.Now()
	err := domain.Create(message)
	if err != nil {
		return err
	}
	return nil
}

func UpdateMessage(message *models.Message) error {
	err := domain.Update(message)
	if err != nil {
		return err
	}
	return nil
}