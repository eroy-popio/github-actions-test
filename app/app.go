package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"helloworld/domain"
)

var (
	router = gin.Default()
)

func StartingFunction(){
	domain.InitialiseMongoDB()
	fmt.Println("DATABASE STARTED")
	routes()
	router.Run(":8080")
}
