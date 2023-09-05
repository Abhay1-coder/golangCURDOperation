package main

import "fmt"

func main() {
	fmt.Println("welcom to struct tutorial")

	abhay := User{"abhay", "abhay@go.com", true, 24}
	fmt.Println("abhay details", abhay)           //output: abhay details {abhay abhay@go.com true 24}
	fmt.Printf("abhay details are %+v \n", abhay) //output: abhay details are {name:abhay email:abhay@go.com status:true age:24}
	fmt.Printf("name is %v \n and email is %v", abhay.Name, abhay.Email)
	//output is
	// 	name is abhay
	//  and email is abhay@go.com

}

//this is how struct is defined
//struct is similar to class but dont have inheritance , no supper keyword
type User struct {
	//Start with capital letter so that it can be defined as public and can be exported
	Name   string
	Email  string
	Etatus bool
	Age    int
}
