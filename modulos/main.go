package main

import (
	"fmt"
	"log"

	"github.com/idanielmeza/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Daniel")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

}
