package model

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
	StatsKey      = "Statistic"
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
		lineNumber  int
		sectionName string
		sm          = make(map[string]map[string]string)
	)

	scanner := bufio.NewScanner(bytes.NewReader(data))

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue // skip blank lines and comments
		}
		lineNumber++

		if sectionName == "" {
			// The first non-blank line must be a section header
			sectionName, err = parseSection(line)
			if err != nil {
				return nil, err
			}
			sm[sectionName] = make(map[string]string)
			continue
		}

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			// This line is a section header
			sectionName, _ = parseSection(line)
			sm[sectionName] = make(map[string]string)
			continue
		}

		key, value, err := parseItem(line)
		if err != nil {
			return nil, err
		}
		sm[sectionName][key] = value
	}
	return sm, nil
}

// parseItem parses an item and returns its key and value.
func parseItem(line string) (string, string, error) {
	group := reItem.FindStringSubmatch(line)
	if group == nil {
		return "", "", fmt.Errorf("invalid item: %q", line)
	}
	return group[1], group[2], nil
}

// parseSection parses a section header and returns its name.
func parseSection(line string) (string, error) {
	group := reSection.FindStringSubmatch(line)
	if group == nil {
		return "", fmt.Errorf("invalid section header: %q", line)
	}
	return group[1], nil
}

// ToDisplayName converts a game name as found in the [AisleRiot Config]
// Recent=name1;name2;name3 item into a name suitable for display.
func ToDisplayName(gameName string) string {
	gameName = strings.ReplaceAll(gameName, ".scm", "")
	gameName = strings.ReplaceAll(gameName, "-", " ")
	gameName = strings.ReplaceAll(gameName, "_", " ")
	names := strings.Fields(gameName)
	if len(names) == 0 {
		return ""
	}
	parts := []string{}
	for _, name := range names {
		parts = append(parts, titleCase(name))
	}
	return strings.Join(parts, " ")
}

// ToSectionName converts a game name to the corresponding section name.
// Hyphens are converted to underscores and ".scm" is appended.
func ToSectionName(gameName string) string {
	sName := strings.TrimSpace(gameName)
	if sName != "" {
		sName = strings.ToLower(sName)
		sName = strings.ReplaceAll(sName, " ", "_")
		sName = strings.ReplaceAll(sName, "-", "_")
		sName += ".scm"
	}
	return sName
}

// titleCase makes the first character of a name uppercase, and the
// remainder (if any) lower case
func titleCase(name string) string {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return name
	}
	if len(name) == 1 {
		return strings.ToUpper(name)
	}
	prefix := name[:1]
	suffix := name[1:]
	return strings.ToUpper(prefix) + strings.ToLower(suffix)
}
