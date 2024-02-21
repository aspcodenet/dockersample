package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func start(c *gin.Context) {
	computerName, _ := os.Hostname() //hallo
	c.String(http.StatusOK, "Tjena "+computerName+" this is cool3")
}

func main() {
	router := gin.Default()
	router.GET("/", start)

	router.Run(":8080")

}
