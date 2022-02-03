package wordle

type Hint int8

const (
	HintCorrect           Hint = iota
	HintIncorrect         Hint = iota
	HintIncorrectPosition Hint = iota
	HintUnused            Hint = iota
)

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
