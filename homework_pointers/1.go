package main 

import "fmt"

func main() {
	var a,b = 4,5
	var p , y int
	tortburchak(a, b, &p, &y)

	fmt.Println(p,y)

}

func tortburchak(a,b int, p, y *int) {
	*y = a * b
	*p = 2 * (a + b)

}