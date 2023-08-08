package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	var (
		currentFlag bool
		listFlag    bool
		gameName    string
	)

	// Parse the command line. There are short and long names for each
	// option
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage: arstats [OPTION]...

Shows statistics for Aisleriot games played by the current user.

Options:
  -c, --current         Show statistics for the most recently played game (default)
  -l, --list            List the names of all games played
  -g, --game=GAMENAME	Name of game for which statistics are desired

Statistics include:
  - Number of wins
  - Number of losses
  - Total games played
  - Winning percentage
  - Best time
  - Average time
  - Worst time
  - Number of wins to next higher percent
  - Number of losses to next lower percent

  `)
	}
	flag.BoolVar(&currentFlag, "c", true, "Use current game")
	flag.BoolVar(&currentFlag, "current", true, "Use current game")
	flag.BoolVar(&listFlag, "l", false, "List all games played")
	flag.BoolVar(&listFlag, "list", false, "List all games played")
	flag.StringVar(&gameName, "g", "", "Game name")
	flag.StringVar(&gameName, "game", "", "Game name")
	flag.Parse()
}
