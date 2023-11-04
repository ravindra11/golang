package main

// Declare a main package (a package is a way to group functions, and its made up of all the files in the same directory)

import (
	"fmt"
	"log"

	"rsc.io/quote"

	"example.com/greetings"

	"example.com/randomStr"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	goQuote := quote.Go()
	hello := quote.Hello()
	greet, _ := greetings.Hello("Ravindra")
	randomString, _ := randomStr.RandomString("Ravindra")
	fmt.Println(goQuote)
	fmt.Println(hello)
	fmt.Println("greeting module --- ", greet)
	fmt.Println("randome greeting --- ", randomString)
	multipleGreetings, error1 := randomStr.MultipleGreetings([]string{"Ravindra", "Babu"})
	fmt.Println("randome greetings to multiple people --- ", multipleGreetings)
	if error1 != nil {
		log.Fatal("emty names ", error1)
	}

	fmt.Println("****************************************")
	_, err := greetings.Hello(" ")
	if err != nil {
		log.Fatal("name empty ", err)
	}
}
