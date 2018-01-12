package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(bandNameGenerator("Dolpine"))
	fmt.Println(bandNameGenerator("tart"))
}

func bandNameGenerator(word string) string {
	// Happy coding
	word1 := strings.Title(word)
	word2 := strings.ToLower(word)
	if (word2[0]) == (word2[len(word2)-1]) {
		word3 := word2[1:len(word2)]
		return "The " + word1 + word3
	}
	return "The " + word1
}
