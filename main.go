package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	portEnvVar := os.Getenv("PORT")
	if portEnvVar == "" {
		fmt.Println("PORT environment variable hasn't been set, aborting...")
		return
	}

	portNumber, err := strconv.Atoi(portEnvVar)
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.GET("/ping", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"result": "pong",
		})
	})

	router.LoadHTMLGlob("public/templates/*")
	router.GET("/timer", func (c *gin.Context) {
		c.HTML(http.StatusOK, "timer.tmpl", gin.H {
			"time": time.Now().UTC().Format(time.RFC822),
		})
	})

	if err = router.Run(fmt.Sprintf(":%d", portNumber)); err != nil {
		panic(err)
	}
}
