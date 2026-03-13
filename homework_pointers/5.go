package main

import "fmt"

func main() {
	var a = 54
	var x int
	ikkixonali(a, &x)
	fmt.Println(x)
}
func ikkixonali(a int, x *int) {
	*x = a/10 + a%10
}
