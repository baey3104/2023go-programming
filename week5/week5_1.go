package main

import (
	"fmt"
	"strings"
)

func main() {
	brokenWords := "cs r?cks~"
	replaceWords := strings.NewReplacer("?", "o")
	fixedWords := replaceWords.Replace(brokenWords)
	fmt.Println((fixedWords))
}
