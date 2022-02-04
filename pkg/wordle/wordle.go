package wordle

import (
	"strings"
)

// Wordle represents a wordle game and the necessary state
type Wordle struct {
	Word       string
	Guesses    int
	MaxGuesses int
	Win        bool
	RuneHints  map[rune]Hint
}

// NewWordle properly initializes a new Wordle game
func NewWordle(word string, maxGuesses int) *Wordle {
	w := Wordle{}
	w.Word = word
	w.Guesses = 0
	w.MaxGuesses = maxGuesses

	// initialize rune hints
	w.RuneHints = make(map[rune]Hint)
	for _, r := range GetAlphabet() {
		w.RuneHints[r] = HintUnused
	}

	return &w
}

// Guess validates a guess against the solution and returns the hints for the guess
func (w *Wordle) Guess(guess string) string {
	w.Guesses++
	hint := ""
	for i, r := range guess {
		if guess[i] == w.Word[i] {
			w.RuneHints[r] = HintCorrect
			hint += HintCorrect.String(r)
		} else if strings.Contains(w.Word, string(guess[i])) {
			w.RuneHints[r] = HintIncorrectPosition
			hint += HintIncorrectPosition.String(r)
		} else {
			w.RuneHints[r] = HintIncorrect
			hint += HintIncorrect.String(r)
		}
	}
	if guess == w.Word {
		w.Win = true
	}
	return hint
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
