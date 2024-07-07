package main

import "fmt"

func main() {
	// Sentencia if - else
	// Esta declaración de variable dentro del if se hace cuando solo necesitemos la variable dentro del if
	// y no en otro lugar del código
	if character := "🦹🏻"; character == "🦸🏻" {
		fmt.Println("Es un super héroe")
	} else if character == "🦹🏻" {
		fmt.Println("Es un super villano")
	} else {
		fmt.Println("Es un personaje normal")
	}

	// Sentencia switch
	// No es necesario colocar el break, ya que Go se encarga de esto
	canSearch := true
	character := "🧟"
	switch {
	case !canSearch:
		fmt.Println("No está permitida la búsqueda")
	case character == "🦸🏻" || character == "🧞‍": // Podemos agrupar casos separándolos con ','
		fmt.Println("Es un super héroe")
	case character == "🦹🏻" || character == "🧟":
		fmt.Println("Es un super villano")
	default:
		fmt.Println("Es un personaje normal")
	}

}
