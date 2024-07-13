package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type MyInt int
type MyIntV2 int

func main() {
	var num1 MyInt = 2
	var num2 MyInt = 6
	var num3 MyIntV2 = 6
	var num4 MyIntV2 = 6

	fmt.Println(sum(2, 5, 8, 9))
	fmt.Println(sum(2.5, 5.2, 8.9, 9.1))
	//fmt.Println(sum("a", "b")) Esto daría error
	fmt.Println(sum(num2, num1))
	fmt.Println(sum(num3, num4))
	fmt.Println(sum2(num3, num4))
	fmt.Println(sum3(4.6, 9.3))

	fmt.Println(Includes([]string{"a", "b", "c"}, "c"))
	fmt.Println(Includes([]string{"a", "b", "c"}, "d"))
	fmt.Println(Includes([]int{1, 12, 24}, 24))

	fmt.Println(filter([]int{1, 12, 34, 56, 7, 89, 0}, lessThanTwenty))
}

// Any NO es la solución, puesto que no se pueden sumar por ejemplo, booleanos o strings
// Para esto existen los parámetros de tipo
// el caracter '~' es para que la función trabaje con todos los tipos relacionados. En este caso MyInt hereda de int y por ende es un tipo asociado a este
// A esto se le conoce como 'elementos de aproximación' ~int → cualquier tipo de dato entero o derivado
func sum[T ~int | float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// Pero los tipos numéricos son MUCHOS, si quisiéramos que la función tomara todos, sería MUY larga
// por eso, existe la posibilidad de usar interfaces en su

type Number interface {
	~int | ~float64 | ~float32 | ~uint
}

// De esta forma la función no se vuelve tan complicada de leer
func sum2[T Number](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// Sin embargo, como esto es altamente utilizado, GO ya cuenta con un paquete para esto
// El paquete es 'constraints' (se debe importar)
func sum3[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// Comparable permite comparar valores, no sumar, ni restar, etc.
// solo podremos usar los operadores '==' y '!='

func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

// 'constraints.Ordered' nos permite trabajar comn operadores como '<', '>', '<=', '>='

func filter[T constraints.Ordered](nums []T, callback func(T) bool) []T {
	result := make([]T, 0, len(nums))

	for _, num := range nums {
		if callback(num) {
			result = append(result, num)
		}
	}
	return result
}

func lessThanTwenty(num int) bool {
	return num < 20
}
