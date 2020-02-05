package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// A function to generate a random number with a random seed.
func generateRanNum(min int, max int) int {
	// Start by getting a new random seed
	rand.Seed(time.Now().UnixNano())

	// Return a random number between min and max
	return rand.Intn(max-min) + min
}

// Define the logic to generate a random word.
func generate(w http.ResponseWriter, r *http.Request) {
	// Declare vowels and consonants
	vowels := [5]string{"a", "e", "i", "o", "u"}
	consonants := [21]string{
		"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n",
		"p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}

	// Declare an empty word variable to hold the word
	word := ""

	// Get a random number for the length of our generated word
	maxLength := generateRanNum(4, 8)

	// Generate the word
	for i := 1; i < maxLength; i++ {
		if i%2 == 0 {
			word += consonants[generateRanNum(0, 21)]
		} else {
			word += vowels[generateRanNum(0, 5)]
		}
	}

	fmt.Fprintf(w, word)
}

// Define a default response for the home route.
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/generate", generate)
	log.Fatal(http.ListenAndServe(":8080", router))
}
