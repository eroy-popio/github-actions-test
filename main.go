package main

import "github.com/gin-gonic/gin"
import "fmt"
import "os"

func NewString(s string) string {
	return s+"12345"
}

func main() {
	r := gin.Default()

	host := os.Getenv("HOSTNAME")
	reply:= fmt.Sprintf("Hello from %s V6 \n", host)
	newStr := NewString("Hello, World! ")
	appName := os.Getenv("APPNAME")
	if appName == "" {
		appName = "app"
	}

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, newStr)
	})
        
	r.GET("/"+appName+"/hello", func(c *gin.Context) {
		c.String(200, reply)
	})

	r.GET("/"+appName+"/exit", func(c *gin.Context) {
		os.Exit(0)
	})

	r.GET("/"+appName+"/health", func(c *gin.Context) {
		c.String(200, "Success")
	})

	r.Run(":3000")
}
