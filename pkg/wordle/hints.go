package wordle

// Hint indicates the state of a character compared to the solution
type Hint int8

const (
	HintCorrect           Hint = iota // Character is in the correct position
	HintIncorrect         Hint = iota // Character is not in the solution
	HintIncorrectPosition Hint = iota // Character is in the wrong position, but is in the solution
	HintUnused            Hint = iota // Character has not been guessed yet
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
	case HintUnused:
		return "\033[37m" + string(r) + "\033[0m"
	default:
		return string(r)
	}
}
