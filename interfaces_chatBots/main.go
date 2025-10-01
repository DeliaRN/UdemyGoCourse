package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
	// Needs a getGreeting() - returns Hi
	// Needs a printGreeting() - probably same logic
	// fmt.Println(eb.getGreeting())
}

type spanishBot struct {
	// Needs a getGreeting() - returns Hola
	// Needs a printGreeting() - probably same logic
	// fmt.Println(sb.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// This kind of overload by arguments is not possible in Go.
/*
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}*/

func (englishBot) getGreeting() string {
	//ommit the value 'eb' since it's not being used
	// VERY custom logic for English blablabla
	return "Hi!"
}

func (sb spanishBot) getGreeting() string {
	// VERY custom logic for Spanish blablabla
	return "Hola!"
}
