package main

import (
    "fmt"
    "strings"
)

func main() {
    wordToGuess := "golang" // Le mot à deviner
    guessedWord := make([]string, len(wordToGuess))

    for i := range guessedWord {
        guessedWord[i] = "_"
    }

    attempts := 6 // Nombre d'essais autorisés

    fmt.Println("Bienvenue dans le jeu du pendu!")

    for attempts > 0 {
        fmt.Println("Mot à deviner: ", strings.Join(guessedWord, " "))
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
            fmt.Println("Bravo! Vous avez deviné le mot:", wordToGuess)
            break
        }
    }

    if attempts == 0 {
        fmt.Println("Désolé, vous avez épuisé toutes vos tentatives. Le mot était:", wordToGuess)
    }
}
