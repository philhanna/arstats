package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/philhanna/aisleriot/model"
	"github.com/philhanna/aisleriot/view"
)

func main() {

	var (
		listFlag    bool
		gameNameArg string
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
  - Best time
  - Average time
  - Worst time
  - Winning percentage
  - Number of wins to next higher percent
  - Number of losses to next lower percent
  `)
	}
	flag.BoolVar(&listFlag, "l", false, "List all games played")
	flag.BoolVar(&listFlag, "list", false, "List all games played")
	flag.StringVar(&gameNameArg, "g", "", "Game name")
	flag.StringVar(&gameNameArg, "game", "", "Game name")
	flag.Parse()

	// Get the data provider
	pdp, err := model.NewDataProvider()
	if err != nil {
		log.Fatal(err)
	}

	// Handle the --list option
	if listFlag {
		view.List(pdp)
		return
	}

	// Handle the --game option
	gameName := ""
	if gameNameArg == "" {
		gameName = pdp.MostRecentGame()
	} else {
		gameName = gameNameArg
	}
	if gameName == "" {
		view.ErrorMessage("No games have been played\n")
		return
	}
	gameName = model.ToDisplayName(gameName)

	// Print the statistics
	view.PrintStatistics(pdp, gameName)
}
