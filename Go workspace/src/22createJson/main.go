package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tag      []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to json creation in go lang")
	encodeJson()
}

func encodeJson() {
	lcoCourse := []courses{
		{"react bootcamp", 259, "learncodeline.in", "abhcd123", []string{"wev-dev", "js"}},
		{"nodejs bootcamp", 599, "learncodeline.in", "abhcd123", []string{"server-dev", "js"}},
		{"angular bootcamp", 299, "learncodeline.in", "abhcd123", nil},
	}

	//package this data as a json data
	finalJson, err := json.MarshalIndent(lcoCourse, "abhay", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
