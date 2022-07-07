package main

import "github.com/gin-gonic/gin"
import "fmt"
import "os"

func main() {
	r := gin.Default()

        host := os.Getenv("HOSTNAME")
        reply:= fmt.Sprintf("Hello from %s V2\n", host)
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, reply)
	})

	r.Run(":3000")
}
