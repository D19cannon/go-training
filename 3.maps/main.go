package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)

	// colors["red"] = "#ff0000"
	// colors["yellow"] = "#f4b745"

	colors := map[string]string{
		"red":    "#ff0000",
		"yellow": "#f4b745",
		"white":  "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("color: %v with the hex code: %v \n", color, hex)
	}
}
