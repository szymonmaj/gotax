package main

import "fmt"


var (
	shoppingCartAni []Product
	shoppingCartKasi []Product
	shoppingCart []Product
)

type Product struct {
	name  string
	price int
}

func main() {


	product1 := createProduct("shoes", 200)
	product2 := createProduct("jacket", 500)
	product3 := createProduct("T-shirt", 100)




	addProduct(product1)
	fmt.Println("Zawartość koszyka Ani", shoppingCartAni)


	addProduct(product2)
	fmt.Println("Zawartość koszyka Ani",shoppingCartAni)


	addProduct(product3)
	fmt.Println("Zawartość koszyka Ani",shoppingCartAni)

	addProduct(product3)
	fmt.Println("Zawartość koszyka Kasi",shoppingCartKasi)


}


func addProduct(product Product)  {


	shoppingCart = append(shoppingCart, product)


}

func createProduct(name string, price int) (product Product) {

	product = Product{name, price}

	return product
}
