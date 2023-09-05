package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to map tutorial")

	//lets define map
	languages := make(map[string]string)

	//lets add some data in map
	languages["jv"] = "java"
	languages["sp"] = "springBoot"
	languages["mg"] = "mongo"
	languages["js"] = "javaScript"

	fmt.Println("list of languages is:", languages)          //output :list of languages is: map[js:javaScript jv:java mg:mongo sp:springBoot]
	fmt.Println("one language from a list", languages["js"]) //output: one language from a list javaScript

	//lets delete one language
	delete(languages, "js")
	fmt.Println("list of languages after deletetion of js is:", languages)

	//loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("for key %v, value is %v \n", key, value) // %v is for value
		//output : for key jv, value is java
		// for key sp, value is springBoot
		// for key mg, value is mongo
	}

	//lets do this with comma okk syntax
	//in this we doesnt bother about key. we want only value, then
	for _, value := range languages {
		fmt.Printf("for key k, value is %v", value) // output : for key k, value is javafor key k, value is springBootfor key k, value is mongo
	}

}
