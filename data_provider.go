package aisleriot

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// DataProvider is a structure holding a map of section names to values,
// obtained from the .config/gnome-games/aisleriot .ini file
type DataProvider struct {
	Sections map[string]map[string]string
}

// ---------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------

const (
	HeaderSection = "Aisleriot Config"
	RecentItem    = "Recent"
)

// ---------------------------------------------------------------------
// Variables
// ---------------------------------------------------------------------

var (
	// Create the regular expression(s) we will use to parse the file.
	// Declared outside any function so that they can be unit tested.
	reSection = regexp.MustCompile(`\[(.*)\]`)
	reItem    = regexp.MustCompile(`([^=]+)=([^=]+)`)
)

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewDataProvider reads the specified configuration file, which is in
// .ini format, and returns a pointer to a DataProvider having the
// file's contents parsed into named sections and their lines.
func NewDataProvider(filenames ...string) (*DataProvider, error) {

	// Create a new, empty data provider structure
	pdp := new(DataProvider)

	// Read the specified .ini file
	var filename string
	switch len(filenames) {
	case 0:
		filename = DefaultFileName()
	default:
		filename = filenames[0]
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Parse its contents
	pdp.Sections, err = ParseData(data)
	if err != nil {
		return nil, err
	}

	// Done
	return pdp, nil
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// GameList returns the list of all games played so far, based on the
// "Recent" list in the header section.  If no games have been played,
// returns nil
func (pdp *DataProvider) GameList() []string {
	item := pdp.Sections[HeaderSection][RecentItem]
	if item == "" {
		return nil
	}
	item = strings.TrimSuffix(item, ";")
	list := strings.Split(item, ";")

	return list
}

// MostRecentGame returns the name of the game most recently played,
// or the empty string if no games have been played.
func (pdp *DataProvider) MostRecentGame() string {
	list := pdp.GameList()
	if list == nil {
		return ""
	}
	return list[0]	
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// DefaultFileName returns the name of the .ini file in the user .config
// directory.
func DefaultFileName() string {
	configDir, _ := os.UserConfigDir()
	filename := filepath.Join(configDir, "gnome-games", "aisleriot")
	return filename
}

// ParseData reads the contents of an .ini file and returns a map of its
// section names and their lines.
func ParseData(data []byte) (map[string]map[string]string, error) {
	var (
		err         error
		group       []string
		key         string
		line        string
		lineNumber  int
		scanner     *bufio.Scanner
		sectionName string
		sm          map[string]map[string]string
		value       string
	)

	// Create an empty map of section names to list of strings
	sm = make(map[string]map[string]string)

	// Scan the lines into sections
	scanner = bufio.NewScanner(bytes.NewReader(data))

	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			// Skip blank lines and comments
			continue
		}
		lineNumber++

		// If this line is a section header, make it the current section
		group = reSection.FindStringSubmatch(line)
		if group != nil {
			sectionName = group[1]
			sm[sectionName] = make(map[string]string)
			continue
		} else if lineNumber == 1 {
			// The first non-blank line must be a section header
			err = fmt.Errorf("data found before any section header: %q", line)
			return nil, err
		}

		// Add this line to the current section.
		group = reItem.FindStringSubmatch(line)
		if group == nil {
			// The item does not consist of key=value
			err = fmt.Errorf("invalid item: %q", line)
			return nil, err
		}
		key, value = group[1], group[2]
		sm[sectionName][key] = value

	}
	return sm, nil
}

// ToSectionName converts a game name to the corresponding section name.
// Hyphens are converted to underscores and ".scm" is appended.
func ToSectionName(gameName string) string {
	var sectionName string
	gameName = strings.TrimSpace(gameName)
	if gameName == "" {
		return ""
	}
	sectionName = strings.ReplaceAll(gameName, "-", "_")
	sectionName += ".scm"
	return sectionName
}