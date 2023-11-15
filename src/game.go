package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

var (
	chosenWord      string
	selectedLetter  rune
	guessedLetters  []rune
	selectedLetters []rune
	Essai           int64 = 6
)

//lance ou quitte le jeu
func Main(screen *ebiten.Image) error {
	screen.DrawImage(backgroundImg, nil)

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

//jeu
func pendu(screen *ebiten.Image) error {
	screen.Fill(color.White)
	s := strconv.FormatInt(Essai, 10)
	if chosenWord == "" || Essai == 0 {
		rand.Seed(time.Now().UnixNano())
		chosenWord = words[rand.Intn(len(words))]
		Essai = 6
		selectedLetters = nil
	}

	guessedWord := make([]string, len(chosenWord))
	text.Draw(screen, "Attempts remaining :", Font, 20, 50, colorBlack)
	text.Draw(screen, s, Font, 360, 50, colorBlack)
	for i := range guessedWord {
		guessedWord[i] = "_"
	}

	for _, button := range gameMenu.Buttons {
		button.TextColor = colorBlack

		mouseX, mouseY := ebiten.CursorPosition()
		if mouseX >= button.X && mouseX <= button.X+button.Width && mouseY >= button.Y && mouseY <= button.Y+button.Height {
			button.TextColor = colorRed
		}

		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			mouseX, mouseY := ebiten.CursorPosition()
			for _, button := range gameMenu.Buttons {
				if mouseX >= button.X && mouseX <= button.X+button.Width && mouseY >= button.Y && mouseY <= button.Y+button.Height {
					if button.Label == "A" {
						selectedLetter = rune('a')
					} else if button.Label == "B" {
						selectedLetter = rune('b')

					} else if button.Label == "C" {
						selectedLetter = rune('c')

					} else if button.Label == "D" {
						selectedLetter = rune('d')

					} else if button.Label == "E" {
						selectedLetter = rune('e')

					} else if button.Label == "F" {
						selectedLetter = rune('f')

					} else if button.Label == "G" {
						selectedLetter = rune('g')

					} else if button.Label == "H" {
						selectedLetter = rune('h')

					} else if button.Label == "I" {
						selectedLetter = rune('i')

					} else if button.Label == "J" {
						selectedLetter = rune('j')

					} else if button.Label == "K" {
						selectedLetter = rune('k')

					} else if button.Label == "L" {
						selectedLetter = rune('l')

					} else if button.Label == "M" {
						selectedLetter = rune('m')

					} else if button.Label == "N" {
						selectedLetter = rune('n')

					} else if button.Label == "O" {
						selectedLetter = rune('o')

					} else if button.Label == "P" {
						selectedLetter = rune('p')

					} else if button.Label == "Q" {
						selectedLetter = rune('q')

					} else if button.Label == "R" {
						selectedLetter = rune('r')

					} else if button.Label == "S" {
						selectedLetter = rune('s')

					} else if button.Label == "T" {
						selectedLetter = rune('t')

					} else if button.Label == "U" {
						selectedLetter = rune('u')

					} else if button.Label == "V" {
						selectedLetter = rune('v')

					} else if button.Label == "W" {
						selectedLetter = rune('w')

					} else if button.Label == "X" {
						selectedLetter = rune('x')

					} else if button.Label == "Y" {
						selectedLetter = rune('y')

					} else if button.Label == "Z" {
						selectedLetter = rune('z')

					}
					if selectedLetter != 0 {
						if strings.Contains(chosenWord, string(selectedLetter)) {
							selectedLetters = append(selectedLetters, selectedLetter)
						} else {
							handleLetterSelection(button.Label)
						}
					}
					break
				}
			}
		}
		guessedWord = updateGuessedWord(chosenWord, selectedLetters)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(button.X), float64(button.Y))
		screen.DrawImage(button.Image, op)

		text.Draw(screen, button.Label, Font, button.X+13, button.Y+34, button.TextColor)
		if gameInMenu == 2 {
			text.Draw(screen, strings.Join(guessedWord, " "), Font, 575, 600, colorBlack)
		} else if gameInMenu == 3 {
			text.Draw(screen, strings.Join(guessedWord, " "), Font, 550, 600, colorBlack)
		} else if gameInMenu == 4 {
			text.Draw(screen, strings.Join(guessedWord, " "), Font, 500, 600, colorBlack)
		} else if gameInMenu == 5 {
			text.Draw(screen, strings.Join(guessedWord, " "), Font, 450, 600, colorBlack)
		}
		//check for victory
		if strings.Join(updateGuessedWord(chosenWord, selectedLetters), "") == chosenWord {
			chosenWord = ""
			Essai = 6
			selectedLetters = nil

			gameInMenu = -2
		}
	}

	return nil
}
func handleLetterSelection(buttonLabel string) {
	selectedLetter = rune(buttonLabel[0])
	if selectedLetter != 0 && !contains(selectedLetters, selectedLetter) {
		selectedLetters = append(selectedLetters, selectedLetter)
		if !strings.Contains(chosenWord, string(selectedLetter)) {
			Essai -= 1

		}
		if Essai == 0 {
			gameInMenu = -1

		}
	}
}
func defeat(screen *ebiten.Image) error {
	screen.DrawImage(backgroundLose, nil)
	text.Draw(screen, "You lose", Font, 600, 475, colorBlack)
	for _, button := range lastMenu.Buttons {
		button.TextColor = colorBlack
		if button.Active {
			mouseX, mouseY := ebiten.CursorPosition()
			if mouseX >= button.X && mouseX <= button.X+button.Width && mouseY >= button.Y && mouseY <= button.Y+button.Height {
				button.TextColor = colorRed
			}
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				if mouseX >= button.X && mouseX <= button.X+button.Width && mouseY >= button.Y && mouseY <= button.Y+button.Height {

					if button.Label == "Return" {
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
func victory(screen *ebiten.Image) error {
	screen.DrawImage(backgroundWin, nil)
	text.Draw(screen, "You won", Font, 655, 610, colorBlack)
	for _, button := range winMenu.Buttons {
		button.TextColor = colorBlack
		if button.Active {
			mouseX, mouseY := ebiten.CursorPosition()
			if mouseX >= button.X && mouseX <= button.X+button.Width && mouseY >= button.Y && mouseY <= button.Y+button.Height {
				button.TextColor = colorRed
			}
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				if mouseX >= button.X && mouseX <= button.X+button.Width && mouseY >= button.Y && mouseY <= button.Y+button.Height {

					if button.Label == "Return" {
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

//relance une partie ou quitte le jeu
func update(screen *ebiten.Image) error {
	if gameInMenu == 1 {
		return Difficulty(screen)
	} else if gameInMenu == 0 {
		return Main(screen)
	} else if gameInMenu > 1 {
		return pendu(screen)
	} else if gameInMenu == -1 {
		return defeat(screen)
	} else if gameInMenu == -2 {
		return victory(screen)
	}
	return nil
}
func updateGuessedWord(word string, guessedLetters []rune) []string {
	guessedWord := make([]string, len(word))

	for i, char := range word {
		if contains(guessedLetters, char) {
			guessedWord[i] = string(char)
		} else {
			guessedWord[i] = "_"
		}
	}

	return guessedWord
}

func contains(arr []rune, r rune) bool {
	for _, v := range arr {
		if v == r {
			return true
		}
	}
	return false
}
