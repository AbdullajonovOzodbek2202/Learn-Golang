package main

import "fmt"

func main() {
	var a = 135
	var x int
	uchxonalialmashtir(a, &x)
	fmt.Println(x)
}
func uchxonalialmashtir(a int, x *int) {
	*x = a%10*100 + a/100 + a/10%10*10
}
