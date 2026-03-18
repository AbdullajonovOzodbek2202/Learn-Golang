package main

import "fmt"

func main() {
	soz := "Universal"
	bogin := 0

	for _, v := range soz {
		if v == 'a' || v == 'o' || v == 'e' || v == 'i' || v == 'u' || v == 'A' || v == 'O' || v == 'E' || v == 'I' || v == 'U' {
			bogin++
		}
	}
	fmt.Println(bogin)
}
