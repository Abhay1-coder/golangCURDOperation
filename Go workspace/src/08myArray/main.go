package main

import "fmt"

func main() {
	welcome := "welcome to array tutorial"
	fmt.Println(welcome)

	//creation of an array
	var fruitsList [5]string
	fruitsList[0] = "apple"
	fruitsList[1] = "orange"
	fruitsList[2] = "banana"
	fruitsList[4] = "leamon"

	fmt.Println("list of fruits is:", fruitsList) //output : list of fruits is: [apple orange banana  leamon]
	//notice something in output. there id double speace between banana and leamon. this is becase index[3] is nil

	//now lets find the lenghth of the array

	var lenghOfArray int = len(fruitsList)
	fmt.Println("lenght of an array is :", lenghOfArray) //since u are giving size of array is 5. than length will be also 5 altough you have only 3 data in an array. than also length will be 5 only

	var vegList = [5]string{"potato", "onion", "beans"}
	fmt.Println("vey list is", vegList)
	fmt.Println("length of vegy list is", len(vegList))

}
