package main

import "fmt"

func main() {
	var arr []int = []int{1, 34, 54, 67, 45, 3, 7, 89, 90, 4, 32, 12, 4, 21, 45, 34, 76, 54, 0}

	var juftlar []int
	var toqlar []int

	for _, v := range arr {
		if v%2 == 0 {
			juftlar = append(juftlar, v)
		} else {
			toqlar = append(toqlar, v)
		}
	}
	fmt.Println(juftlar)
	fmt.Println(toqlar)

}
