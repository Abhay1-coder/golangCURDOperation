package main

import (
	"fmt"
	"time"
)

func main() {
	welcome := "welcome to date and time"
	fmt.Println(welcome)

	presentTime := time.Now() //output: 2023-08-31 23:55:28.226421 +0530 IST m=+0.002347601
	fmt.Println(presentTime)

	//lets formate this
	//you might be thinking why "01-02-2006", then it is predefined by golang developer we have to use this only
	fmt.Println(presentTime.Format("01-02-2006")) //output: 08-31-2023

	//and this this is the pre determined for date time and day
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	//now lets create our own date
	createDate := time.Date(2020, time.August, 26, 12, 45, 0, 0, time.UTC) // this is all parameter you are suppose to pass to create day.. dont worry when you give time.Date. it will only guid you how to do it
	fmt.Println(createDate.Format("01-02-2006 Monday"))

}
