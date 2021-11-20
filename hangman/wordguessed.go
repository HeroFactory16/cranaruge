package hangman

import (
	"fmt"
)

func WordGuessed(guess, word string, wordToGuess []rune) bool {
	if guess == word+"\n" {
		for c, j := range word {
			wordToGuess[c] = j
		}
		fmt.Println("Bravo, vous avez gagné !")
		return true
	}
	return false
}
