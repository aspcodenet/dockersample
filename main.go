package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func start(c *gin.Context) {
	computerName, _ := os.Hostname()
	c.String(http.StatusOK, "Tjena "+computerName+" this is cool")
}

func main() {
	router := gin.Default()
	router.GET("/", start)

	router.Run(":8080")

}
