package main

import (

	"fmt"
	"strings"
)

func qidir(soz , x string) bool {

	if soz == x {
		return true
	}

	

	return false

}


func main() {

	gap := "men golang dasturlash tilini o'rganyapman.BU dasturlash tili juda oson"
	x := "tili"
	words := strings.FieldsFunc(gap, func(r rune) bool {
		return r == ' ' || r == ',' || r == '.'
	})

	for _, v := range words {

		y := qidir(v , x)

		if y == true{
			fmt.Println("Qidirilgan so'z bor")
			return
		}

	}

	fmt.Println("Qidirilgan so'z yo'q")

}
