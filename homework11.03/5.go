package main

import "fmt"

func main() {
	var arr []int = []int{1, 12, 32, 45, 65, 67, 5, 34, 36, 78, 91, 12}

	sana := 0

	for _, v := range arr {
		if v%2 != 0 {
			sana++
		}
	}
	fmt.Println(sana)

}
