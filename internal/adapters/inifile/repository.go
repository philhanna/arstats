package inifile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/philhanna/aisleriot/internal/domain"
	"github.com/philhanna/aisleriot/internal/ports"
)

const (
	headerSection = "Aisleriot Config"
	recentItem    = "Recent"
	statsKey      = "Statistic"
)

// IniRepository implements ports.GameRepository by reading an INI config file.
type IniRepository struct {
	sections map[string]map[string]string
}

// Compile-time check that IniRepository satisfies the port interface.
var _ ports.GameRepository = (*IniRepository)(nil)

// NewIniRepository reads and parses the given INI file (defaults to the
// system Aisleriot config if no filename is provided).
func NewIniRepository(filenames ...string) (*IniRepository, error) {
	filename := DefaultFileName()
	if len(filenames) > 0 {
		filename = filenames[0]
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	sections, err := ParseData(data)
	if err != nil {
		return nil, err
	}
	return &IniRepository{sections: sections}, nil
}

// DefaultFileName returns the path to the Aisleriot config file.
func DefaultFileName() string {
	configDir, _ := os.UserConfigDir()
	return filepath.Join(configDir, "gnome-games", "aisleriot")
}

// GameList returns all game names from the Recent list, or nil if none.
func (r *IniRepository) GameList() []string {
	item := r.sections[headerSection][recentItem]
	if item == "" {
		return nil
	}
	item = strings.TrimSuffix(item, ";")
	return strings.Split(item, ";")
}

// MostRecentGame returns the first game in the Recent list, or "".
func (r *IniRepository) MostRecentGame() string {
	list := r.GameList()
	if list == nil {
		return ""
	}
	return list[0]
}

// GetStatistics returns the Statistics for the given INI section name.
func (r *IniRepository) GetStatistics(sectionName string) (*domain.Statistics, error) {
	section, ok := r.sections[sectionName]
	if !ok {
		return nil, fmt.Errorf("game %q not found", sectionName)
	}
	return domain.NewStatisticsFromString(section[statsKey])
}
