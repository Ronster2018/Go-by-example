package main

import (
	"fmt"
	"net/http"
)

func index_handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Woah, Go is neat!") // going to format based on what's specified and the output what we ask for
}

func about_handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Welcome to this cool about page")
}

func main() {
	// Handlers
	http.HandleFunc("/", index_handler) // Takes in the path we want to use and then what ever function
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000", nil) // Port and other server options
}
