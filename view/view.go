package view

import (
	"fmt"
	"github.com/philhanna/aisleriot/model"
	"log"
	"strings"
)

func ErrorMessage(msg string) {
	fmt.Println(msg)
}

func List(pdp *model.DataProvider) {
	gameNames := pdp.GameList()
	if gameNames != nil {
		for i, gameName := range pdp.GameList() {
			fmt.Printf("%d: %s\n", i+1, model.ToDisplayName(gameName))
		}
	} else {
		fmt.Printf("No games have been played\n")
	}

}

// Prints the statistics for the specified game
func PrintStatistics(pdp *model.DataProvider, gameName string) {
	sName := model.ToSectionName(gameName)
	section, ok := pdp.Sections[sName]
	if !ok {
		log.Fatalf("Game %q not found\n", gameName)
	}

	// Get the statistics for this game
	statString := section[model.StatsKey]
	ps, err := model.NewStatisticsFromString(statString)
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
	parts = PadParts(parts)

	// Append the statistic to each line
	parts[0] += fmt.Sprintf(" %s", gameName)
	parts[1] += fmt.Sprintf(" %d", ps.Wins())
	parts[2] += fmt.Sprintf(" %d", ps.Losses())
	parts[3] += fmt.Sprintf(" %d", ps.Total())
	parts[4] += fmt.Sprintf(" %s", SecondsToTime(ps.Best()))
	parts[5] += fmt.Sprintf(" %s", SecondsToTime(ps.Average()))
	parts[6] += fmt.Sprintf(" %s", SecondsToTime(ps.Worst()))
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

// PadParts pads all the strings to the length of the longest part
func PadParts(parts []string) []string {
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

// SecondsToTime converts a number of seconds into a mm:ss string
func SecondsToTime(seconds int) string {
	if seconds == 0 {
		return "N/A"
	}
	mm := int(seconds / 60)
	ss := seconds % 60
	return fmt.Sprintf("%02d:%02d", mm, ss)
}
