// main.go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", getHomePage)

	r.Run(":28080")
}

func getHomePage(c *gin.Context) {
	c.String(200, "Hello, Gin")
}
