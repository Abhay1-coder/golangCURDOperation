package main

import "fmt"

func main() {
	//when we denote anything with the defer keyword. it will be placed at the last of the program
	//but what if there is more than two defer. yes, it will be executed in LPFO concept. example is below
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	defer fmt.Println("world")
	fmt.Println("hello")

	add()
	//output
	// 	hello
	// world
	// three
	// two
	// one

}

//lets understand this with function
func add() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

//output
// hello
// 4
// 3
// 2
// 1
// 0
// world
// three
// two
// one
