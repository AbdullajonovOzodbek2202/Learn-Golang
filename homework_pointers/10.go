package main

import "fmt"

func main() {
	var a = 3978
	var x int
	teskari(a, &x)
	fmt.Println(x)
}
func teskari(a int, x *int) {
	z := 0
	for a > 0 {
		z = z*10 + a%10
		a = a / 10
	}
	*x = z
}
