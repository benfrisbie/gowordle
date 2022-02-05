package wordle

import (
	"bufio"
	"math/rand"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

// Words is a struct to hold the list of words and solutions
type Words struct {
	solutions []string
	words     map[string]struct{}
}

// NewWords properly initializes a new Words struct
func NewWords(wordPath string, solutionsPath string) *Words {
	w := Words{}

	// ingest words into map to make existence checks faster
	words, err := ingestWordsFromFile(wordPath)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to ingest words")
	}
	w.words = make(map[string]struct{})
	for _, word := range words {
		w.words[word] = struct{}{}
	}

	// ingest solutions
	w.solutions, err = ingestWordsFromFile(solutionsPath)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to ingest solution words")
	}

	return &w
}

func ingestWordsFromFile(path string) ([]string, error) {
	words := make([]string, 0)

	// Open file
	f, err := os.Open(path)
	if err != nil {
		return words, err
	}
	defer f.Close()

	// Scan file line by line adding words to the list
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		words = append(words, word)
	}

	return words, nil
}

// RandomSolution chooses a random word from the solution list
func (w *Words) RandomSolution(random *rand.Rand) string {
	return w.solutions[random.Intn(len(w.solutions))]
}

// Exists checks if a word exists in the list of words
func (w *Words) Exists(word string) bool {
	_, ok := w.words[word]
	return ok
}

// GetAlphabet returns all runes in the alphabet
func GetAlphabet() []rune {
	return []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
}
