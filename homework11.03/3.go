package main

import "fmt"

func main() {
	var arr []int = []int{1, 3, 4, 6, 78, 98, 43, 23, 12, 43, 5, 87}

	var max = arr[0]
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	fmt.Println(max)

}
