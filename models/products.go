package models

import (
	"gocrud/db"
)

// We can create a entity using a struct with a type,
// this is similar to create a interface on typescript.
// Case we create a property where the first letter is in lower case,
// the property cannot be used on html ( can use on another places??)

// Go does not have class like other programming languages, to use object oriented programming
// we can use structs and functions just like this bellow.
// It's important to remember that to a function become public the first letter must be upper case
type Product struct {
	Id          int
	Name        string
	Price       float64
	Quantity    int
	Description string
}

func SearchAllProducts() []Product {
	db := db.ConnectWithDatabase()
	products, dbErr := db.Query("select * from products;")

	if dbErr != nil {
		panic(dbErr.Error())
	}
	defer products.Close()

	p := Product{}
	productsCollection := []Product{}

	for products.Next() {
		err := products.Scan(&p.Id, &p.Name, &p.Description, &p.Quantity, &p.Price)
		if err != nil {
			panic(err.Error())
		}
		productsCollection = append(productsCollection, p)
	}

	defer db.Close()
	return productsCollection
}

func PostProduct(name, description string, price float64, quantity int) {
	db := db.ConnectWithDatabase()

	query, err := db.Prepare("insert into products(name, description, price, quantity) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	query.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectWithDatabase()

	query, queryErr := db.Prepare("delete from products where id = $1")
	if queryErr != nil {
		panic(queryErr.Error())
	}

	query.Exec(id)
	defer db.Close()
}

func FindProduct(id string) Product {
	db := db.ConnectWithDatabase()

	product, err := db.Query("select * from products where id = $1 order by id desc", id)
	if err != nil {
		panic(err.Error())
	}

	productSerialized := Product{}
	for product.Next() {
		err = product.Scan(
			&productSerialized.Id,
			&productSerialized.Name,
			&productSerialized.Description,
			&productSerialized.Price,
			&productSerialized.Quantity,
		)

		if err != nil {
			panic(err.Error())
		}
	}

	defer db.Close()
	return productSerialized
}

func UpdateProduct(id, quantity int, name, description string, price float64) {
	db := db.ConnectWithDatabase()
	updateProduct, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(name, description, price, quantity, id)
	defer db.Close()
}
