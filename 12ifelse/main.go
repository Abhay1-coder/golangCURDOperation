package main

import "fmt"

func main() {
	fmt.Println("welcome to ifesle tutorial in golang")

	var logincount int = 23
	var result string

	if logincount < 10 {

		result = "result is less than 10"

	} else if logincount > 10 {
		result = "count is greater than 10"
	} else {
		result = "count is exactly 10"
	}

	fmt.Println(result)

	if num := 20; num < 10 {
		fmt.Println("number is less than 10")
	} else {
		fmt.Println("number is greater than 10")
	}

}
