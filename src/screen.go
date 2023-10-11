package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/image/font"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font/opentype"
)
type Button struct { //caractéristique des boutons
	X, Y, Width, Height int
	Label               string
	Image               *ebiten.Image
	Active              bool
}

var (
	backgroundImg *ebiten.Image
	buttons          []Button
	Font            font.Face
)

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
	buttons = []Button{
		{X: 720, Y: 450, Width: 200, Height: 40, Label: "Play", Image: playImg, Active: true},
		{X: 720, Y: 550, Width: 200, Height: 40, Label: "Quit", Image: quitImg, Active: true},
	}

}
func screen(screen *ebiten.Image) error {

	if err := screen.DrawImage(backgroundImg, nil); err != nil {
		return err
	}
	for _, button := range buttons {
		color := color.Black
		if button.Active {
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				mouseX, mouseY := ebiten.CursorPosition()
				if mouseX >= button.X && mouseX <= button.X+button.Width && mouseY >= button.Y && mouseY <= button.Y+button.Height {
					// button click
					if button.Label == "Play" {
						return fmt.Errorf("Play")
					} else if button.Label == "Quit" {
						fmt.Println("Goodbye!")
						os.Exit(0)
					}
				}
			}
		}

		// Draw the button image
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(button.X), float64(button.Y))
		screen.DrawImage(button.Image, op)

		// Draw the button label
		text.Draw(screen, button.Label, Font, button.X+65, button.Y+45, color)
	}
	return nil

	
}
