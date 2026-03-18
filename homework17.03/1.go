package main

import "fmt"

func main() {
	var soz string = "hello world"

	a := soz[0]
	a = a - 32

	soz = string(a) + soz[1:len(soz)]

	fmt.Println(soz)

}