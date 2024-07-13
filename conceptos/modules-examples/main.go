package main

import (
	"conceptos/modules-examples/slices"
	"fmt"
	"rsc.io/quote/v3"
	"strings"
)

func main() {
	list := []string{"EDteam", "gophers", "golang", quote.HelloV3()}
	slices.Includes(list, "gophers")
	slices.Filter(list, func(item string) bool {
		return strings.HasPrefix(strings.ToLower(item), "g")
	})

	fmt.Println(quote.Concurrency())
}
