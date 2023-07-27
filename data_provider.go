package aisleriot

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// DataProvider is a structure holding a map of section names to values,
// obtained from the .config/gnome-games/aisleriot .ini file
type DataProvider struct {
	Sections map[string][]string
}

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

	// Parse its contentws
	pdp.Sections, err = ParseData(data)
	if err != nil {
		return nil, err
	}

	return pdp, nil
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
func ParseData(data []byte) (map[string][]string, error) {

	// Create an empty map of section names to list of strings
	sm := make(map[string][]string)

	// Scan the lines into sections
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {

	}
	return sm, nil
}