package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("welcome to webrequst")

	response, err := http.Get(url)
	checkNilError(err)
	fmt.Printf("lets check what type of responce we are getting from get request %T\n", response) //output: *http.Response // this is reference address of actul data. i.e we are getting pointer of actul memory address

	defer response.Body.Close() // this is responsibility of coder to close the request at the end

	databytes, err := ioutil.ReadAll(response.Body) //store responce into databytes
	checkNilError(err)

	content := string(databytes) //conversion of byte to string
	fmt.Println(content)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
