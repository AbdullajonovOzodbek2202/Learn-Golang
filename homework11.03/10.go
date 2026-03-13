package main

import "fmt"

func main() {
	var arr []int = []int{1, 78, 77, 99, 56, 44, 99, 121, 191, 123, 88}
	var polindromes []int

	for _, v := range arr {
		rev := 0
		y := v
		for y > 0 {
			rev = rev*10 + y%10
			y = y / 10
		}
		if rev == v {
			polindromes = append(polindromes, rev)
		}
	}
	fmt.Println(polindromes)
}
