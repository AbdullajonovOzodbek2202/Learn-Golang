package main

import "fmt"

func main() {
	var a = 78
	var x int

	ikkialmashtir(a, &x)
	fmt.Println(x)

}
func ikkialmashtir(a int, x *int) {
	*x = a%10*10 + a/10
}
