package main

import (
	"fmt"
)

// Speaker is an interface with a Speak method

type Speaker interface {
	Speak() string
}

// Dog is a type that implements the Speaker interface
type Dog struct{}

// Speak is a method of the Dog type, satisfying the Speaker interface
func (d Dog) Speak() string {
	return "Woof!"
}

func main() {
	// Create a variable of type Speaker and assign it a Dog
	var s Speaker = Dog{}

	if dog, ok := s.(Dog); ok {
		fmt.Println("It's a dog", dog.Speak())
	} else {
		// Type assertion failed
		fmt.Println("It's not a dog!")
	}
}
