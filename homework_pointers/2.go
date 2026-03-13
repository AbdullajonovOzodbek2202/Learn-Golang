package main

import "fmt"

func main() {
	var r float64
	r = 5
	var l, s float64
	doira(r, &l, &s)

	fmt.Println(l, s)

}

func doira(r float64, l, s *float64) {
	*s = 3.14 * r * r
	*l = 3.14 * 2 * r
}
