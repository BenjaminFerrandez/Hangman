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

var (
	backgroundImg  *ebiten.Image
	Font           font.Face
	colorBlack     = color.RGBA{0, 0, 0, 255}
	colorRed       = color.RGBA{255, 0, 0, 255}
	mainMenu       MainMenu
	difficultyMenu DifficultyMenu
	gameInMenu     int
)

type MainMenu struct {
	Buttons []Button
}

type DifficultyMenu struct {
	Buttons []Button
}

func init() {
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
		{X: 720, Y: 450, Width: 200, Height: 60, Label: "Play", Image: playImg, Active: true},
		{X: 720, Y: 550, Width: 200, Height: 60, Label: "Quit", Image: quitImg, Active: true},
	}
	easyImg, _, err := ebitenutil.NewImageFromFile("./images/button.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	mediumImg, _, err := ebitenutil.NewImageFromFile("./images/button.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	hardImg, _, err := ebitenutil.NewImageFromFile("./images/button.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	larousseImg, _, err := ebitenutil.NewImageFromFile("./images/button.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	returnImg, _, err := ebitenutil.NewImageFromFile("./images/button.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	difficultyMenu.Buttons = []Button{
		{X: 720, Y: 450, Width: 200, Height: 40, Label: "Easy", Image: easyImg, Active: true},
		{X: 720, Y: 550, Width: 200, Height: 40, Label: "Medium", Image: mediumImg, Active: true},
		{X: 720, Y: 650, Width: 200, Height: 40, Label: "Hard", Image: hardImg, Active: true},
		{X: 720, Y: 750, Width: 200, Height: 40, Label: "Larousse", Image: larousseImg, Active: true},
		{X: 720, Y: 850, Width: 200, Height: 40, Label: "Return", Image: returnImg, Active: true},
	}

}
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
func Difficulty(screen *ebiten.Image) error {
	if err := screen.DrawImage(backgroundImg, nil); err != nil {
		return err
	}
	for _, button := range difficultyMenu.Buttons {
		button.TextColor = colorBlack
		if button.Active {
			mouseX, mouseY := ebiten.CursorPosition()
			if mouseX >= button.X && mouseX <= button.X+button.Width && mouseY >= button.Y && mouseY <= button.Y+button.Height {
				button.TextColor = colorRed
			}
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				if mouseX >= button.X && mouseX <= button.X+button.Width && mouseY >= button.Y && mouseY <= button.Y+button.Height {
					if button.Label == "Easy" {

					} else if button.Label == "Medium" {

					} else if button.Label == "Hard" {

					} else if button.Label == "Larousse" {

					} else if button.Label == "Return" {
						gameInMenu = 0
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

func update(screen *ebiten.Image) error {
	
	if gameInMenu == 1 {

		return Difficulty(screen)
	} else if gameInMenu == 0 {
		return Main(screen)
	}
	return nil
}
