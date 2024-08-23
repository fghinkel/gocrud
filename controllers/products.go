package controllers

import (
	"fmt"
	"gocrud/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// This is used to bring all the html files from templates file
var templates = template.Must(template.ParseGlob("templates/*.html"))

// To embed the go code on html, we need to use
// the following code on the top of the html file
// {{ define "template_name" }}
// And in the end, we close the file using
// {{ end }}
func GetProducts(writer http.ResponseWriter, r *http.Request) {
	productsCollection := models.SearchAllProducts()
	err := templates.ExecuteTemplate(writer, "Index", productsCollection)

	if err != nil {
		log.Fatalln(err)
	}
}

func NewProduct(writer http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(writer, "NewProduct", nil)

	if err != nil {
		log.Fatalln(err)
	}
}

func PostProduct(writer http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("Error:", "The requested source only allows POST Method")
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	quantity := r.FormValue("quantity")

	priceConverted, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Println("Erro na conversão da quantidade:", err)
	}

	quantityConverted, err := strconv.Atoi(quantity)
	if err != nil {
		log.Println("Erro na conversão do preço", err)
	}

	models.PostProduct(name, description, priceConverted, quantityConverted)

	http.Redirect(writer, r, "/", http.StatusMovedPermanently)
}

func DeleteProduct(writer http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProduct(id)
	http.Redirect(writer, r, "/", http.StatusMovedPermanently)

}

func EditProduct(writer http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := models.FindProduct(id)

	err := templates.ExecuteTemplate(writer, "Edit", product)

	if err != nil {
		log.Fatalln(err)
	}
}

func PutProduct(writer http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("Error:", "The requested source only allows POST Method")
	}
	id := r.FormValue("id")
	name := r.FormValue("name")
	price := r.FormValue("price")
	quantity := r.FormValue("quantity")
	description := r.FormValue("description")

	idToInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Erro na conversão do id:", err)
	}

	priceToFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Println("Erro na conversão do preço:", err)
	}

	quantityToInt, err := strconv.Atoi(quantity)
	if err != nil {
		log.Println("Erro na conversão da quantidade:", err)
	}

	models.UpdateProduct(idToInt, quantityToInt, name, description, priceToFloat)
	http.Redirect(writer, r, "/", http.StatusMovedPermanently)
}
