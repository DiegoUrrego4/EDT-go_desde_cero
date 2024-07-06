package main

import "fmt"

func main() {
	// bool string, numeric
	var a bool = true
	var b string = "EDteam"

	// c uint → Valores positivos
	// c int → Valores positivos y negativos (mirar archivo datatype en esta mismo package)
	var c uint8 = 33
	// Si usamos 'uint' sin número se tomará por defecto dependiendo de la arquitectura de nuestro ordenador
	// si nuestro ordenador es x32 uint será uint32, si es x64 será uint64

	// byte → es un alias para uint8, usar cualquiera de los dos sería lo mismo
	var d byte = 255

	// rune → alias para int32 o representaciones de códigos unicode
	// Para los códigos unicode debemos usar las comillas simples ''
	var e rune = 'a'

	var f float32 = 124.77

	fmt.Printf("Tipo: %T, Valor: %v\n", a, a)
	fmt.Printf("Tipo: %T, Valor: %v\n", b, b)
	fmt.Printf("Tipo: %T, Valor: %v\n", c, c)
	fmt.Printf("Tipo: %T, Valor: %v\n", d, d)
	fmt.Printf("Tipo: %T, Valor: %v\n", e, e) // Esto imprimiría: Tipo: int32, Valor: 97
	// 97 sería porque este es el unicode de ese carácter
	fmt.Printf("Tipo: %T, Valor: %v\n", f, f)

	// Operaciones de tipo
	// NO PODEMOS hacer operaciones entre tipos de datos distintos
	// Ejemplos:
	var g uint8 = 255
	var h uint16 = 2550

	// i:= g + h → Esto provocará un error

	// Para corregirlo debemos castear, y se hace de la siguiente manera:
	i := uint16(g) + h // Se castea usando la función del mismo tipo de dato que requiramos (normalmente se castea el de menor valor a mayor)
	// NOTA: a pesar de castear a 'g' este seguirá siendo de tipo uint8 en cualquier otra línea de código distinta a la anterior

	fmt.Printf("Tipo: %T, Valor: %v\n", i, i)

	// Si necesitáramos una lógica que usaremos luego podemos 'omitir' esa variable de la siguiente manera
	_ = uint16(g) + h
	// a '_' se le conoce como el identificador 'blank' y nos permite no perder la lógica que estemos haciendo, pero no usando
	// al no tener una variable asignada GO nos permite correr el código sin problema así lo estemos 'usando' acá
	// IMPORTANTE: NO SE PUEDE USAR el operador de asignación corta ':=', solo el operador de asignación '=' o nuestro programa fallará

	// VALOR ZERO: Es el valor que le asigna GO a las variables que fueron declaradas, pero NO inicializadas
	// Para cada tipo de dato GO le asigna un valor diferente
	// Ejemplos:
	var j string
	fmt.Printf("Tipo: %T, Valor: %v\n", j, j) // Esto imprimirá un string vacío ""
	// Para los tipos numéricos será 0
	// Y para los tipos booleanos será false
}
