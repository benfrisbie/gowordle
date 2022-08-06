package wordle

// Hint indicates the state of a character compared to the solution
type Hint int8

const (
	HintCorrect           Hint = iota // Character is in the correct position
	HintIncorrectPosition Hint = iota // Character is in the wrong position, but is in the solution
	HintIncorrect         Hint = iota // Character is not in the solution
	HintUnknown           Hint = iota // Character has not been guessed yet
)

// String returns the string representation of the Hint for a character
func (h Hint) String(r rune) string {
	switch h {
	case HintCorrect:
		return "\033[32m" + string(r) + "\033[0m"
	case HintIncorrect:
		return "\033[31m" + string(r) + "\033[0m"
	case HintIncorrectPosition:
		return "\033[33m" + string(r) + "\033[0m"
	case HintUnknown:
		return "\033[37m" + string(r) + "\033[0m"
	default:
		return string(r)
	}
}

// NewHintMap returns a new map of rune hints for the whole alphabet
func NewHintMap() map[rune]Hint {
	hints := make(map[rune]Hint)
	for _, r := range GetAlphabet() {
		hints[r] = HintUnknown
	}
	return hints
}

// UpdateHintMap updates the hint map the given character only if the new hint is better than the current
func (w Wordle) UpdateHintMap(r rune, new Hint) {
	if new < w.RuneHints[r] {
		w.RuneHints[r] = new
	}
}
