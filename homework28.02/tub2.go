package main

import "fmt"

func allPrime(num1, num2 int) {
	for i := num1; i < num2; i++ {
		sana := 0
		for j := 1; j <= i/2; j++ {
			if i%j == 0 {
				sana++
			}
		}
		if sana == 1 {
			fmt.Println(i)
		}
	}
}

func main() {
	var num1, num2 int
	fmt.Print("Son kiriting: ")
	fmt.Scan(&num1)
	fmt.Print("2-sonni kiriting: ")
	fmt.Scan(&num2)

	allPrime(num1, num2)

}
