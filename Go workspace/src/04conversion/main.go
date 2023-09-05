package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "welcome to comversion"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter ratting for pizza")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your ratting:", input)
	fmt.Println("the type of ratting is %T", input)

	// numratting := input +1 // this will through error because  the type of input is string so we cannot add integer to it. To do so we have to do type conversion

	//type conversion
	// numratting, err := strconv.ParseFloat(input, 64)
	numratting, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // will remove space \n

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to your ratting :", numratting+1)
		//opps you got this error
		//strconv.ParseFloat: parsing "4\r\n": invalid syntax
		//now we have to trim \r\n

	}

}
