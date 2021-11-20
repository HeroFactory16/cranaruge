package hangman

import (
	"fmt"
)

func UpdateWord(wordToGuess []rune, guess, word string) ([]rune, int) {
	tmp := string(wordToGuess)
	for c, j := range word {
		for _, k := range guess {
			if k == j {
				wordToGuess[c] = j
				if word == string(wordToGuess) {
					fmt.Println("Bravo, vous avez gagné !")
				}
				if tmp == string(wordToGuess) {
					return wordToGuess, 1
				}
			}
		}
	}
	return wordToGuess, 0
}
