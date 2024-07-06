package main

import "fmt"

func main() {
	// Para declarar un array debemos hacerlo de la siguiente manera:
	/*
		var flags [3]string
	*/
	// Los [] es para definir el tamaño que tendrá el array, en este caso 3
	// y seguido a ellos debemos poner el tipo de dato que manejará ese array, en este caso strings
	// Al GO ser un lenguaje tipado solo nos permitirá almacenar strings y NO cualquier otro tipo de dato

	// Para agregar elementos:
	/*
			flags[0] = "🇲🇽"
			flags[1] = "🇺🇸"
			flags[2] = "🇨🇴"
			//flags[3] = "🇨🇴" -> Esto daría un error
			// IMPORTANTE:
		En Go los arrays son de tamaño fijo y esto no puede cambiar
	*/

	// podemos usar otra manera -> Arrays literales, sirven para asignar como para declarar
	/*
		flags := [3]string{"🇲🇽", "🇺🇸", "🇨🇴"} → Esto sigue siendo de tamaño fijo
	*/
	// Pero podemos hacer que GO lo infiera por nosotros, con base en los valores que le pasemos
	flags := [...]string{"🇲🇽", "🇺🇸", "🇨🇴", "🏴󠁧󠁢󠁷󠁬󠁳󠁿", "🇷🇴"} // Sin embargo, sigue siendo de tamaño fijo

	fmt.Println(flags)
}
