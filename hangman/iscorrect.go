package hangman

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Replace(word string) []rune {
	hangmanStatus := ReadHangman("words/hangman.txt")
	health := 10
	word = Randomize(word)
	wordToGuess := WordToGuess(word)
	var tried string
	clearscreen()
	for hp := health; hp > 0; {
		HP := 10 - hp
		fmt.Println(hangmanStatus[HP])
		fmt.Println("Il vous reste", hp, "essais.")
		if word != string(wordToGuess) {
			fmt.Printf("Votre progression : %c\n\n", wordToGuess)
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("Entrez une lettre : ")
			guess, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			tried += guess
			tried = Tried(tried)
			fmt.Printf("Caractères et mots essayés : %s\n", tried)
			if IsValidGuess(wordToGuess) {
				if WordGuessed(guess, word, wordToGuess) {
					clearscreen()
					fmt.Println(hangmanStatus[HP])
					fmt.Println("Bravo, vous avez gagné en ayant commis", HP, "erreurs ! Le mot à deviner était :", word)
					return wordToGuess
				}
				if len(guess) > 2 && guess != word {
					hp = hp - 2
					fmt.Println("Dommage !")
				}
				if len(guess) == 2 {
					tmp := string(wordToGuess)
					for c, j := range word {
						for _, k := range guess {
							if k == j {
								wordToGuess[c] = j
								if word == string(wordToGuess) {
									clearscreen()
									fmt.Println(hangmanStatus[HP])
									fmt.Println("Bravo, vous avez gagné en ayant commis", HP, "erreurs ! Le mot à deviner était :", word)
									return wordToGuess
								}
							}
						}
					}
					if tmp == string(wordToGuess) {
						hp = hp - 1
						fmt.Println("Dommage !")
					} else {
						fmt.Println("Bien joué !")
					}
				}
			} else {
				return wordToGuess
			}
		}
	}
	fmt.Println("Vous avez perdu ! Le mot à deviner était :", word)
	return wordToGuess
}
