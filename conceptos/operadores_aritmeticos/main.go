package main

import "fmt"

func main() {
	// Operadores aritméticos: (), *, /, %, +, -
	var a = (2 + 3) * 2 // si no usamos el paréntesis se realizará primero la multiplicación
	fmt.Println(a)

	// Operadores de asignación: =, +=, -=, *=, /=, %=
	var b int = 5
	b += 2
	fmt.Println(b) // b: 7

	// Declaración post-incremento y post-decremento: ++, --
	// NO son una expresión, sino una declaración
	var c int = 6
	// fmt.Println(c++) -> Esto NO es permitido en GO, ya que 'c++' es una declaración NO un valor
	// Lo correcto sería esto:
	c++ // Acá se modifica el valor de la variable
	fmt.Println(c)
}
