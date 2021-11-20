package hangman

import (
	"os"
)

func ReadArgs() string {
	arg := os.Args[1]
	return arg
}
