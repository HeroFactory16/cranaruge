package hangman

import (
	"math/rand"
	"time"
)

func WordToGuess(word string) []rune {
	var placeHolder []rune = make([]rune, len(word))
	for i := 0; i < len(word); i++ {
		placeHolder[i] += '_'
	}
	rand.Seed(time.Now().UnixNano())
	letterClue := (len((word))/2 - 1)
	for i := 0; i < letterClue; i++ {
		index := rand.Intn(len(word))
		placeHolder[index] = []rune(word)[index]
	}
	return placeHolder
}
