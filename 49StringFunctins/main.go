package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "Hello, World!"
	length := len(str) // length is 13
	fmt.Println(length)

	str1 := "Hello, "
	str2 := "World!"
	result := str1 + str2 // result is "Hello, World!"
	fmt.Println(result)
	equal := (str1 == str2) // equal is false
	fmt.Println(equal)

	str3 := "Golang"
	char := str3[0] // char is 'G'
	fmt.Println(char)

	str4 := "Hello, World!"
	sub := str4[7:12] // sub is "World"
	fmt.Println(sub)

	str5 := "Hello, World!"
	contains := strings.Contains(str5, "World") // contains is true
	fmt.Println(contains)

	str6 := "apple,banana,grape"
	parts := strings.Split(str6, ",") // parts is []string{"apple", "banana", "grape"}
	fmt.Println(parts)

	str7 := "Hello, World!"
	replaced := strings.Replace(str7, "World", "Golang", -1) // replaced is "Hello, Golang!"
	fmt.Println(replaced)

	str8 := "   Trim me    "
	trimmed := strings.TrimSpace(str8) // trimmed is "Trim me"
	fmt.Println(trimmed)

	num := 42
	str9 := strconv.Itoa(num) // str is "42"
	fmt.Println(str9)

}
