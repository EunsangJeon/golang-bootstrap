package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type koreanBot struct{}

func main() {
	eb := englishBot{}
	kb := koreanBot{}

	printGreeting(eb)
	printGreeting(kb)
}

func (eb englishBot) getGreeting() string {
	// Assume that custom logic is here
	return "Hello!"
}

func (kb koreanBot) getGreeting() string {
	// Assume that custom logic is here
	return "안녕!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
