package main

import "fmt"

func main() {
	var arr []int = []int{1,34,54,67,45,3,7,89,90,4,32,12,4,}
	var sana = 0

	for _,v := range arr {
		if v % 2 == 0 {
			sana ++
		}
	}
	fmt.Println(sana)

}
