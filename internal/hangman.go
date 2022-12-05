package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var inputReader = bufio.NewReader(os.Stdin)
var dictionary = []string{
	"Zombie",
	"Gopher",
	"United States of America",
	"Indonesia",
	"Nazism",
	"Apple",
	"Programming",
}

func Start() (string, map[rune]bool) {
	rand.Seed(time.Now().UnixNano())

	targetWord := getRandomWord()
	guessedLetters := initializeGuessedWords(targetWord)

	return targetWord, guessedLetters
	// for !isGameOver(targetWord, guessedLetters, hangmanState) {
	// 	printGameState(targetWord, guessedLetters, hangmanState)
	// 	input := readInput()
	// 	if len(input) != 1 {
	// 		fmt.Println("Invalid input. Please use letters only...")
	// 		continue
	// 	}

	// 	letter := rune(input[0])
	// 	if isCorrectGuess(targetWord, letter) {
	// 		guessedLetters[letter] = true
	// 	} else {
	// 		hangmanState++
	// 	}
	// }

	// printGameState(targetWord, guessedLetters, hangmanState)
	// fmt.Print("Game Over... ")
	// if isWordGuessed(targetWord, guessedLetters) {
	// 	fmt.Println("You win!")
	// } else if isHangmanComplete(hangmanState) {
	// 	fmt.Println("You lose!")
	// } else {
	// 	panic("invalid state. Game is over and there is no winner!")
	// }
}

func getRandomWord() string {
	targetWord := dictionary[rand.Intn(len(dictionary))]
	return targetWord
}

func initializeGuessedWords(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true

	return guessedLetters
}

func isGameOver(
	targetWord string,
	guessedLetters map[rune]bool,
	hangmanState int,
) bool {
	return isWordGuessed(targetWord, guessedLetters) ||
		isHangmanComplete(hangmanState)
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

// func printGameState(
// 	targetWord string,
// 	guessedLetters map[rune]bool,
// 	hangmanState int,
// ) {
// 	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
// 	fmt.Println()
// 	fmt.Println(getHangmanDrawing(hangmanState))
// }

func getWordGuessingProgress(
	targetWord string,
	guessedLetters map[rune]bool,
) string {
	result := ""
	for _, ch := range targetWord {
		if ch == ' ' {
			result += " "
		} else if guessedLetters[unicode.ToLower(ch)] {
			result += fmt.Sprintf("%c", ch)
		} else {
			result += "_"
		}

		result += " "
	}

	return result
}

func isCorrectGuess(targetWord string, letter rune) bool {
	return strings.ContainsRune(targetWord, letter)
}
