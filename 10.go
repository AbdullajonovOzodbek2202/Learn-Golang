package main

import (
	"fmt"
	"strings"
)

func sana(soz,x string ) int{

	sana := 0

	if soz == x {
		sana ++
	}

	return sana 
}

func main() {

	gap := "men golang dasturlash tilini o'rganyapman.BU dasturlash tili juda oson,tili"
	x := "tili"

	words := strings.FieldsFunc(gap, func(r rune) bool {
		return r == ' ' || r == ',' || r == '.'
	})

	res := 0

	for _, v := range words {
		res += sana(v,x)
	}

	fmt.Println("Berilgan soz:", res)

}