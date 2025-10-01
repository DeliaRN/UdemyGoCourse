package main

import "fmt"

func main() {

	//var colorMap1 map[string]string
	//colorMap2 := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// Adding elements to a map
	colors["pink"] = "#e4a0d5ff"
	colors["grey"] = "#afafafff"
	colors["uglycolor"] = "#4e3035ff"

	// Removing them
	delete(colors, "uglycolor")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
