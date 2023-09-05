package main

import "fmt"

func main() {
	fmt.Println("welcome to funtion tutorial in golang")

	//function call
	greetng()

	result := addition(4, 9)
	fmt.Println("result of additions", result)

	//calling advance addtions method
	//this is how we catch both return type of a method
	prores, myMessage := advanceAdditions(2, 3, 4, 5, 6, 6, 4, 3, 2)
	fmt.Println("pro result is ", prores)
	fmt.Println(myMessage)

}

//function should be decleared outside the main function and called in main function
//this is how function is defined in golang.  int is return type
func addition(a int, b int) int {
	return a + b
}

//here is no return type because we are returning nothing.
func greetng() {
	fmt.Println("namastey inida")
}

//In this method we are passing slice of value. which will store n number of value. all the value will be of integer type
//here we can return two types of data ex below
func advanceAdditions(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "this is how we can print message along with integer data"
}
