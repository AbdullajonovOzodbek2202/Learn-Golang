package main

import (
	"fmt"
	"errors"
)

func main() {
	var res float64 = 0
	var i int = 1
	var operator string
	var a, b float64
	for true {
		a = res
		if i != 0 {
			fmt.Print("Son kiriting: ")
			fmt.Scan(&a)
		}

		fmt.Print("Amal kiriting:")
		fmt.Scan(&operator)

		if operator == "=" {
			fmt.Println(res)
			return
		}

		if operator != "+" && operator != "-" && operator != "*" && operator != "/" && operator != "=" && operator != "%" {
			fmt.Println("Notog'ri amal kiritdingiz!")
			return
		}

		fmt.Print("Son kiriting: ")
		fmt.Scan(&b)

		switch operator {
		case "+":
			res = add(a, b)
		case "-":
			res = sub(a, b)
		case "*":
			res = mul(a, b)
		case "/":
			a,err := divide(a, b)
			if err != nil {
				fmt.Println(err)
				return
			}
			res = a
		case "%":
			res = percent(a, b)
		case "=":
			fmt.Println(res)
			return
		}

		i = 0
	}

}

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func mul(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64,error) {
	if b == 0 {
		err := errors.New("Sonni 0 ga bo'lib bo'lmaydi")
		return a/b,err
	}
	return a / b, nil
}

func percent(a, b float64) float64 {
	return a * b / 100
}

func root(a float64) float64 {
	return a
}
