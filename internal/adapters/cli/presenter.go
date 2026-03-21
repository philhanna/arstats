package cli

import (
	"fmt"
	"strings"

	"github.com/philhanna/aisleriot/internal/domain"
)

// ErrorMessage prints an error message to stdout.
func ErrorMessage(msg string) {
	fmt.Println(msg)
}

// List prints a numbered list of game display names.
func List(displayNames []string) {
	if len(displayNames) == 0 {
		fmt.Println("No games have been played")
		return
	}
	for i, name := range displayNames {
		fmt.Printf("%d: %s\n", i+1, name)
	}
}

// PrintStatistics prints the formatted statistics table for a game.
func PrintStatistics(displayName string, ps *domain.Statistics) {
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

	parts = PadParts(parts)

	parts[0] += fmt.Sprintf(" %s", displayName)
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

	fmt.Println(strings.Join(parts, "\n"))
}

// PadParts pads all strings to the length of the longest one.
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

// SecondsToTime converts seconds to a "mm:ss" string, or "N/A" if zero.
func SecondsToTime(seconds int) string {
	if seconds == 0 {
		return "N/A"
	}
	return fmt.Sprintf("%02d:%02d", seconds/60, seconds%60)
}
