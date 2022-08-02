package app

import "helloworld/controllers"

func routes() {
	router.POST("/messages", controllers.CreateMessage)
	router.PUT("/messages/:message_id", controllers.UpdateMessage)
}