package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type dutchBot struct{}

func main() {
	eb := englishBot{}
	db := dutchBot{}

	printGreeting(eb)
	printGreeting(db)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hi There!"
}

func (db dutchBot) getGreeting() string {
	return "Hoi"
}
