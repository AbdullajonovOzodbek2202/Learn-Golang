package main

import "fmt"

func main() {
	var a = 1234
	var x int
	tortxonali(a,&x)
	fmt.Println(x)
	
}
func tortxonali(a int, x *int) {
	*x = (a%10 + a/100%10) - (a/10%10 + a/1000)
}