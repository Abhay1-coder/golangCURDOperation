package main

import "fmt"

func mypanic() {
	panic("a problem")
}
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered error:\n", r)
		}
	}()

	mypanic()
	fmt.Println("after mypanic()")
}
