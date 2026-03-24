package main

import "fmt"

func sozSana(gap string) (int) {

	var res = 0 

	for _,v := range gap {
		if v == ' ' || v == '.' {
			res ++
		}
	}

	return res

}

func main() {

	var gap = "Men golang dastrulash tilini o'rganyapman."

	res := sozSana(gap) 

	fmt.Println(res)
}