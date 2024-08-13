package main

// Go only let us import packages that are in use, sometimes, packages aren't always in use,
// in this cases we can pass the _ before the package name

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func connectWithDatabsae() *sql.DB {
	connection := "user=postgres dbname=gocrud password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}

// We can create a entity using a struct with a type,
// this is similar to create a interface on typescript.
// Case we create a property where the first letter is in lower case,
// the property cannot be used on html ( can use on another places??)
type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

// This is used to bring all the html files from templates file
var templates = template.Must(template.ParseGlob("templates/*.html"))

// The main function create a http server running on 8000 port
// to start the server run `go run main.go` . The serve is up by the
// ListenAndServe function.

// The HandleFunc function serves to bind a route with a function,
// on this case is binding "/" with the index function
func main() {
	db := connectWithDatabsae()
	defer db.Close()

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

// Here we pass to our html file a writter, who is being executed ( index )
// and in case we want to pass any information to the index we can use the last param
func index(writer http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "Camiseta", Description: "Camisa grêmio", Price: 39, Quantity: 5},
		{Name: "Tenis", Description: "Novo", Price: 89, Quantity: 3},
		{Name: "Fone", Description: "Grande", Price: 59, Quantity: 3},
		{Name: "Monitor", Description: "4k", Price: 59, Quantity: 3},
	}

	err := templates.ExecuteTemplate(writer, "Index", products)
	if err != nil {
		log.Fatalln(err)
	}
}

// To embed the go code on html, we need to use
// the following code on the top of the html file
// {{ define "template_name" }}
// And in the end, we close the file using
// {{ end }}
