package hangman

import (
	"io/ioutil"
	"log"
)

func ReadHangman(content string) []string {
	var tmp []byte
	var result []string
	tmp, err := ioutil.ReadFile(content)
	result = SplitHangman(string(tmp))
	if err != nil {
		log.Fatal(err)
	}
	return result
}
