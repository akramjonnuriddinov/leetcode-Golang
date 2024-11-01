package main

import (
	"fmt"
)

func main() {
	words := []string{"go", "is", "a", "programming", "language", "go", "is", "simple"}

	wordLength := make(map[string]int)

	// Store the length of each word
	for _, word := range words {
		wordLength[word] = len(word) // Store the length of the word
	}

	// Print the word and its length
	for word, length := range wordLength {
		fmt.Printf("%s has %d letters\n", word, length)
	}
	fmt.Println(wordLength)
}
