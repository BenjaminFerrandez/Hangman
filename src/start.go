package main

import (
    "fmt"
	"math/rand"
	"time"
    "strings"
    "os"
)

func start() {
	rand.Seed(time.Now().UnixNano())
    words := []string{"leave", "good", "thing", "child", "young", "little",
	"great", "tractor", "chicken", "expensive", "photograph", "generous",
	"butterfly", "success", "lollipop", "answer", "beautiful",
	"bridge", "country", "bucket", "beauty", "hundred", "thousand", "student", "slowly",
	"window", "pattern", "sentence", "school", "mountain", "summer", "winter",
	"thing", "probably", "upstairs", "everything", "remember", "clothes", "shoulder",
	"himself", "perhaps", "always", "birthday", "forgive", "become", "shower",
	"strawberry", "walking", "yellow"} // Le mot à deviner
	wordToGuess := words[rand.Intn(len(words))]
    guessedWord := make([]string, len(wordToGuess))

    for i := range guessedWord {
        guessedWord[i] = "_"
    }

    attempts := 6 // Nombre d'essais autorisés

    fmt.Print("\033[H\033[2J")
    fmt.Println("\nBienvenue dans le jeu du pendu!")

    for attempts > 0 {
        fmt.Println("\nMot à deviner: ", strings.Join(guessedWord, " "))
        fmt.Printf("Tentatives restantes: %d\n", attempts)

        var guess string
        fmt.Print("Devinez une lettre: ")
        fmt.Scanln(&guess)

        if len(guess) != 1 {
            fmt.Println("Veuillez entrer une seule lettre.")
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
            fmt.Printf("%s n'est pas dans le mot.\n", guess)
            attempts--
        }

        if strings.Join(guessedWord, "") == wordToGuess {
            fmt.Print("\033[H\033[2J")
            fmt.Println("\nBravo! Vous avez deviné le mot:", wordToGuess)
            victory()
        }
    }

    if attempts == 0 {
        fmt.Print("\033[H\033[2J")
        fmt.Println("\nDésolé, vous avez épuisé toutes vos tentatives. Le mot était:", wordToGuess)
        lose()
    }
}

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

func lose() {
    fmt.Println("You have lose")
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