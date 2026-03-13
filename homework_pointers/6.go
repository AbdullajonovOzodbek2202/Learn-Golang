package main

import "fmt"

func main() {
	var a = 147
	var x int
	uchxonali(a, &x)

	fmt.Println(x)

}
func uchxonali(a int, x *int) {
	*x = a/100 + a/10%10 + a%10
}
