package main

import "fmt"

func main() {
	// slice: son apuntadores a un array, no poseen datos
	things := [7]string{"🚓", "🚙", "🚕", "🏎️", "🚒", "🚨", "🎈"}

	// Para crear un slice debemos indicar a donde apuntaremos
	// Slice N.º 1
	// El primer índice es incluyente y el último excluyente [inicial:final]
	// No es necesario indicar el índice inicial si nos interesa desde el índice 0
	cars := things[:5] // "🚓", "🚙", "🚕", "🏎️", "🚒"

	// Slice N.º 2
	// Si dejamos el último índice sin llenar por defecto GO inferirá que queremos el último índice
	redThings := things[3:] // "🏎️", "🚒", "🚨", "🎈"
	// Para cambiar elementos
	redThings[1] = "🚑" // con esto cambiaremos esto "🚒" → "🚑" SE MODIFICAN TODOS LOS ARRAYS

	// Entonces
	fmt.Println("Things", things)        // "🚓", "🚙", "🚕", "🏎️", "🚑", "🚨", "🎈"
	fmt.Println("Cars", cars)            // "🚓", "🚙", "🚕", "🏎️", "🚑"
	fmt.Println("Red Things", redThings) // "🏎️", "🚑", "🚨", "🎈"
	// Fijémonos en que el cambio de la ambulancia se cambió en todos los arrays/slices puesto, que los slices SON APUNTADORES

	// ¿Cómo accedemos a los elementos de los slices?
	fmt.Println("redThings[0]", redThings[0]) // "🏎️"

	// Tamaño y capacidad
	// len(): # de elementos en el slice
	// cap(): # de elementos del array origen, a partir del índice donde se creó el slice

	animals := [5]string{"🦍", "🐈", "🦮", "🦜", "🐘"}
	pets := animals[1:3]

	// Agregar elementos a los slices
	// Se debe hacer con la función pre construida append
	pets = append(pets, "🐈‍⬛", "🐕", "😸") // Acá se modifica el elemento original, en este caso "🦜" → "🐈‍⬛" y "🐘" → "🐕"
	// pero entonces, que pasa cuando ya el array NO tiene más capacidad y el array original está de la siguiente forma:
	// ["🦍", "🐈", "🦮", "🐈‍⬛", "🐕"]
	// Pues lo que pasaría es lo siguiente:
	// La capacidad inicial de pets era 4, una vez la capacidad máxima es alcanzada, GO crea un nuevo array y aumenta la capacidad al doble
	// En este caso 4 → 8
	// Por este motivo el array original queda intacto, ya que la función append devuelve un nuevo array y queda apuntando a ese

	fmt.Println("animals", animals) // Array[4]["🦍", "🐈", "🦮", "🦜", "🐘"]
	fmt.Println("pets", pets)       // new Array[8]{"🐈", "🦮", "🐈‍⬛", "🐕", "😸"}
	fmt.Println("Tamaño pets", len(pets))
	fmt.Println("Capacidad pets", cap(pets))

	// Pero también podemos crear slices sin necesidad de tener un array existente
	// Forma N.º 1
	/*
		sports := []string{"🏃‍", "🧗🏻"}
	*/

	// Forma N.º 2
	sports := make([]string, 0, 3)
	// make(tipo, tamaño, capacidad) entonces en este caso, el tamaño de sports sería 0 y tendría una capacidad inicial de 3
	sports = append(sports, "🏃‍", "🧗🏻", "🏊🏼‍", "🏋🏻")
	// Al agregar 4 elementos entonces lo que pasaría con mi slice es:
	// El tamaño aumentaría a 4 y la capacidad al doble de la inicial, es decir, de 3 pasaría a ser 6

	fmt.Println("sports", sports)                // "🏃‍", "🧗🏻", "🏊🏼‍", "🏋🏻"
	fmt.Println("Tamaño sports", len(sports))    // 4
	fmt.Println("Capacidad sports", cap(sports)) // 6
	fmt.Println("Valor cero", sports == nil)     // Esto si no se pasaran llaves '{}' si se pasan NO será nil

}
