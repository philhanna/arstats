package application

import (
	"github.com/philhanna/aisleriot/internal/domain"
	"github.com/philhanna/aisleriot/internal/ports"
)

// GameService implements the application use cases, orchestrating the
// domain and secondary ports. It has no knowledge of I/O or presentation.
type GameService struct {
	repo ports.GameRepository
}

// NewGameService creates a GameService backed by the given repository.
func NewGameService(repo ports.GameRepository) *GameService {
	return &GameService{repo: repo}
}

// ListGames returns the display names of all games played, in order.
// Returns nil if no games have been played.
func (s *GameService) ListGames() []string {
	raw := s.repo.GameList()
	if raw == nil {
		return nil
	}
	names := make([]string, len(raw))
	for i, g := range raw {
		names[i] = domain.ToDisplayName(g)
	}
	return names
}

// MostRecentGame returns the display name of the most recently played
// game, or the empty string if no games have been played.
func (s *GameService) MostRecentGame() string {
	raw := s.repo.MostRecentGame()
	if raw == "" {
		return ""
	}
	return domain.ToDisplayName(raw)
}

// GetStatistics returns the Statistics for the given display-name game.
func (s *GameService) GetStatistics(displayName string) (*domain.Statistics, error) {
	sectionName := domain.ToSectionName(displayName)
	return s.repo.GetStatistics(sectionName)
}
