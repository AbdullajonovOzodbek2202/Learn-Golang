package main

import "fmt"

func main() {
	gap := "Men golang dasturlash tilini o'rganyapman."

	sozlar := 0

	for _, v := range gap {
		if v == ' ' || v == '.' {
			sozlar++
		}
	}
	fmt.Println("Berilgan gapdaGI so'zlar soni: ", sozlar)
}
