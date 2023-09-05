package main

import "fmt"

func main() {
	fmt.Println("welcome to loop in golang")

	var days = []string{"sunday", "monday", "tuesday", "thusday", "friday"}

	//1 way...let use for loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	//lets use another way of using for loop in go
	for i := range days { //remmber i here give index number not a value
		fmt.Println(days[i])
	}

	//lets implement foreach loop in golang

	for index, day := range days {
		fmt.Printf("index is %v and day is %v\n", index, day)
	}

	//lets understand another form of for loop. here we will understand break and continue

	number := 1
	// for number < 10 {
	// 	if number == 3 {
	// 		break
	// 	}
	// 	fmt.Println(number)
	// 	number++
	// }
	for number < 10 {
		if number == 2 {
			goto pic //when number is 2. it will go to pic and print what is there
		}
		if number == 3 {
			number++
			continue //before continue we have to do increment in golang
		}
		fmt.Println(number)
		number++
	}

	//let understand goto in golang
pic:
	fmt.Println("i am goto coming from pic defined below")

}
