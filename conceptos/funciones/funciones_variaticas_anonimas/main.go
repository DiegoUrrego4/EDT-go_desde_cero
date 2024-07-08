package main

import "fmt"

func main() {
	fmt.Println(sum(2, 3))
	fmt.Println(sum(2, 3, 12))
	fmt.Println(sum(2, 3, 12, 1, 24))

	// Funciones anonimas → funciones que no tienen nombre
	greet := func(name string) {
		fmt.Printf("👋🏼 Hello %s\n", name)
	}
	greet("gophers")

}

// Funciones variaticas, permite recibir un número 'n' de argumentos
func sum(nums ...int) (total int) {
	for _, num := range nums {
		total += num
	}
	return
}
