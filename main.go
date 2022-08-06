package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
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
	random := rand.New(rand.NewSource(*seed))
	var game *wordle.Wordle

	// load words
	words := wordle.NewWords(*wordPath, *solutionsPath)
	var reader = bufio.NewReader(os.Stdin)

	for {
		game = wordle.NewWordle(words.RandomSolution(random), *maxGuesses)

		// start game loop
		for {
			// prompt user for guess
			fmt.Printf("Guess #%d/%d: ", game.Guesses+1, *maxGuesses)
			guess, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal().Err(err).Msg("error reading input")
			}
			guess = strings.ToLower(strings.TrimSuffix(guess, "\n"))

			// validate guess
			if len(guess) != len(game.Solution) {
				fmt.Printf("Incorrect length. Enter a %d letter word\n", len(game.Solution))
				continue
			} else if !words.Exists(guess) {
				fmt.Printf("Not a real word\n")
				continue
			}

			// send guess to game and print hints
			hint := game.Guess(guess)
			fmt.Println(hint + " - " + game.GetAlphabetHints())

			// Check for end of game
			if game.IsWin() {
				fmt.Printf("You won! Correctly guessed \"%s\" in %d/%d guesses!\n", game.Solution, game.Guesses, game.MaxGuesses)
				break
			} else if game.IsLose() {
				fmt.Printf("You lost! The solution was \"%s\"\n", game.Solution)
				break
			}
		}

		fmt.Printf("Would you like to play again? (y/N)\n")
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal().Err(err).Msg("error reading input")
		}
		response = strings.ToLower(strings.TrimSpace(response))
		if response == "y" || response == "yes" {
			fmt.Println()
			continue
		} else {
			break
		}
	}

}
