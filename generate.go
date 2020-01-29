package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRanNum(min int, max int) int {
	// Start by getting a new random seed
	rand.Seed(time.Now().UnixNano())
	// Return a random number between min and max
	return rand.Intn(max-min) + min
}

func generate() string {
	// Declare vowels and consonants
	var vowels = [5]string{"a", "e", "i", "o", "u"}
	var consonants = [21]string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}

	// Declare an empty word variable to hold the word
	var word = ""

	// Get a random number for the length of our generated word
	var maxLength = generateRanNum(4, 8)

	// Generate the word
	for i := 1; i < maxLength; i++ {
		if i%2 == 0 {
			word += consonants[generateRanNum(0, 21)]
		} else {
			word += vowels[generateRanNum(0, 5)]
		}
	}

	return word
}

func main() {

	// Generate a word
	fmt.Println(generate())

}
