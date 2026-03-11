package main

import "fmt"

func main(){
	var arr = [10]int{1,-6,7,9,5,7,3,-10,23,10}
	fmt.Println(arr)

	var sum = 0

	for _, v :=  range arr {
		sum += v
	}

	fmt.Println(sum)
}