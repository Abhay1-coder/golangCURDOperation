package main

import "fmt"

// outside method we can not remove var to declear any data type
//  lname :="bijay kalwar"
// 	fmt.Println(lname)

//If variable name start with capital letter, it denotes that it is public
var Jwttoken = "akdjf;kajdfh;k"
var Loggedin = true

func main() {
	fmt.Println("hi i am abhay kalwar")

	//lets define data types
	var fname string = "abhay"
	fmt.Println(fname)
	fmt.Printf("ok my fname data type is: %T \n", fname)

	//lets give boolean data types

	var isLoggedin bool = true
	fmt.Println("abhay is loggedin to system:", isLoggedin)

	var count int = 34567890987654
	fmt.Println(count)

	// lets remove data type decleartion
	var name = "neha sriwastav"
	fmt.Println(name)

	var countNumber = 6789876
	fmt.Println(countNumber)

	//we can also remove var but only inside method. outside method we can not remove var
	lname := "bijay kalwar"
	fmt.Println(lname)
}
