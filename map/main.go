package main

import "fmt"

func main() {

	//var colorMap map[string]string

	//colorsMap := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	colors["white"] = "#ffffff"

	fmt.Println(colors)

}
