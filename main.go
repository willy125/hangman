package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

var dictionary = []string{
	"Palabras",
	"Barco",
	"Cañon",
	"Teclado",
	"Programacion",
	"Estados Unidos",
	"Mexico",
	"Telegrama",
}

func main() {
	//obtener una palabra
	rand.Seed(time.Now().UnixNano())
	targetWord := getRandomWorld()
	guessedLetters := initializeGuessedWorlds(targetWord)

	//fmt.Println(targetWord)

	printGameState(targetWord, guessedLetters)

}
func initializeGuessedWorlds(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[rune(targetWord[len(targetWord)-1])] = true
	return guessedLetters
}
func getRandomWorld() string {
	targetWord := dictionary[rand.Intn(len(dictionary))]
	return targetWord
}
func printGameState(targetWorld string, guessedLetters map[rune]bool) {
	for _, ch := range targetWorld {
		if ch == ' ' { //si el caracter es un espacio entonces añadimos un espacio
			fmt.Print(" ")
		} else if guessedLetters[unicode.ToLower(ch)] == true { // unicode.ToLower is for converting characters from capital letter to lower
			fmt.Printf("%c", ch)
		} else {
			fmt.Print("_")
		}
		fmt.Print(" ")
	}
	fmt.Println()
}
