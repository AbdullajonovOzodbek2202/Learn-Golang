package main

import (
	"fmt"
	"strings"
)

func Polindrom(soz string) bool {

	res := ""
	i := len(soz)

	for j := i - 1; j >= 0; j-- {
		res += string(soz[j])
	}

	res = kichikHarf(res)
	soz = kichikHarf(soz)

	return soz == res
}

func kichikHarf(soz string) string {

	res := ""

	for _, v := range soz {

		if v >= 65 && v <= 90 {
			v = v + 32

			res = res + string(v)
		} else {
			res = res + string(v)
		}
	}
	return res
}

func main() {
	gap := "Aziza maktab ana oy"

	words := strings.Split(gap, " ")

	for _, v := range words {
		x := Polindrom(v)
		if x == true {
			fmt.Println(v)
		}
	}
}
