package main

import "fmt"

func main() {
	soz := "hello WORLD"
	new := ""

	for _, v := range soz {

		if 65 <= v && v <= 90 {
			v = v + 32
			new += string(v)

		} else if 97 <= v && v <= 122 {
			v = v - 32
			new += string(v)

		} else {
			new += string(v)
		}
	}

	fmt.Println(new)
}
