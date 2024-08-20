package main

import (
	"gocrud/routes"
	"net/http"
)

// The main function create a http server running on 8000 port
// to start the server run `go run main.go` . The serve is up by the
// ListenAndServe function.

// The HandleFunc function serves to bind a route with a function,
// on this case is binding "/" with the index function
func main() {
	routes.GetRoutes()
	http.ListenAndServe(":8077", nil)
}
