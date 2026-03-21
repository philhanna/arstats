package ports

import "github.com/philhanna/aisleriot/internal/domain"

// GameRepository is the secondary (driven) port for reading game data.
// Implementations are provided by adapters (e.g. the INI file adapter).
type GameRepository interface {
	// GameList returns all game names in the order they appear in the
	// Recent list. Returns nil if no games have been played.
	GameList() []string

	// MostRecentGame returns the raw name of the most recently played
	// game, or the empty string if no games have been played.
	MostRecentGame() string

	// GetStatistics returns the Statistics for the given INI section name
	// (e.g. "spider.scm"). Returns an error if not found or malformed.
	GetStatistics(sectionName string) (*domain.Statistics, error)
}
