package hangman

import (
	"math/rand"
	"time"
)

func Randomize(word string) string {
	rand.Seed(time.Now().UnixNano())
	words := Read(word)
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	result := words[1]
	return result
}
