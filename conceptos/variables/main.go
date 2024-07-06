package main

import "fmt"

func main() {
	// Variables en GO
	// Forma N.º 1 - Declarar variables y luego asignar
	/*
		var apple string
		var banana string
		var orange string
		apple = "🍎"
		banana = "🍌"
		orange = "🍊"
	*/

	// Forma N.º 2 - Declarar y asignar las variables en una misma línea
	/*
		var apple string = "🍎"
		var banana string = "🍌"
		var orange string = "🍊"
	*/

	// Forma N.º 3 - Agrupar variables
	/*
		var (
			apple  string = "🍎"
			banana string = "🍌"
			orange string = "🍊"
		)
	*/

	// Forma N.º 4 - Declarar y asignar TODAS las variables en una sola línea
	/*var apple, banana, orange string = "🍎", "🍌", "🍊"*/

	// Forma N.º 4 - Operador de variable corta
	apple, banana, orange := "🍎", "🍌", "🍊"
	// El operador de variable corta solo se puede usar cuando declaremos variables nuevas, no para reasignar un valor
	apple, lemon := "manzana", "🍋"

	fmt.Println(apple, banana, orange, lemon)
}
