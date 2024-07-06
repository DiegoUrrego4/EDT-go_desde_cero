package main

import "fmt"

// Forma N.º 4 - Constantes a nivel paquete
/*
const os, domain = "linux", "ed.team"
*/

// Forma N.º 5 - Agrupamiento
const (
	os     = "linux"
	domain = "ed.team"
)

// Identificador iota
const (
	Jan = iota + 1
	Feb
	Mar
	Apr
	May
	Jun
)

func main() {
	/*
		Para declarar la variable se usa una keyword distinta
		En las constantes DEBEMOS inicializarlas de forma inmediata cuando las declaramos
		Este valor NO PUEDE modificarse luego de esta línea
		Las constantes NO NECESITAN SER USADAS para su compilación
	*/

	// Forma N.º 1 - una sola variable usando keyword y tipo
	/*
		const os, domain string = "linux", "ed.team"
	*/

	// Forma N.º 2 - varias constantes de una vez con tipo
	/*
		const os, domain string = "linux", "ed.team"
	*/

	// Forma N.º 3 - varias constantes de una vez sin tipo (inferido por GO)
	/*
		const os, domain = "linux", "ed.team"
	*/

	fmt.Println(os, domain, May)
}
