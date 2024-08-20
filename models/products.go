package models

import "gocrud/db"

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
	Description string
	Price       float64
	Quantity    int
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
