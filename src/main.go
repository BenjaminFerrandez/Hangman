package main

import (
	"fmt"
	"os"
)

var choice string

func main() {
	fmt.Println("Bienvenue dans le jeu du pendu")
	fmt.Println("\nClick 's' to start or 'q' to quit")

	fmt.Scanln(&choice)
	switch choice {
	case "s":
		start()
	case "q":
		os.Exit(0)
	default:
		fmt.Println("Incorrect Choice. Please select an option from the menu (s or q)")
		main()
	}
}