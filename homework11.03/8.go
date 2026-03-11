package main

import "fmt"

func main() {
	var arr []int = []int{1, -54, 6, 4, -3, 23, 3, -4, 56, 6, -67, 7 - 8, -89, 90, 9, -76, 213, 658}

	var manfiy []int
	var musbat []int

	for _, v := range arr {
		if v >= 0 {
			musbat = append(musbat, v)
		} else {
			manfiy = append(manfiy, v)
		}
	}
	
	fmt.Println(musbat)
	fmt.Println(manfiy)

}
