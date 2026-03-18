package main

import (
	"fmt"
	"strings"
)

func isPolindrom(soz string) bool {
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
	soz := "Aziza"
	soz1 := "maktab"

	res := isPolindrom(soz)
	res1 := isPolindrom(soz1)
	fmt.Println(res)
	fmt.Println(res1)
}
