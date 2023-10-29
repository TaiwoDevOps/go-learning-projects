package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

/**!HTTP Request and Response
ResponseWriter is the response sent to the users request
Request is the request sent by the user and can be handled by the server based on different methods
*/

// handle form
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err) //%v this is a formatter that prints its value compare to %T that prints the type of the value
		return
	}

	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name) //%s this is a formatter and is a string interpolation
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// confirm if path is not correct and throw and error
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// if request method is not GET
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	// send response if everything is ok
	fmt.Fprintf(w, "hello!")
}

func main() {
	// serve files from the static directory
	serveFiles := http.FileServer(http.Dir("./static"))
	http.Handle("/", serveFiles) // if path is /, then serve the index file from the static directory

	// func for handling other directories
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	// start the server and listen  for error
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
