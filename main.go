package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func start(c *gin.Context) {
	computerName2, _ := os.Hostname() //hallo
	c.String(http.StatusOK, "Tjena "+computerName2+" this is cool3")
}

func main() {
	//hej
	//hej2
	router := gin.Default()
	router.GET("/", start)

	router.Run(":8080")

}
