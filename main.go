package main

import (
	"fmt"
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

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	fmt.Println(fmt.Sprintf("Launching on port %d...", portNumber))
	err = http.ListenAndServe(fmt.Sprintf(":%d", portNumber), nil)

	if err != nil {
		panic(err)
	}
}
