package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"math/rand"
	"time"

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
	blancImg       *ebiten.Image
	backgroundImg  *ebiten.Image
	Font           font.Face
	colorBlack     = color.RGBA{0, 0, 0, 255}
	colorRed       = color.RGBA{255, 0, 0, 255}
	mainMenu       MainMenu
	difficultyMenu DifficultyMenu
	gameMenu       GameMenu
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

//met en place les images et boutons du jeu
func init() {
	// Download image
	img, _, err := ebitenutil.NewImageFromFile("./images/fond.png", ebiten.FilterDefault) //chemin du fichier
	if err != nil {
		panic(err)
	}
	backgroundImg = img
	blanc, _, err := ebitenutil.NewImageFromFile("./images/blanc.png", ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	blancImg = blanc
	fontData, err := ioutil.ReadFile("./images/FFF_Tusj.ttf") //police d'écriture
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
		log.Fatal(err) //pour vérifier si on clique sur le boutons
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
		{X: 300, Y: 650, Width: 200, Height: 0, Label: "A", Image: bkgImg, Active: true},
		{X: 350, Y: 650, Width: 200, Height: 0, Label: "B", Image: bkgImg, Active: true},
		{X: 400, Y: 650, Width: 200, Height: 0, Label: "C", Image: bkgImg, Active: true},
		{X: 450, Y: 650, Width: 200, Height: 0, Label: "D", Image: bkgImg, Active: true},
		{X: 500, Y: 650, Width: 200, Height: 0, Label: "E", Image: bkgImg, Active: true},
		{X: 550, Y: 650, Width: 200, Height: 0, Label: "F", Image: bkgImg, Active: true},
		{X: 600, Y: 650, Width: 200, Height: 0, Label: "G", Image: bkgImg, Active: true},
		{X: 650, Y: 650, Width: 200, Height: 0, Label: "H", Image: bkgImg, Active: true},
		{X: 700, Y: 650, Width: 200, Height: 0, Label: "I", Image: bkgImg, Active: true},
		{X: 750, Y: 650, Width: 200, Height: 0, Label: "J", Image: bkgImg, Active: true},
		{X: 800, Y: 650, Width: 200, Height: 0, Label: "K", Image: bkgImg, Active: true},
		{X: 850, Y: 650, Width: 200, Height: 0, Label: "L", Image: bkgImg, Active: true},
		{X: 900, Y: 650, Width: 200, Height: 0, Label: "M", Image: bkgImg, Active: true},
		{X: 300, Y: 700, Width: 200, Height: 0, Label: "N", Image: bkgImg, Active: true},
		{X: 350, Y: 700, Width: 200, Height: 0, Label: "O", Image: bkgImg, Active: true},
		{X: 400, Y: 700, Width: 200, Height: 0, Label: "P", Image: bkgImg, Active: true},
		{X: 450, Y: 700, Width: 200, Height: 0, Label: "Q", Image: bkgImg, Active: true},
		{X: 500, Y: 700, Width: 200, Height: 0, Label: "R", Image: bkgImg, Active: true},
		{X: 550, Y: 700, Width: 200, Height: 0, Label: "S", Image: bkgImg, Active: true},
		{X: 600, Y: 700, Width: 200, Height: 0, Label: "T", Image: bkgImg, Active: true},
		{X: 650, Y: 700, Width: 200, Height: 0, Label: "U", Image: bkgImg, Active: true},
		{X: 700, Y: 700, Width: 200, Height: 0, Label: "V", Image: bkgImg, Active: true},
		{X: 750, Y: 700, Width: 200, Height: 0, Label: "W", Image: bkgImg, Active: true},
		{X: 800, Y: 700, Width: 200, Height: 0, Label: "X", Image: bkgImg, Active: true},
		{X: 850, Y: 700, Width: 200, Height: 0, Label: "Y", Image: bkgImg, Active: true},
		{X: 900, Y: 700, Width: 200, Height: 0, Label: "Z", Image: bkgImg, Active: true},
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

//choix du niveau
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
					
							words = []string{"ane", "axe", "coq", "cou", "cri", "gag", "gaz", "gel", "jus", "nul", "ski", "tas", "tic",
								"beau", "boxe", "brun", "cerf", "cire", "dame", "dent", "dodo", "drap", "dune", "jazz", "joli", "joue", "logo", "loin", "long", "lune", "lynx", "mine", "ours", "pion", "seau", "test", "trou", "truc", "vert",
								"aimer", "assez", "avion", "balai", "banjo", "barbe", "bruit", "buche", "capot", "carte", "chien", "cycle", "essai", "jambe", "koala", "livre", "noeud", "ortie", "poire", "pomme", "prune", "radar", "radis", "robot", "route", "rugby", "taupe", "tenue", "texte", "valse"}
							gameInMenu = 2
						
					} else if button.Label == "Medium" {
					
							words = []string{"acajou", "agneau", "alarme", "ananas", "animal", "arcade", "aviron", "balade", "billet", "bouche", "boucle", "bronze", "cabane", "cloche", "coccyx", "crayon", "garage", "goulot", "gramme", "grelot", "humour", "limite", "lionne", "menthe", "oiseau", "podium", "poulpe", "poumon", "puzzle", "rapide", "tomate", "walabi", "whisky",
								"abriter", "batavia", "billard", "bretzel", "chariot", "clairon", "corbeau", "cortège", "crapaud", "cymbale", "dentier", "djembé", "drapeau", "exemple", "fourmis", "grandir", "iceberg", "javelot", "journal", "journee", "losange", "mondial", "oxygene", "panique", "petrole", "poterie", "pouvoir", "scooter", "sifflet", "spirale", "sucette", "strophe", "tonneau", "trousse", "tunique", "ukulele", "vautour", "zozoter",
								"aquarium", "araignee", "arbalete", "archipel", "banquise", "batterie", "brocante", "brouhaha", "cloporte", "debutant", "diapason", "gangster", "gothique", "hautbois", "herisson", "logiciel", "objectif", "parcours", "question", "scorpion", "symptome", "tabouret", "taboulet", "toujours", "tourisme", "triangle", "utopique"}
							gameInMenu = 3
						
					} else if button.Label == "Hard" {
				
							words = []string{"accordeon", "ascenseur", "ascension", "aseptiser", "autoroute", "avalanche", "bilboquet", "bourricot", "brillance", "cabriolet", "cornemuse", "dangereux", "epluchage", "forteresse", "graphique", "horoscope", "intrepide", "klaxonner", "mascarade", "metaphore", "narrateur", "populaire", "printemps", "tambourin", "vestiaire", "xylophone",
								"acrostiche", "apocalypse", "attraction", "aventurier", "bouillotte", "citrouille", "controverse", "coquelicot", "dissimuler", "flibustier", "grenouille", "impossible", "labyrinthe", "prudemment", "quadriceps", "soliloquer", "subjective"}
							gameInMenu = 4
						
					} else if button.Label == "Larousse" {
					
							words = []string{"baccalaureat", "abracadabra", "francophile", "pandemonium", "chlorophylle", "metallurgie", "metamorphose", "montgolfiere", "kaleidoscope", "conquistador", "conspirateur", "rhododendron", "qualification", "protozoaire", "quadrilatère", "zygomatique", "sorcellerie", "belligerant"}
							gameInMenu = 5
						
					} else if button.Label == "Return" {
							gameInMenu = 0
					}
				}
			}
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(button.X), float64(button.Y))
		screen.DrawImage(button.Image, op)

		text.Draw(screen, button.Label, Font, button.X+45, button.Y+45, button.TextColor)
	}
	return nil
}
var chosenWord string

//jeu
func pendu(screen *ebiten.Image) error {
    screen.Fill(color.White)

    if chosenWord == "" {
        rand.Seed(time.Now().UnixNano())
        chosenWord = words[rand.Intn(len(words))]
    }

    guessedWord := make([]string, len(chosenWord))
    text.Draw(screen, "Atempts remaining :", Font, 20, 50, colorBlack)
    for i := range guessedWord {
        guessedWord[i] = "_"
    }
	if gameInMenu == 2 {
		text.Draw(screen, strings.Join(guessedWord, " "), Font, 575, 600, colorBlack)
	} else if gameInMenu == 3 {
		text.Draw(screen, strings.Join(guessedWord, " "), Font, 550, 600, colorBlack)
	} else if gameInMenu == 4 {
		text.Draw(screen, strings.Join(guessedWord, " "), Font, 500, 600, colorBlack)
	} else if gameInMenu == 5 {
		text.Draw(screen, strings.Join(guessedWord, " "), Font, 450, 600, colorBlack)
	}
    for _, button := range gameMenu.Buttons {
        button.TextColor = colorBlack
        op := &ebiten.DrawImageOptions{}
        op.GeoM.Translate(float64(button.X), float64(button.Y))
        screen.DrawImage(button.Image, op)

        text.Draw(screen, button.Label, Font, button.X+13, button.Y+34, button.TextColor)
    }
    return nil
}

//relance une partie ou quitte le jeu
func update(screen *ebiten.Image) error {
	if gameInMenu == 1 {
		return Difficulty(screen)
	} else if gameInMenu == 0 {
		return Main(screen)
	} else if gameInMenu > 1 {
		return pendu(screen)
	}
	return nil
}
