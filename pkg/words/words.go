package words

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
)

type WordDictionary struct {
	words    []string
	wordsMap map[string]struct{}
}

func NewWordDictionary(path string, length int) *WordDictionary {
	ed := WordDictionary{}
	ed.ingestFromFile(path, length)
	return &ed
}

func (ed *WordDictionary) ingestFromFile(path string, length int) error {
	// Open file
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Scan words into map line by line
	ed.words = make([]string, 0)
	ed.wordsMap = make(map[string]struct{})
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		if len(word) == length {
			ed.words = append(ed.words, word)
			ed.wordsMap[word] = struct{}{}
		}
	}
	return nil
}

func (ed *WordDictionary) Random(seed int64) string {
	i := rand.New(rand.NewSource(seed)).Intn(len(ed.words))
	return ed.words[i]
}

func (ed *WordDictionary) Exists(word string) bool {
	_, ok := ed.wordsMap[word]
	return ok
}

// GetAlphabet returns all runes in the alphabet
func GetAlphabet() []rune {
	return []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
}
