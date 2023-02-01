package main

import (
	"fmt"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./root"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("String server at 8080\n")

}
