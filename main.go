package main

import "net/http"

func main() {
	fileServer := http.FileServer(http.Dir("./root"))
	http.Handle("/", fileServer)
}
