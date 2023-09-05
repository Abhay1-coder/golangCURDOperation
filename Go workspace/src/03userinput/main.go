package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter ratting for pizza")

	//comma ok || err ok
	input, _ := reader.ReadString('\n') //this line will take input
	fmt.Println("thanks for ratting: ", input)
	fmt.Printf("the type of input is %T", input) //find waht is the type of input is comming

}
