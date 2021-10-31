package main

import (
	"fmt"
	"log"
)

type User struct {
	FirstName   string
	PhoneNumber string
}

func (user *User) printFirstName() string {
	return user.FirstName
}

// interface
type Animal interface {
	Says() string
}

func PrintInfo(a Animal) {
	fmt.Println(a.Says())
}

type Dog struct {
}

func (d *Dog) Says() string {
	return "Woooof"
}

type Gorlla struct {
}

func (g *Gorlla) Says() string {
	return "Ooof"
}

func main() {

	// fmt
	whatWasSiad, theOther := say()
	fmt.Println(whatWasSiad, theOther)

	// Pointer
	var myString = "Green"
	changePointer(&myString)
	log.Println(myString)

	// struct
	user := User{
		FirstName:   "Trevor",
		PhoneNumber: "000-9999-0000",
	}
	log.Println(user.FirstName, "PhoneNumber:", user.PhoneNumber)
	log.Println(user.printFirstName())

	// map
	myMap := make(map[string]string)
	myMap["dog"] = "Samson"
	// slices
	numbers := []int{1, 2, 3, 4, 5}
	log.Println(numbers)
	log.Println(numbers[0:2])

	// loops
	for i, number := range numbers {
		log.Println(i, number)
	}

	// interface
	dog := Dog{}
	PrintInfo(&dog)
}

func say() (string, string) {
	return "sa", "else"
}

func changePointer(s *string) {
	newValue := "Red"
	*s = newValue
}
