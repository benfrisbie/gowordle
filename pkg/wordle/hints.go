package wordle

// HintStatus indicates the state of a character compared to the solution
type HintState int8

const (
	HintUnknown           HintState = iota // Character has not been guessed yet
	HintIncorrect         HintState = iota // Character is not in the solution
	HintIncorrectPosition HintState = iota // Character is in the wrong position, but is in the solution
	HintCorrect           HintState = iota // Character is in the correct position
)

type Hints struct {
	BackgroundColor bool
	AlphabetHintMap map[rune]HintState
}

func NewHints(backgroundColor bool) Hints {
	h := Hints{}
	h.BackgroundColor = backgroundColor

	h.AlphabetHintMap = make(map[rune]HintState)
	for _, r := range GetAlphabet() {
		h.AlphabetHintMap[r] = HintUnknown
	}

	return h
}

// UpdateHintMap updates the hint map the given character only if the new hint is better than the current
func (h Hints) UpdateAlphabetHintMap(r rune, new HintState) {
	if new > h.AlphabetHintMap[r] {
		h.AlphabetHintMap[r] = new
	}
}

// String returns the string representation of the Hint for a character
func (h Hints) String(r rune, hs HintState) string {
	switch h.BackgroundColor {
	case true:
		switch hs {
		case HintIncorrect:
			return "\033[41m" + string(r) + "\033[0m"
		case HintIncorrectPosition:
			return "\033[43m" + string(r) + "\033[0m"
		case HintCorrect:
			return "\033[42m" + string(r) + "\033[0m"
		}

	case false:
		switch hs {
		case HintIncorrect:
			return "\033[31m" + string(r) + "\033[0m"
		case HintIncorrectPosition:
			return "\033[33m" + string(r) + "\033[0m"
		case HintCorrect:
			return "\033[32m" + string(r) + "\033[0m"
		}
	}
	return string(r)
}
