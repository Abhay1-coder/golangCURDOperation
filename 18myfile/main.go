package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to file tutorial in go")

	content := "This text should go in a file."

	file, err := os.Create("./myfile.txt") //this will create a file in same directory with name myfile.txt

	// if err != nil { //find if creation of file through any error
	// 	panic(err) // this will stop imidatly when any eror is encounter
	// }
	checkNilError(err)

	len, err := io.WriteString(file, content) //this is storing content into a file. this will return length and err

	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)
	fmt.Println("length of the text in a file is", len)
	defer file.Close()
	readFile("./myfile.txt") // function call
}

//lets create a function where we will read our file.
//rember we have another util to read file. Os is use for creating and io is use for reading

func readFile(file string) {
	gotBytedata, err := ioutil.ReadFile(file)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)
	fmt.Println("this is data in our file", gotBytedata)                              //it will return byte data
	fmt.Println("Here we are converting byte data into string:", string(gotBytedata)) //byte to string conversion
}

// handling error in simple way
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
