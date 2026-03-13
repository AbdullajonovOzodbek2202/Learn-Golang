package main

import "fmt"

func main() {
	var l = 860
	var m int
	meters(l, &m)

	fmt.Println(m)

}

func meters(l int, m *int) {
	*m = l / 100
}
