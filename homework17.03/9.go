package main

import (
	"fmt"
	"strings"
)

func main() {
	gap := "men golang dasturlash tilini o'rganyapman.BU dasturlash tili juda oson"
	x := "tili"
	words := strings.FieldsFunc(gap, func(r rune) bool {
		return r == ' ' || r == ',' || r == '.'
	})

	for _, v := range words {
		if x == v {
			fmt.Println("Qidirilgan soz bor: ", v)
		}
	}
}
