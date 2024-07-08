package main

import "fmt"

func main() {
	division(100, 10)
	division(200, 25)
	division(34, 0)
	division(124, 8)
}

func division(dividend, divisor int) {
	defer func() {
		// Con la función recover nos podemos recuperar de un pánico.
		// Esta función devuelve un valor, que nos permite a nosotros saber si en efecto se generó un pánico
		if r := recover(); r != nil {
			fmt.Println("Me recuperé del pánico 😮‍💨")
		}
	}()

	validateZero(divisor)
	fmt.Println(dividend / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("🤕 No puedes dividir por cero.")
		// Cuando ocurre un pánico, el programa se detiene y GO nos otorgará la traza
		// Pero nosotros podemos recuperarnos de un pánico (ver línea 14)
	}
}
