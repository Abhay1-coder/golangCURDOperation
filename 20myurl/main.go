package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ytgaakhgah"

func main() {
	fmt.Println("welcome to url tutorial in golang")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery) //output : https://lco.dev:3000/learn?coursename=reactjs&paymentid=ytgaakhgah
	// https
	// lco.dev:3000
	// /learn
	// 3000
	// coursename=reactjs&paymentid=ytgaakhgah

	//now lets find query in more prisce way
	qparams := result.Query()
	fmt.Printf("the type of our query param is: %T\n", qparams) //the type of our query param is: url.Values

	//since query param is in key value pair. lets find value of each key
	fmt.Println(qparams["coursename"]) //[reactjs]

	//let use forloop to find all the params in url
	for index, val := range qparams {
		fmt.Println("params is: ", index, val)
	}
	//or if you dont want index number. just want values
	for _, val := range qparams {
		fmt.Println("params is: ", val)
	}

	//lets construct our own url form small small part
	partsOfUrl := &url.URL{ //& is use for passing reference of url
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String() //convert url into a string
	fmt.Println(anotherUrl)

}
