package hangman

import (
	"io/ioutil"
	"log"
)

func Read(content string) []string {
	var tmp []byte
	var result []string
	tmp, err := ioutil.ReadFile(content)
	result = SplitStrings(string(tmp))
	if err != nil {
		log.Fatal(err)
	}
	return result
}
