package aisleriot

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// Captures the wins, losses, and percentages
type Statistics struct {
	wins    int // Number of wins
	losses  int // Number of losses
	total   int // Total games played
	best    int // Time in seconds
	average int // Time in seconds
	worst   int // Time in seconds
	pct     int // Multiplied by 100 and rounded to nearest integer
}

// ---------------------------------------------------------------------
// Constructors
// ---------------------------------------------------------------------
// Creates a new Statistics object from the basic integer values
// that AisleRiot keeps:
//   - wins
//   - total
//   - best
//   - worst
//
// It then calculates the other three values:
//   - losses
//   - average
//   - percentage of wins
func NewStatistics(wins, total, best, worst int) *Statistics {
	stats := new(Statistics)
	stats.wins = wins
	stats.total = total
	stats.best = best
	stats.worst = worst
	stats.average = int(math.Round(float64(best+worst) / 2.0))
	stats.losses = total - wins
	if total != 0 {
		stats.pct = int(math.Round(100.0 * float64(wins) / float64(total)))
	}
	return stats
}

// Creates a new Statistics object from the string representation
// that is in the configuration file, e.g., "99;150;144;208;"
func NewStatisticsFromString(statString string) (*Statistics, error) {
	statString = strings.TrimSuffix(statString, ";")
	tokens := strings.Split(statString, ";")
	if len(tokens) != 4 {
		return nil, fmt.Errorf("expected 4 values, got %d from %q", len(tokens), statString)
	}
	wins, err := strconv.Atoi(tokens[0])
	if err != nil {
		return nil, fmt.Errorf("invalid 'wins' value: %q", fmt.Sprintf("%v", err))
	}

	total, err := strconv.Atoi(tokens[1])
	if err != nil {
		return nil, fmt.Errorf("invalid 'total' value: %q", fmt.Sprintf("%v", err))
	}

	best, err := strconv.Atoi(tokens[2])
	if err != nil {
		return nil, fmt.Errorf("invalid 'best' value: %q", fmt.Sprintf("%v", err))
	}

	worst, err := strconv.Atoi(tokens[3])
	if err != nil {
		return nil, fmt.Errorf("invalid 'worst' value: %q", fmt.Sprintf("%v", err))
	}

	return NewStatistics(wins, total, best, worst), nil
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

func (ps *Statistics) Wins() int {
	return ps.wins
}

func (ps *Statistics) Losses() int {
	return ps.total - ps.wins
}

func (ps *Statistics) Total() int {
	return ps.total
}

func (ps *Statistics) Best() int {
	return ps.best
}

func (ps *Statistics) Average() int {
	return ps.average
}

func (ps *Statistics) Worst() int {
	return ps.worst
}

func (ps *Statistics) Percentage() int {
	return ps.pct
}

func (ps *Statistics) WinsToNextHigherPercentage() int {
	if ps.Wins() == 0 {
		return -1
	}
	currentPct := ps.Percentage()
	wins, losses := ps.Wins(), ps.Losses()
	for {
		wins++
		nextPct := int(math.Round(100 * float64(wins) / float64(wins+losses)))
		if nextPct > currentPct {
			return wins - ps.Wins()
		}
	}
}

func (ps *Statistics) LossesToNextLowerPercentage() int {
	if ps.Wins() == 0 {
		return -1
	}
	currentPct := ps.Percentage()
	wins, losses := ps.Wins(), ps.Losses()
	for {
		losses++
		nextPct := int(math.Round(100 * float64(wins) / float64(wins+losses)))
		if nextPct < currentPct {
			return losses - ps.Losses()
		}
	}
}
