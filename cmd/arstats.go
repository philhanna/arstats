package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	ar "github.com/philhanna/aisleriot"
)

func main() {

	var (
		listFlag bool
		gameName string
	)

	// Parse the command line. There are short and long names for each
	// option
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			`Usage: arstats [OPTION]...

Shows statistics for Aisleriot games played by the current user.

Options:
  -g, --game=GAMENAME	Name of game for which statistics are desired
                        (Default is most recently played game)
  -l, --list            List the names of all games played

Output includes:
  - Game name
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
	flag.BoolVar(&listFlag, "l", false, "List all games played")
	flag.BoolVar(&listFlag, "list", false, "List all games played")
	flag.StringVar(&gameName, "g", "", "Game name")
	flag.StringVar(&gameName, "game", "", "Game name")
	flag.Parse()

	// Get the data provider
	pdp, err := ar.NewDataProvider()
	if err != nil {
		log.Fatal(err)
	}

	// Handle the --list option
	if listFlag {
		gameNames := pdp.GameList()
		if gameNames == nil {
			fmt.Printf("No games have been played\n")
		} else {
			for i, gameName := range pdp.GameList() {
				fmt.Printf("%d: %s\n", i, gameName)
			}
			return
		}
	}
}
