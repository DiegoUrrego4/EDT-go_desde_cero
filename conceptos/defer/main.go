package main

import (
	"fmt"
	"os"
)

// Cuando se difieren varias funciones los que hace GO es inyectarlas en una pila
// En este caso específico pasaría de la siguiente manera
/*
	pila = [
		defer fmt.Println(1)
		defer fmt.Println(2)
		defer fmt.Println(3)
		]
 Una vez armada la pila, GO espera a que la función principal termine, una vez ocurrido esto GO empieza a sacar
las funciones de la pila en el orden contrario, es decir, la primera en entrar es la última en salir y viceversa
*/

func main() {
	// Nos permite diferir (aplazar) algo
	defer fmt.Println(3) // la primera entra a la 'pila'
	defer fmt.Println(2) // luego entra esta, pero queda encima de la anterior (ver línea 8)
	defer fmt.Println(1) // y luego entra esta, que queda encima de las dos anteriores

	// Impresión en pantalla: 1,2,3

	num := 4
	defer fmt.Println("imprime", num)

	num = 10
	fmt.Println(num)

	// Pero en la práctica… ¿Esto para qué nos sirve?
	// Limpiar recursos, cerrar archivos, cerrar conexiones de red o cerrar controladores de BB DD

	// Ejemplo:
	file, err := os.Create("conceptos/defer/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("Hello gophers"))
	if err != nil {
		// file.Close()  Estaríamos repitiendo código!
		fmt.Println(err)
		return
	}

}
