package main

import (
	"fmt"
	"log"

	"example.local/greetings"
	"example.local/greetings/childGreetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// names := []string{"Gladys", "Samantha", "Darrin"}
	// message, err := greetings.Hellos(names)
	// message, err := greetings.Hello("Mr TuanNH")
	message, err := greetings.Hello("Mr TuanNH")
	message2, _ := childGreetings.Hello2("Mr TuanNH")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
	fmt.Println(message2)
}
