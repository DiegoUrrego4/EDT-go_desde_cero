package main

import "fmt"

func main() {
	// punteros: variables que almacenan la dirección en memoria de un valor
	var color string = "🟥"

	fmt.Printf("Tipo: %T, Valor: %s\n", color, color)
	// Con esto sabemos que el valor que tiene 'color' es el cuadrado rojo
	// Ahora queremos conocer su dirección en memoria, para ello debemos usar el operador de dirección '&'
	fmt.Printf("Tipo: %T, Valor: %s, Dirección: %v\n", color, color, &color)

	// Ahora declararemos nuestro puntero:
	// Para ello debemos agregar un '*' al tipo de dato para indicarle a GO que lo que queremos almacenar es una dirección en memoria
	var colorPointer *string // 'colorPointer' será de tipo puntero y apuntará a cualquier variable de tipo string
	colorPointer = &color
	*colorPointer = "🟦" // Acá cambia tanto el color como el puntero de rojo a azul

	fmt.Printf("Tipo: %T, Valor: %s, Dirección: %v\n", color, color, &color)
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciación: %s\n", colorPointer, colorPointer, *colorPointer)
}
