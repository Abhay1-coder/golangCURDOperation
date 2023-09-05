package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices tutorial")

	//lets define slices
	//this is how we define slices in go.
	var fruitsList = []string{"apple", "orange", "mango"}
	fmt.Println(fruitsList) //output : [apple orange mango]

	//lets add some more fruits in slices
	fruitsList = append(fruitsList, "banana", "lemon")
	fmt.Println(fruitsList) //output :[apple orange mango banana lemon]

	//lets perform some operation in slices
	// i am creating new slice i.e newFruits, i dont have to mention this is slice. since i am storing value of slice data in a varible it will automaticlly take it as slice
	var newFruits = append(fruitsList[1:3]) //output : [orange mango]
	fmt.Println(newFruits)
	fmt.Println("size of newFruits", len(newFruits)) // output : 2, it is just like arrayList. size is not defined. it will take size on the basis of data avilable in the slice

	// fruitsList = append(fruitsList[1:3])
	// fruitsList = append(fruitsList[:3])                    // this will print from 0 to 3
	// fruitsList = append(fruitsList[1:])                    // this will print from index 1 to all
	// fruitsList = append(fruitsList[1 : len(fruitsList)-2]) // this will  print all the fruits except last two fruits in slice
	fmt.Println("fruits list after append", fruitsList)
	fmt.Println("size of fruitslist is", len(fruitsList))

	//lets play with make() in slice
	highscores := make([]int, 4) // make() will create memory size of 4
	//this is defult allocation of memeory, so we cannot store data more than size
	highscores[0] = 234
	highscores[1] = 456
	highscores[2] = 874
	highscores[3] = 678
	fmt.Println(" highscores data is", highscores) // this will print all the value in highscores

	//lets test, what happen when we store value which is out of range
	// highscores[4] = 03455
	// fmt.Println("print value of highscores which is exceding out of range", highscores)//output : runtime error: index out of range [4] with length 4

	//now lets  append value in highscores
	//when we use append method. it will reallocate the memory and we can store n nummber of data
	highscores = append(highscores, 432, 87654, 2345, 2345)
	fmt.Println("print value after appending value in highscores", highscores) //output : print value after appending value in highscores [234 456 874 678 432 87654 2345 2345]

	//lets play with sorting in slice. In array we cannot perform  this operations
	//lets sort our highscores
	sort.Ints(highscores)
	fmt.Println("sorted highscores is ", highscores)
	fmt.Println(sort.IntsAreSorted(highscores)) // return bool value that slice is sorrted or not

	//how to remove value from sclice
	var cources = []string{"java", "springBoot", "mongo", "heibernate", "sql"}
	fmt.Println(cources)
	//lets remove mongo from the cource
	index := 2
	cources = append(cources[:index], cources[index+1:]...) // ... dot means that after this print rest all
	fmt.Println(cources)                                    //output : [java springBoot heibernate sql]

}
