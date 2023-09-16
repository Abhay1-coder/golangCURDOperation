package main

import (
	"fmt"
	"time"
)

func main() {
	//c1 channel id define to recive data from go function
	c1 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second) //this will wait for 2 second and result -1 will be stored to c1
		c1 <- "result-1"
	}()

	select { //select function is use either for result-1 or timeout-1
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout-1")
	}

	c2 := make(chan string)
	//c2 channel id define to recive data from go function
	go func() {
		time.Sleep(1 * time.Second) //this will wait for 1 second and result -2 will be stored to c2
		c2 <- "result-2"
	}()

	select { //select function is use either for result-2 or timeout-2
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
