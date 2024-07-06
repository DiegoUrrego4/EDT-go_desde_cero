package main

import "fmt"

func main() {
	// Operadores de comparación: >, <, ==, !=, <=, >=
	fmt.Println((4 + 5) > 6) // true
	fmt.Println(4 == 4)      // true
	fmt.Println(4 != 4)      // false

	// Operadores lógicos: && (y), || (o)
	var age uint = 33
	fmt.Println("¿Es adulto?", age >= 18 && age <= 60)
	fmt.Println("¿Es niño o anciano?", age <= 18 && age >= 60)

	// Operador lógico unario: !
	fmt.Println(!(4 == 4)) // devolverá false

}
