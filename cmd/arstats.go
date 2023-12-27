package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	ar "github.com/philhanna/aisleriot"
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
	pdp, err := ar.NewDataProvider()
	if err != nil {
		log.Fatal(err)
	}

	// Handle the --list option
	if listFlag {
		gameNames := pdp.GameList()
		if gameNames != nil {
			for i, gameName := range pdp.GameList() {
				fmt.Printf("%d: %s\n", i+1, ar.ToDisplayName(gameName))
			}
		} else {
			fmt.Printf("No games have been played\n")
		}
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
		fmt.Printf("No games have been played\n")
		return
	}
	gameName = ar.ToDisplayName(gameName)

	// Print the statistics
	printStatistics(pdp, gameName)
}

// Prints the statistics for the specified game
func printStatistics(pdp *ar.DataProvider, gameName string) {
	sName := ar.ToSectionName(gameName)
	section, ok := pdp.Sections[sName]
	if !ok {
		log.Fatalf("Game %q not found\n", gameName)
	}

	// Get the statistics for this game
	statString := section[ar.StatsKey]
	ps, err := ar.NewStatisticsFromString(statString)
	if err != nil {
		log.Fatal(err)
	}

	// Start forming the list of statistical strings
	parts := make([]string, 0)
	parts = append(parts, "Game name:")
	parts = append(parts, "Number of wins:")
	parts = append(parts, "Number of losses:")
	parts = append(parts, "Total games played:")
	parts = append(parts, "Best time:")
	parts = append(parts, "Average time:")
	parts = append(parts, "Worst time:")
	parts = append(parts, "Winning percentage:")
	parts = append(parts, fmt.Sprintf("Number of wins to %d%%:", ps.Percentage()+1))
	parts = append(parts, fmt.Sprintf("Number of losses to %d%%:", ps.Percentage()-1))

	// Pad them all to the length of the longest part
	parts = padParts(parts)

	// Append the statistic to each line
	parts[0] += fmt.Sprintf(" %s", gameName)
	parts[1] += fmt.Sprintf(" %d", ps.Wins())
	parts[2] += fmt.Sprintf(" %d", ps.Losses())
	parts[3] += fmt.Sprintf(" %d", ps.Total())
	parts[4] += fmt.Sprintf(" %s", secondsToTime(ps.Best()))
	parts[5] += fmt.Sprintf(" %s", secondsToTime(ps.Average()))
	parts[6] += fmt.Sprintf(" %s", secondsToTime(ps.Worst()))
	parts[7] += fmt.Sprintf(" %d%%", ps.Percentage())
	parts[8] += fmt.Sprintf(" %d", ps.WinsToNextHigher())
	parts[9] += fmt.Sprintf(" %d", ps.LossesToNextLower())
	if strings.HasSuffix(parts[9], "-1") {
		parts = parts[:9]
	}
	if strings.HasSuffix(parts[8], "-1") {
		parts = parts[:8]
	}

	// Join parts with newlines and print
	stats := strings.Join(parts, "\n")
	fmt.Println(stats)
}

// padParts pads all the strings to the length of the longest part
func padParts(parts []string) []string {
	maxLen := -1
	for _, s := range parts {
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}
	newParts := make([]string, len(parts))
	for i, s := range parts {
		for len(s) < maxLen {
			s += " "
		}
		newParts[i] = s
	}
	return newParts
}

// secondsToTime converts a number of seconds into a mm:ss string
func secondsToTime(seconds int) string {
	if seconds == 0 {
		return "N/A"
	}
	mm := int(seconds / 60)
	ss := seconds % 60
	return fmt.Sprintf("%02d:%02d", mm, ss)
}
