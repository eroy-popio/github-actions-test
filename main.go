package main

import "github.com/gin-gonic/gin"
import "fmt"
import "os"

func main() {
	r := gin.Default()

	host := os.Getenv("HOSTNAME")
	reply:= fmt.Sprintf("Hello from %s V2 \n", host)
	appName := os.Getenv("APPNAME")
	if appName == "" {
		appName = "app"
	}

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, reply)
	})
        
	r.GET("/"+appName+"/hello", func(c *gin.Context) {
		c.String(200, reply)
	})

	r.GET("/"+appName+"/exit", func(c *gin.Context) {
		os.Exit(0)
	})

	r.Run(":3000")
}
