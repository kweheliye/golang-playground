package main

import "fmt"

// Speaker interface
type Speaker interface {
	Speak() string
}

// Dog struct implementing Speaker
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says Woof!"
}

// Cat struct implementing Speaker
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return c.Name + " says Meow!"
}

// Robot struct implementing Speaker
type Robot struct {
	Model string
}

func (r Robot) Speak() string {
	return "Beep Boop! I am " + r.Model
}

// Function to make any Speaker speak (Polymorphism in action)
func LetThemSpeak(speakers []Speaker) {
	for _, s := range speakers {
		fmt.Println(s.Speak()) // Calls the Speak() method dynamically
	}
}

// SpeakerDemo function to demonstrate polymorphism
func RunMainInterfaceExample1() {
	// Create multiple Speaker instances
	fmt.Println("*******----Start Interface Speaker Examples----***********")
	speakers := []Speaker{
		Dog{Name: "Buddy"},
		Cat{Name: "Whiskers"},
		Robot{Model: "XJ-9"},
	}

	// Let all speakers speak using polymorphism
	LetThemSpeak(speakers)

	fmt.Println("*******----End Interface Speaker Examples----***********")

}
