package main

import "fmt"

func isPrime(num int) {
	for i := 1; i < num; i++ {
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
	var num int
	fmt.Print("Son kiriting: ")
	fmt.Scan(&num)

	isPrime(num)

}
