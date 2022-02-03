package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/benfrisbie/gowordle/pkg/wordle"
	"github.com/benfrisbie/gowordle/pkg/words"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// parse command line flags
	length := flag.Int("length", 5, "length of word to guess")
	maxGuesses := flag.Int("max_guesses", 6, "max number of guesses to allow")
	dictionaryPath := flag.String("dictionary_path", "words.txt", "path to a text file containing list of words seperated by new lines")
	seed := flag.Int64("seed", time.Now().UnixNano(), "seed for the random word generator")
	debug := flag.Bool("debug", false, "enable debug logging")
	flag.Parse()

	// setup logging
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	log.Debug().Int64("seed", *seed).Send()

	dictionary := words.NewWordDictionary(*dictionaryPath, *length)

	var reader = bufio.NewReader(os.Stdin)

	var game *wordle.Wordle

	// Setup signal handler
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		if game != nil {
			fmt.Printf("The correct word was %s\n", game.Word)
		}
		os.Exit(0)
	}()

	game = wordle.NewWordle(dictionary.Random(*seed), *maxGuesses)
	fmt.Printf("Random word of length %d generated.\n", *length)

	for {
		fmt.Printf("Guess: ")
		guess, _ := reader.ReadString('\n')
		guess = strings.ToLower(strings.TrimSuffix(guess, "\n"))
		if len(guess) != *length {
			fmt.Printf("incorrect length\n")
			continue
		} else if !dictionary.Exists(guess) {
			fmt.Printf("not a real word\n")
			continue
		}
		hint := game.Guess(guess)
		fmt.Println(hint + " - " + game.GetRuneHints())

		if game.IsWin() {
			fmt.Printf("You won in %d out of %d guesses!\n", game.Guesses, game.MaxGuesses)
			break
		} else if game.IsLose() {
			fmt.Printf("You lost! The word was %v\n", game.Word)
			break
		}
	}

}
