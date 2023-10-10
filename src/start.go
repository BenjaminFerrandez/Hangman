package main

import (
    "fmt"
	"math/rand"
	"time"
    "strings"
    "os"
)

var difficulty string

//jeu
func start() {
    words := []string{""}

    fmt.Print("\033[H\033[2J")
    fmt.Println("\nChoose a difficulty")
    fmt.Println("1. Easy (3, 4 or 5 letters in the word)")
    fmt.Println("2. Medium (6, 7 or 8 letters in the word)")
    fmt.Println("3. Hard (9 or 10 letters in the word)")
    fmt.Println("4. Larousse (11, 12 or 13 letters in the word)")

    fmt.Scanln(&difficulty)
    switch difficulty {
    case "1":
        words = []string{"ane", "axe", "coq", "cou", "cri", "gag", "gaz", "gel", "jus", "nul", "ski", "tas", "tic",
        "beau", "boxe", "brun", "cerf", "cire", "dame", "dent", "dodo", "drap", "dune", "jazz", "joli", "joue", "logo", "loin", "long", "lune", "lynx", "mine", "ours", "pion", "seau", "test", "trou", "truc", "vert",
        "aimer", "assez", "avion", "balai", "banjo", "barbe", "bruit", "buche", "capot", "carte", "chien", "cycle", "essai", "jambe", "koala", "livre", "noeud", "ortie", "poire", "pomme", "prune", "radar", "radis", "robot", "route", "rugby", "taupe", "tenue", "texte", "valse"}
    case "2":
        words = []string{"acajou", "agneau", "alarme", "ananas", "animal", "arcade", "aviron", "balade", "billet", "bouche", "boucle", "bronze", "cabane", "cloche", "coccyx", "crayon", "garage", "goulot", "gramme", "grelot", "humour", "limite", "lionne", "menthe", "oiseau", "podium", "poulpe", "poumon", "puzzle", "rapide", "tomate", "walabi", "whisky",
        "abriter", "batavia", "billard", "bretzel", "chariot", "clairon", "corbeau", "cortège", "crapaud", "cymbale", "dentier", "djembé", "drapeau", "exemple", "fourmis", "grandir", "iceberg", "javelot", "journal", "journee", "losange", "mondial", "oxygene", "panique", "petrole", "poterie", "pouvoir", "scooter", "sifflet", "spirale", "sucette", "strophe", "tonneau", "trousse", "tunique", "ukulele", "vautour", "zozoter",
        "aquarium", "araignee", "arbalete", "archipel", "banquise", "batterie", "brocante", "brouhaha", "cloporte", "debutant", "diapason", "gangster", "gothique", "hautbois", "herisson", "logiciel", "objectif", "parcours", "question", "scorpion", "symptome", "tabouret", "taboulet", "toujours", "tourisme", "triangle", "utopique"}
    case "3":
        words = []string{"accordeon", "ascenseur", "ascension", "aseptiser", "autoroute", "avalanche", "bilboquet", "bourricot", "brillance", "cabriolet", "cornemuse", "dangereux", "epluchage", "forteresse", "graphique", "horoscope", "intrepide", "klaxonner", "mascarade", "metaphore", "narrateur", "populaire", "printemps", "tambourin", "vestiaire", "xylophone",
        "acrostiche", "apocalypse", "attraction", "aventurier", "bouillotte", "citrouille", "controverse", "coquelicot", "dissimuler", "flibustier", "grenouille", "impossible", "labyrinthe", "prudemment", "quadriceps", "soliloquer", "subjective"}
    case "4":
        words = []string{"baccalaureat", "abracadabra", "francophile", "pandemonium", "chlorophylle", "metallurgie", "metamorphose", "montgolfiere", "kaleidoscope", "conquistador", "conspirateur", "rhododendron", "qualification", "protozoaire", "quadrilatère", "zygomatique", "sorcellerie", "belligerant"}
    }

	rand.Seed(time.Now().UnixNano())
	wordToGuess := words[rand.Intn(len(words))]
    guessedWord := make([]string, len(wordToGuess))

    for i := range guessedWord {
        guessedWord[i] = "_"
    }

    attempts := 6 // Nombre d'essais autorisés
    fmt.Print("\033[H\033[2J")

    for attempts > 0 {
        fmt.Println("\nWord to find: ", strings.Join(guessedWord, " "))
        fmt.Printf("attempts remaining: %d\n", attempts)

        var guess string
        fmt.Print("Try a letter: ")
        fmt.Scanln(&guess)

        if len(guess) != 1 {
            fmt.Println("Please try only one letter")
            continue
        }

        guess = strings.ToLower(guess)

        if strings.Contains(wordToGuess, guess) {
            for i, letter := range wordToGuess {
                if string(letter) == guess {
                    guessedWord[i] = guess
                }
            }
        } else {
            fmt.Printf("%s isn't in the word\n", guess)
            attempts--
        }

        if strings.Join(guessedWord, "") == wordToGuess {
            fmt.Print("\033[H\033[2J")
            fmt.Println("\nWell done! You have find the word", wordToGuess)
            victory()
        }
    }

    if attempts == 0 {
        fmt.Print("\033[H\033[2J")
        fmt.Println("\nSorry, you have lose. The word was:", wordToGuess)
        lose()
    }
}

//si on gagne
func victory() {
    var win string
    fmt.Println("You have won!")
    fmt.Println("Click 'c' to continue or 'q' to quit")
    fmt.Scanln(&win)
    switch win {
    case "c":
        start()
    case "q":
        os.Exit(0)
    default:
        fmt.Println("Incorrect choice")
        victory()
    }
}

//si on perd
func lose() {
    fmt.Println("Click 'r' to restart or 'q' to quit")
    var l string
    fmt.Scanln(&l)
    switch l {
    case "r":
        start()
    case "q":
        os.Exit(0)
    default:
        fmt.Println("Incorrect choice")
        lose() 
    }
}