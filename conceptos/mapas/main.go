package main

import "fmt"

func main() {
	// Mapas. Son estructuras llave - valor

	// CREAR MAPAS:
	// Forma N.º 1:
	// Para crear mapas podemos hacerlo con la función pre construida make
	// map[tipo de la llave] tipo del valor
	music := make(map[string]string)
	music["guitar"] = "🎸"
	music["violin"] = "🎻"

	fmt.Println(music)

	// Forma N.º 2:
	tech := map[string]string{
		"computer": "🖥️",
		"mouse":    "🖱️",
	}

	fmt.Println(tech)

	// ELIMINAR ELEMENTOS
	// delete(nombre del mapa, llave del elemento que queremos eliminar)
	delete(tech, "computer")

	fmt.Println(tech)

	// LEER VALORES DEL MAPA
	fmt.Println(music["fake"]) // Cuando no encuentra un valor, retornará vacío
	// Sin embargo, el mapa devuelve dos valores
	// contenidos (con el tipo que tenga), si existe o no (valor booleano)
	value, ok := music["violin"]
	fmt.Printf("Valor retornado: %v, Existe en el mapa: %t\n", value, ok)
}
