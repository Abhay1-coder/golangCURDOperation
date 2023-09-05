package main

import (
	"crypto/rand"
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to switch case tutorial of golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("your dicenumber is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("you got 1 and you can open")
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move 3 spot")
	case 4:
		fmt.Println("you can move 4 spot")
	case 5:
		fmt.Println("you can move 5 spot")
	case 6:
		fmt.Println("you can move 6 spot")
	default:
		fmt.Println("what is this!")
	}
}
