package main

import "fmt"

func main() {
	// GO solo cuenta con la declaración for

	// Forma N.º1 - For clásico
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Forma N.º2 - For while
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// Forma N.º3 - For forever
	for {
		if i == 6 {
			break // La palabra break se usa para salir del bucle for
		}
		fmt.Println(i)
		i++
	}

	// Forma N.º4 - For each (range)
	// Slices
	food := []string{"🍕", "🍔", "🍣", "🥙"}
	for i, v := range food {
		fmt.Println("índice:", i, "Valor:", v)
	}

	numbers := []int{2, 4, 6, 8}
	for i := range numbers {
		// Cuando estamos trabajando con slice Go maneja copias
		numbers[i] *= 2
	}
	fmt.Println(numbers)

	// Mapas
	foodTwo := map[string]string{
		"pizza":     "🍕",
		"hamburger": "🍔",
		"sushi":     "🍣",
		"taco":      "🥙",
	}

	for key, value := range foodTwo {
		fmt.Println("Key", key, "Value", value)
	}

	// Strings
	for i, v := range "EDteam" {
		// v será un valor unicode
		//fmt.Println("Índice", i, "Value", v)
		// Si quisiéramos la letra sería castearlo
		fmt.Println("Índice", i, "Value", string(v))
	}
}
