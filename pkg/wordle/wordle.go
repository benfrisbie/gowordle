package wordle

import (
	"strings"
)

// Wordle represents a wordle game and the necessary state
type Wordle struct {
	Solution   string
	Guesses    int
	MaxGuesses int
	Win        bool
	RuneHints  map[rune]Hint
}

// NewWordle properly initializes a new Wordle game
func NewWordle(solution string, maxGuesses int) *Wordle {
	w := Wordle{}
	w.Solution = solution
	w.Solution = "exits"
	w.Guesses = 0
	w.MaxGuesses = maxGuesses

	// initialize rune hints
	w.RuneHints = NewHintMap()

	return &w
}

// Guess validates a guess against the solution and returns the hints for the guess
func (w *Wordle) Guess(guess string) string {
	w.Guesses++
	hint := ""
	for i, guessRune := range guess {
		solutionRune := rune(w.Solution[i])
		if guessRune == solutionRune {
			w.UpdateHintMap(guessRune, HintCorrect)
			hint += HintCorrect.String(guessRune)
		} else if count := strings.Count(w.Solution, string(guessRune)); count > 0 && count != w.getCorrectCountForLetter(guess, guessRune) {
			w.UpdateHintMap(guessRune, HintIncorrectPosition)
			hint += HintIncorrectPosition.String(guessRune)
		} else {
			w.UpdateHintMap(guessRune, HintIncorrect)
			hint += HintIncorrect.String(guessRune)
		}
	}
	if guess == w.Solution {
		w.Win = true
	}
	return hint
}

// getCorrectCountForLetter returns the number of times a letter is in the correct position in the solution
func (w *Wordle) getCorrectCountForLetter(guess string, r rune) int {
	count := 0
	for i, solutionRune := range w.Solution {
		if solutionRune == r && rune(guess[i]) == r {
			count++
		}
	}
	return count
}

// GetAlphabetHints returns the current games hints for all letters in the alphabet
func (w *Wordle) GetAlphabetHints() string {
	hint := ""
	for _, r := range GetAlphabet() {
		hint += w.RuneHints[r].String(r) + ","
	}
	return hint
}

// IsWin returns true if the game has been won
func (w Wordle) IsWin() bool {
	return w.Win
}

// IsLose returns true if the game has been lost
func (w Wordle) IsLose() bool {
	return w.Guesses >= w.MaxGuesses
}
