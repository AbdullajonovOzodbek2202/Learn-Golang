package main

import "fmt"

func main() {
	arr := []int{1, 7, 4, 3, 2, 0, 6, 5, 8, 10, 9, 12}

	sum := 0
	n := len(arr)

	for _, v := range arr {
		sum += v
	}

	total := n * (n + 1) / 2

	fmt.Println(total - sum)

}
