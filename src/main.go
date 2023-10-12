package main

import (
	"fmt"
	"os"
	"github.com/hajimehoshi/ebiten" //import de bibliotheque ebiten pour graphique

)
const ( //taille de l'écran
	screenWidth  = 1280
	screenHeight = 940
)

var choice string

//premiere fonction
func main() {

	ebiten.Run(update, screenWidth, screenHeight, 1, "Hangman game") //titre
	fmt.Println("Welcome in our Hangman game!")
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