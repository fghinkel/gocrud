package routes

import (
	"gocrud/controllers"
	"net/http"
)

// Here we pass to our html file a writter, who is being executed ( index )
// and in case we want to pass any information to the index we can use the last param
func GetRoutes() {
	http.HandleFunc("/", controllers.GetProducts)
	http.HandleFunc("/new", controllers.NewProduct)
	http.HandleFunc("/insert", controllers.PostProduct)
}
