package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "ff0000",
		"green": "ff1000",
		"white": "ff2000",
	}

	//colors["white"] = "ff0000"
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, "is", hex)
	}
}
