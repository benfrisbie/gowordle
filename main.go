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

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// parse command line flags
	maxGuesses := flag.Int("max_guesses", 6, "max number of guesses to allow")
	wordPath := flag.String("word_path", "words.txt", "path to a text file containing list of all possible words seperated by new lines")
	solutionsPath := flag.String("solutions_path", "solutions.txt", "path to a text file containing list of all solution words seperated by new lines")
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

	log.Debug().Int64("seed", *seed).Msg("initializing game")
	var game *wordle.Wordle

	// setup signal handler
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

	// load words and game
	words := wordle.NewWords(*wordPath, *solutionsPath)
	game = wordle.NewWordle(words.RandomSolution(*seed), *maxGuesses)
	fmt.Printf("Random word selected. You have %d guesses. Ready, Set, Go!\n", *maxGuesses)

	// start game loop
	var reader = bufio.NewReader(os.Stdin)
	for {
		// prompt user for guess
		fmt.Printf("Guess: ")
		guess, _ := reader.ReadString('\n')
		guess = strings.ToLower(strings.TrimSuffix(guess, "\n"))

		// validate guess
		if len(guess) != len(game.Word) {
			fmt.Printf("incorrect length\n")
			continue
		} else if !words.Exists(guess) {
			fmt.Printf("not a real word\n")
			continue
		}

		// send guess to game and print hints
		hint := game.Guess(guess)
		fmt.Println(hint + " - " + game.GetAlphabetHints())

		// Check for end of game
		if game.IsWin() {
			fmt.Printf("You won! You used %d out of %d guesses!\n", game.Guesses, game.MaxGuesses)
			break
		} else if game.IsLose() {
			fmt.Printf("You lost! The word was %v\n", game.Word)
			break
		}
	}

}
