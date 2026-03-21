package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/philhanna/aisleriot/internal/adapters/cli"
	"github.com/philhanna/aisleriot/internal/adapters/inifile"
	"github.com/philhanna/aisleriot/internal/application"
)

func main() {

	var (
		listFlag    bool
		gameNameArg string
	)

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

	// Wire the secondary adapter to the port and build the service.
	repo, err := inifile.NewIniRepository()
	if err != nil {
		log.Fatal(err)
	}
	svc := application.NewGameService(repo)

	// Handle the --list option.
	if listFlag {
		cli.List(svc.ListGames())
		return
	}

	// Resolve the game name.
	gameName := gameNameArg
	if gameName == "" {
		gameName = svc.MostRecentGame()
	}
	if gameName == "" {
		cli.ErrorMessage("No games have been played")
		return
	}

	// Fetch and present statistics.
	stats, err := svc.GetStatistics(gameName)
	if err != nil {
		log.Fatal(err)
	}
	cli.PrintStatistics(gameName, stats)
}
