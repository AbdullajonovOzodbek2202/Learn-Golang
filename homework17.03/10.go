package main

import (
	"fmt"
	"strings"
)

func main() {
	gap := "men golang dasturlash tilini o'rganyapman.BU dasturlash tili juda oson,tili"
	x := "tili"
	sana := 0
	words := strings.FieldsFunc(gap, func(r rune) bool {
		return r == ' ' || r == ',' || r == '.'
	})

	for _, v := range words {
		if x == v {
			sana += 1
		}
	}
	fmt.Println("Berilgan soz:", sana)
}
