package main

import "fmt"

func main() {
	welcome := "welcome to pointer class"
	fmt.Println(welcome)

	//what is pointer?
	//ans : When we declear any varibale and store any data there and if we are calling that variabe in different method,
	// at that time it will pass reference value of actul memory address. Pointer dont we that, pointer will pass
	// what exactly in that memory address. it will directly take data from memorry address rather than reference value.

	//pointer decleartion
	var ptr *int
	fmt.Println(ptr) // since there is no data it is showing nill

	mynumber := 23

	var newptr = &mynumber                    //& infront of any variable that will pass reference address
	fmt.Println("value of mynumber", newptr)  //output: value of mynumber 0xc000016088// this is reference address
	fmt.Println("value of mynumber", *newptr) //outputn: value of mynumber 23 //* will point to actul value

	*newptr = *newptr + 2
	fmt.Println(newptr) // when we perform any operation on pointer. think our are performing that operation on actul value not with reference
}
