package main

import "fmt"

func main() {
	var arr []int = []int{100, 3, 5, 42, 2, 1, 3, 4, 5, 6, 6, -12, -23, 87, -27}

	var min = arr[0]

	for _, v := range arr {
		if min > v {
			min = v
		}
	}
	fmt.Println(min)

}
