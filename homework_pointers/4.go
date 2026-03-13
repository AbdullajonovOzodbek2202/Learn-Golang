package main

import "fmt"

func main() {

	var a, b int = 5, 7
	fmt.Println(a, b)
	var aa, bb *int = &a, &b

	c := *aa
	*aa = *bb
	*bb = c

	fmt.Println(a, b)
}
