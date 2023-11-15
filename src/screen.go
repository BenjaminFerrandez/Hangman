package main

import (
	"image/color"
	"io/ioutil"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
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
	words = []string{""}

	backgroundImg  *ebiten.Image
	backgroundLose *ebiten.Image
	backgroundWin  *ebiten.Image
	Font           font.Face
	colorBlack     = color.RGBA{0, 0, 0, 255}
	colorRed       = color.RGBA{255, 0, 0, 255}
	mainMenu       MainMenu
	difficultyMenu DifficultyMenu
	gameMenu       GameMenu
	lastMenu       LastMenu
	winMenu        WinMenu
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
type LastMenu struct {
	Buttons []Button
}
type WinMenu struct {
	Buttons []Button
}

func init() { //met en place les images et boutons du jeu
	// Download image
	img, _, err := ebitenutil.NewImageFromFile("./images/fond.png", ebiten.FilterDefault) //chemin du fichier
	if err != nil {
		panic(err)
	}
	backgroundImg = img

	loseImg, _, err := ebitenutil.NewImageFromFile("./images/lose.png", ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	backgroundLose = loseImg

	winImg, _, err := ebitenutil.NewImageFromFile("./images/win.png", ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	backgroundWin = winImg

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
	lastMenu.Buttons = []Button{
		{X: 550, Y: 550, Width: 170, Height: 60, Label: "Return", Image: returnImg, Active: true},
	}
	winMenu.Buttons = []Button{
		{X: 625, Y: 670, Width: 170, Height: 60, Label: "Return", Image: returnImg, Active: true},
	}
	bkgImg, _, err := ebitenutil.NewImageFromFile("./images/bkg.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

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
