package main

import (
	"fmt"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	portNumber := 8008
	fmt.Println(fmt.Sprintf("Launching on port %d...", portNumber))
	err := http.ListenAndServe(fmt.Sprintf(":%d", portNumber), nil)

	if err != nil {
		panic(err)
	}
}
