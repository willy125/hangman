package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

var inputReader = ""
var dictionary = []string{
	"Palabras",
	"Barco",
	"Cañon",
	"Teclado",
	"Programacion",
	"Estados Unidos",
	"Mexico",
	"Telegrama",
	"Minimo",
}

func main() {
	//obtener una palabra
	rand.Seed(time.Now().UnixNano())

	targetWord := getRandomWorld()
	guessedLetters := initializeGuessedWorlds(targetWord)

	//fmt.Println(targetWord)
	//guessedLetters['s'] = true
	//guessedLetters['i'] = true
	hangmanState := 0
	for !isGameOver(targetWord, guessedLetters, hangmanState) {
		printGameState(targetWord, guessedLetters, hangmanState)
		input := readinput()
		if len(input) != 1 {
			fmt.Println("Input invalido. Porfavor ingresa solo una letra")
			continue
		}
		letter := rune(input[0])
		if isCorrectGuess(targetWord, input) {
			guessedLetters[letter] = true
		} else {
			hangmanState++
		}
	}
	printGameState(targetWord, guessedLetters, hangmanState)
	fmt.Print("Game Over...")
	if isWordGuessed(targetWord, guessedLetters) {
		fmt.Println("you win!")
	} else if isHangmanComplete(hangmanState) {
		fmt.Println("You Lose!")
	} else {
		panic("invalid state. Game is over and there is no winner!")
	}
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
func isGameOver(targetWord string,
	guessedLetters map[rune]bool,
	hangmanState int) bool {
	return isWordGuessed(targetWord, guessedLetters) || isHangmanComplete(hangmanState)
}
func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, ch := range targetWord {
		if !guessedLetters[unicode.ToLower(ch)] {
			return false
		}
	}
	return true
}
func isHangmanComplete(hangmanState int) bool {
	return hangmanState >= 9
}
func printGameState(
	targetWorld string,
	guessedLetters map[rune]bool,
	hangmanState int) {
	fmt.Println(getWorldGuesseingProgress(targetWorld, guessedLetters))
	fmt.Println()
	fmt.Println(getHangmanDrawing(hangmanState))
}
func getWorldGuesseingProgress(targetWord string, guessedLetters map[rune]bool) string {
	result := ""
	for _, ch := range targetWord {
		if ch == ' ' { //si el caracter es un espacio entonces añadimos un espacio
			result += " "
		} else if guessedLetters[unicode.ToLower(ch)] == true { // unicode.ToLower is for converting characters from capital letter to lower
			result += fmt.Sprintf("%c", ch)
		} else {
			result += "_"
		}
		result += " "
	}
	return result
}

func getHangmanDrawing(hangmanState int) string {
	data, err := ioutil.ReadFile(fmt.Sprintf("states/hangman%d.txt", hangmanState))
	if err != nil {
		panic(err)
	}
	return string(data)
}
func readinput() string {
	fmt.Print("> ")
	fmt.Scanln(&inputReader)
	return strings.TrimSpace(inputReader)
}
func isCorrectGuess(targetWord string, letter string) bool {
	return strings.Contains(targetWord, letter)
}
