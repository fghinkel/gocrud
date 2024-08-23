package routes

import (
	"gocrud/controllers"
	"log"
	"net/http"
)

func init() {
	log.Println("Server running on http://localhost:8077/")
}

// Here we pass to our html file a writter, who is being executed ( index )
// and in case we want to pass any information to the index we can use the last param
func GetRoutes() {
	http.HandleFunc("/", controllers.GetProducts)
	http.HandleFunc("/new", controllers.NewProduct)
	http.HandleFunc("/edit", controllers.EditProduct)
	http.HandleFunc("/insert", controllers.PostProduct)
	http.HandleFunc("/delete", controllers.DeleteProduct)
	http.HandleFunc("/update", controllers.PutProduct)
}
