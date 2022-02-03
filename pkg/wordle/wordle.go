package wordle

import (
	"strings"

	"github.com/benfrisbie/gowordle/pkg/words"
)

type Wordle struct {
	Word       string
	Guesses    int
	MaxGuesses int
	Win        bool
	RuneHints  map[rune]Hint
}

func NewWordle(word string, maxGuesses int) *Wordle {
	w := Wordle{}
	w.Word = word
	w.Guesses = 0
	w.MaxGuesses = maxGuesses
	// initialize rune hints
	w.RuneHints = make(map[rune]Hint)
	for _, r := range words.GetAlphabet() {
		w.RuneHints[r] = HintUnused
	}
	return &w
}

func (w *Wordle) Guess(word string) string {
	w.Guesses++
	hint := ""
	for i, r := range word {
		if word[i] == w.Word[i] {
			w.RuneHints[r] = HintCorrect
			hint += HintCorrect.String(r)
		} else if strings.Contains(w.Word, string(word[i])) {
			w.RuneHints[r] = HintIncorrectPosition
			hint += HintIncorrectPosition.String(r)
		} else {
			w.RuneHints[r] = HintIncorrect
			hint += HintIncorrect.String(r)
		}
	}
	return hint
}

func (w *Wordle) GetRuneHints() string {
	hint := ""
	for _, r := range words.GetAlphabet() {
		hint += w.RuneHints[r].String(r) + ","
	}
	return hint
}

func (w Wordle) IsWin() bool {
	return w.Win
}

func (w Wordle) IsLose() bool {
	return w.Guesses >= w.MaxGuesses
}
