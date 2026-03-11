package main

import "fmt"

func main() {
	var arr = [6]int{1, 5, 7, 8, 9, 5}

	fmt.Println(arr)

	length := len(arr)
	for i := length - 1; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}
}
