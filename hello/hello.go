package hello

import (
	"fmt"
	"log"

	"github.com/mark1002/go_practice/greetings"
)

func Execute() {
	// Get a greeting message and print it.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
