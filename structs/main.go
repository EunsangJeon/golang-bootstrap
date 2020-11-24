package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// john := person{"john", "lennon"}
	// john := person{firstName: "john", lastName: "lennon"}
	// var john person
	// john.firstName = "john"
	// john.lastName = "lennon"

	paul := person{
		firstName: "paul",
		lastName:  "mccartney",
		contactInfo: contactInfo{
			email:   "paul@beatles.com",
			zipCode: 12345,
		},
	}

	paulPointer := &paul
	paulPointer.updateName("polly")
	paul.print()
}

func (p person) print() {
	fmt.Println(p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
