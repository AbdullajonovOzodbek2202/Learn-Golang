package main

import "fmt"

func main() {
	soz := "kompyuter"
	unli := 0 

	for _,v := range soz {
		if v == 'a' || v == 'o' || v == 'e' || v == 'i' || v == 'u' || v == 'A' || v == 'O' || v == 'E' || v == 'I' || v == 'U'{
			unli ++
		}
	}
	fmt.Println("Unlilar: ",unli)
	undosh := len(soz) - unli
	fmt.Println("Undoshlar: ",undosh)
}