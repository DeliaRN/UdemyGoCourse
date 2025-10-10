package main

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

type germanGreeter struct{}

func SayHello(n string, g Greeter) string {
	language := g.LanguageName()
	greeting := g.Greet(n)
	return fmt.Sprintf("I can speak %s: %s", language, greeting)
}

func (g germanGreeter) Greet() string {
	name := g.LanguageName()
	return fmt.Sprintf("Hallo %s!", name)
}

func (g germanGreeter) LanguageName() string {
	return "German"
}
