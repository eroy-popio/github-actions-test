package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"helloworld/models"
	"helloworld/services"
	"net/http"
	"strconv"
)

func CreateMessage(c *gin.Context) {
	var message models.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		fmt.Println(err)
		return
	}
	err := services.CreateMessage(&message)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusCreated, "Message Created Successfully")
}

func getMessageId(msgIdParam string) (int64, error) {
	msgId, err := strconv.ParseInt(msgIdParam, 10, 64)
	if err != nil {
		return 0, err
	}
	return msgId, nil
}


func UpdateMessage(c *gin.Context) {
	msgId, err := getMessageId(c.Param("message_id"))
	if err != nil {
		panic(err)
		return
	}
	var message models.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		panic(err)
		return
	}
	message.Id = msgId
	err = services.UpdateMessage(&message)
	if err != nil {
		panic(err)
		return
	}
	c.JSON(http.StatusOK, "Message Updated Successfully")
}
