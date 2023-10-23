package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Button struct { //caractéristique des boutons
	X, Y, Width, Height int
	Label               string
	Image               *ebiten.Image
	Active              bool
	TextColor           color.Color
}

var ( //initialisation des variables
	words          = []string{""}
	backgroundImg  *ebiten.Image
	Font           font.Face
	colorBlack     = color.RGBA{0, 0, 0, 255}
	colorRed       = color.RGBA{255, 0, 0, 255}
	mainMenu       MainMenu
	difficultyMenu DifficultyMenu
	gameMenu       GameMenu
	chosenWord     string
	selectedLetter rune
	gameInMenu     int
)

type MainMenu struct {
	Buttons []Button
}

type DifficultyMenu struct {
	Buttons []Button
}
type GameMenu struct {
	Buttons []Button
}

func init() { //met en place les images et boutons du jeu
	// Download image
	img, _, err := ebitenutil.NewImageFromFile("./images/fond.png", ebiten.FilterDefault) //chemin du fichier
	if err != nil {
		panic(err)
	}
	backgroundImg = img
	
	fontData, err := ioutil.ReadFile("./images/FFF_Tusj.ttf")
	if err != nil {
		log.Fatal(err)
	}

	font, err := opentype.Parse(fontData)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	Font, err = opentype.NewFace(font, &opentype.FaceOptions{
		Size: 32,
		DPI:  dpi,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Load images for buttons
	playImg, _, err := ebitenutil.NewImageFromFile("./images/button.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	quitImg, _, err := ebitenutil.NewImageFromFile("./images/button.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	//placement des boutons
	mainMenu.Buttons = []Button{
		{X: 720, Y: 450, Width: 170, Height: 60, Label: "Play", Image: playImg, Active: true},
		{X: 720, Y: 550, Width: 170, Height: 60, Label: "Quit", Image: quitImg, Active: true},
	}
	easyImg, _, err := ebitenutil.NewImageFromFile("./images/Picsart.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	mediumImg, _, err := ebitenutil.NewImageFromFile("./images/Picsart.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	hardImg, _, err := ebitenutil.NewImageFromFile("./images/Picsart.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	larousseImg, _, err := ebitenutil.NewImageFromFile("./images/Picsart.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	returnImg, _, err := ebitenutil.NewImageFromFile("./images/Picsart.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	difficultyMenu.Buttons = []Button{
		{X: 550, Y: 450, Width: 170, Height: 60, Label: "Easy", Image: easyImg, Active: true},
		{X: 900, Y: 450, Width: 170, Height: 60, Label: "Medium", Image: mediumImg, Active: true},
		{X: 550, Y: 550, Width: 170, Height: 60, Label: "Hard", Image: hardImg, Active: true},
		{X: 900, Y: 550, Width: 170, Height: 60, Label: "Larousse", Image: larousseImg, Active: true},
		{X: 720, Y: 650, Width: 170, Height: 60, Label: "Return", Image: returnImg, Active: true},
	}
	bkgImg, _, err := ebitenutil.NewImageFromFile("./images/bkg.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	//quitImg, _, err := ebitenutil.NewImageFromFile("./images/bkg.png", ebiten.FilterDefault)
	//if err != nil {
	//log.Fatal(err)
	//}

	gameMenu.Buttons = []Button{ //placement des boutons de lettre
		{X: 300, Y: 650, Width: 40, Height: 40, Label: "A", Image: bkgImg, Active: true},
		{X: 350, Y: 650, Width: 40, Height: 40, Label: "B", Image: bkgImg, Active: true},
		{X: 400, Y: 650, Width: 40, Height: 40, Label: "C", Image: bkgImg, Active: true},
		{X: 450, Y: 650, Width: 40, Height: 40, Label: "D", Image: bkgImg, Active: true},
		{X: 500, Y: 650, Width: 40, Height: 40, Label: "E", Image: bkgImg, Active: true},
		{X: 550, Y: 650, Width: 40, Height: 40, Label: "F", Image: bkgImg, Active: true},
		{X: 600, Y: 650, Width: 40, Height: 40, Label: "G", Image: bkgImg, Active: true},
		{X: 650, Y: 650, Width: 40, Height: 40, Label: "H", Image: bkgImg, Active: true},
		{X: 700, Y: 650, Width: 40, Height: 40, Label: "I", Image: bkgImg, Active: true},
		{X: 750, Y: 650, Width: 40, Height: 40, Label: "J", Image: bkgImg, Active: true},
		{X: 800, Y: 650, Width: 40, Height: 40, Label: "K", Image: bkgImg, Active: true},
		{X: 850, Y: 650, Width: 40, Height: 40, Label: "L", Image: bkgImg, Active: true},
		{X: 900, Y: 650, Width: 40, Height: 40, Label: "M", Image: bkgImg, Active: true},
		{X: 300, Y: 700, Width: 40, Height: 40, Label: "N", Image: bkgImg, Active: true},
		{X: 350, Y: 700, Width: 40, Height: 40, Label: "O", Image: bkgImg, Active: true},
		{X: 400, Y: 700, Width: 40, Height: 40, Label: "P", Image: bkgImg, Active: true},
		{X: 450, Y: 700, Width: 40, Height: 40, Label: "Q", Image: bkgImg, Active: true},
		{X: 500, Y: 700, Width: 40, Height: 40, Label: "R", Image: bkgImg, Active: true},
		{X: 550, Y: 700, Width: 40, Height: 40, Label: "S", Image: bkgImg, Active: true},
		{X: 600, Y: 700, Width: 40, Height: 40, Label: "T", Image: bkgImg, Active: true},
		{X: 650, Y: 700, Width: 40, Height: 40, Label: "U", Image: bkgImg, Active: true},
		{X: 700, Y: 700, Width: 40, Height: 40, Label: "V", Image: bkgImg, Active: true},
		{X: 750, Y: 700, Width: 40, Height: 40, Label: "W", Image: bkgImg, Active: true},
		{X: 800, Y: 700, Width: 40, Height: 40, Label: "X", Image: bkgImg, Active: true},
		{X: 850, Y: 700, Width: 40, Height: 40, Label: "Y", Image: bkgImg, Active: true},
		{X: 900, Y: 700, Width: 40, Height: 40, Label: "Z", Image: bkgImg, Active: true},
	}
}

//lance ou quitte le jeu
func Main(screen *ebiten.Image) error {
	if err := screen.DrawImage(backgroundImg, nil); err != nil {
		return err
	}

	for _, button := range mainMenu.Buttons {
		button.TextColor = colorBlack
		if button.Active {
			mouseX, mouseY := ebiten.CursorPosition()
			if mouseX >= button.X && mouseX <= button.X+button.Width && mouseY >= button.Y && mouseY <= button.Y+button.Height {
				button.TextColor = colorRed
			}
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				if mouseX >= button.X && mouseX <= button.X+button.Width && mouseY >= button.Y && mouseY <= button.Y+button.Height {
					if button.Label == "Play" {

						gameInMenu = 1
					} else if button.Label == "Quit" {
						fmt.Println("Goodbye!")
						os.Exit(0)
					}
				}
			}
		}

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(button.X), float64(button.Y))
		screen.DrawImage(button.Image, op)

		text.Draw(screen, button.Label, Font, button.X+65, button.Y+45, button.TextColor)
	}
	return nil
}

