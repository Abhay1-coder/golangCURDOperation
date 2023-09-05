package main

import "fmt"

func main() {
	fmt.Println("welcom to struct tutorial")

	abhay := User{"abhay", "abhay@go.com", true, 24}
	fmt.Println("abhay details", abhay)           //output: abhay details {abhay abhay@go.com true 24}
	fmt.Printf("abhay details are %+v \n", abhay) //output: abhay details are {name:abhay email:abhay@go.com status:true age:24}
	fmt.Printf("name is %v \n and email is %v\n", abhay.Name, abhay.Email)
	//output is
	// 	name is abhay
	//  and email is abhay@go.com

	//here is how we call our mthod
	abhay.GetStatus()

	//call method which manupulate this data
	//this will not omit original value of email.
	abhay.newEmail()
	//you can see original value is not change
	fmt.Printf("name is %v \n and email is %v\n", abhay.Name, abhay.Email)

}

//this is how struct is defined
//struct is similar to class but dont have inheritance , no supper keyword
type User struct {
	//Start with capital letter so that it can be defined as public and can be exported
	Name   string
	Email  string
	status bool
	Age    int
}

//this is how mothod is defined
func (u User) GetStatus() {
	fmt.Println("is user is active :", u.status)
}

func (u User) newEmail() {
	//this will not omit original value of email.  since we are creating object of User. its copy is being created. once we change value in copy of object. original value is not change
	u.Email = "test@go.dev"
	fmt.Println("the new email is :", u.Email)
}
