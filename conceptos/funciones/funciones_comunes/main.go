package main

import (
	"fmt"
	"strings"
)

func main() {
	//greet("Jennyfer", "Salcedo")
	// name := "Jennyfer"
	// toUpperCase(&name)
	// fmt.Println(name)

	// result := sum(2, 3)
	// fmt.Println(result)

	lower, upper := convert("EDteam")
	fmt.Println(lower, upper)
}

// funciones_comunes con parámetros por valor y referencia
/*
func greet(firstName, lastname string) {
	fmt.Printf("Hello %s %s", firstName, lastname)
}*/

func toUpperCase(text *string) { // Las funciones_comunes siempre reciben una copia del valor
	// Si le pasamos un puntero, debemos referenciar el valor
	*text = strings.ToUpper(*text)
}

// Funciones con retorno
func sum(a, b int) int {
	return a + b
}
func convert(text string) (lower string, upper string) {
	lower = strings.ToLower(text)
	upper = strings.ToUpper(text)
	return
}
