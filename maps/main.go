package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#afsdsadas",
		"green": "#adwsdasd12",
		"black": "#0000",
	}
	printColors(colors)
}

func printColors(c map[string]string) {
	for color, hex := range c {
		fmt.Println("The hex value for color", color, "is", hex)
	}
}
