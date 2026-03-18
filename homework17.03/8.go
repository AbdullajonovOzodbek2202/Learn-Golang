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
	res = strings.ToLower(res)
	soz = strings.ToLower(soz)
	return soz == res
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
