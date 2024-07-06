package main

import "fmt"

// Las estructuras nos permiten almacenar una colección de campos
// Generalmente lo debemos hacer a nivel paquete para que todos los elementos del paquete tengan acceso

// Person is a structure that defines a person
type Person struct {
	Name        string
	Age         uint8
	HasChildren bool
}

func main() {
	// Formas de instanciar una persona
	// Forma N.º 1 - Especificando nombre atributo
	alexys := Person{
		Name:        "Alexys",
		Age:         42,
		HasChildren: false,
	}

	// Forma N.º 2 - No especificando nombre de atributo
	// Go infiere que los campos corresponden al orden en que se definen
	// en este caso, Name, Age, HasChildren
	beto := Person{
		"Beto",
		33,
		true,
	}

	// Forma N.º 3 - Usando el operador de composición
	// Nos permite especificar solo los campos que nos interesan
	// en este caso, Age
	alejo := Person{Age: 32}
	//los demás valores quedarán con sus respectivos zero values
	//string = ""
	//uint8 = 0
	//bool = false

	// Punteros a estructuras
	alvaro := &alexys
	alvaro.Age = 50 // Esto modifica la información del puntero y la estructura a la que estanos apuntando

	fmt.Printf("%+v\n", alexys.Name)
	fmt.Printf("%+v\n", beto.HasChildren)
	fmt.Printf("%+v\n", alejo.Age)
	fmt.Printf("%+v\n", alexys)
	fmt.Printf("%+v\n", alvaro)
}
