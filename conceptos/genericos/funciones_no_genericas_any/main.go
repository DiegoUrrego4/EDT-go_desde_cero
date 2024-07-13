package main

import "fmt"

func main() {
	PrintList("EDteam", "gopher", "Go desde cero")
	PrintList(1, 2, 3)
}

// any aceptaría CUALQUIER tipo de dato

func PrintList(list ...any) {
	for _, item := range list {
		fmt.Println(item)
	}
}
