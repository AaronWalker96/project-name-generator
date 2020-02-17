package main

import "testing"

func TestGenerateRanNum(t *testing.T) {
	sum := 1

	// Generate a random number 10 times, if any of the numbers
	// are outside the specified bounds, fail the test.
	for sum < 10 {
		sum += sum
		randomNumber := generateRanNum(0, 10)
		if randomNumber < 0 {
			t.Errorf("Randomly generated number is smaller than the specified bounds.")
		} else if randomNumber > 10 {
			t.Errorf("Randomly generated number is larger than the specified bounds.")
		}

	}
}

func TestGenerateWord(t *testing.T) {
	minWord := 4
	maxWord := 8
	sum := 1

	// Generate a random number 10 times, if any of the numbers
	// are outside the specified bounds, fail the test.
	for sum < 10 {
		sum += sum
		generatedWord := generateWord(minWord, maxWord)
		if len(generatedWord)+1 < minWord {
			t.Errorf("Randomly generated word is shorter than the specified bounds. Was %d characters", len(generatedWord))
		} else if len(generatedWord)+1 > maxWord {
			t.Errorf("Randomly generated word is longer than the specified bounds. Was %d characters", len(generatedWord))
		}
	}

}
