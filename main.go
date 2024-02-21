package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func start(c *gin.Context) {
	computerName2, _ := os.Hostname()
	c.String(http.StatusOK, "Tjena "+computerName2+" this is cool")
}

func main() {
	//hej
	router := gin.Default()
	router.GET("/", start)

	router.Run(":8080")

}
