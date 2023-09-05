package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) { // this is the mothod to respond on get request
	w.Write([]byte("<h1>Welcome to golang series on YT</h1>")) //here we are writting response what will be printed out when get request is hitted, response is always byte
}

//some usefull commands
//go mod init github.com/abhay1-coder/mymodule
//go get -u github.com/gorilla/mux // to install gorilla/mux
//go build . // this is use to make exe file
//go mod verify //this is use to verify all the modubles and version
//go mod vendor // it is use to make ready your project to sent in production
