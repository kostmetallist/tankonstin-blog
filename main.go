package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
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

	ginEngine := gin.Default()
	ginEngine.GET("/ping", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"result": "pong",
		})
	})

	fmt.Println(fmt.Sprintf("Launching on port %d...", portNumber))
	err = ginEngine.Run(fmt.Sprintf(":%d", portNumber))

	// fileServer := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fileServer)
	// err = http.ListenAndServe(fmt.Sprintf(":%d", portNumber), nil)

	if err != nil {
		panic(err)
	}
}
