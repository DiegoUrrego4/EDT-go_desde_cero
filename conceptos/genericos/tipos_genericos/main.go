package main

import "fmt"

// Product es una estructura que maneja tipos genéricos para el campo 'ID'
type Product[T uint | string] struct {
	ID          T
	Description string
	Price       float64
}

func main() {
	product1 := Product[uint]{ID: 1, Description: "Zapatos", Price: 30}
	product2 := Product[string]{ID: "c2f301b4-7ede-43ab-bad0-8d981879d234", Description: "Camisa", Price: 40.6}
	fmt.Println(product1)
	fmt.Println(product2)
}
