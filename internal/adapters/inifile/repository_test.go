package inifile

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testdata string

func init() {
	cwd, _ := os.Getwd()
	testdata = filepath.Join(cwd, "..", "..", "..", "testdata")
	testdata, _ = filepath.Abs(testdata)
}

func TestNewIniRepositoryWithBogusFilename(t *testing.T) {
	_, err := NewIniRepository("Bogus")
	assert.NotNil(t, err)
}

func TestNewIniRepositoryDefault(t *testing.T) {
	_, _ = NewIniRepository()
}

func TestDefaultFileName(t *testing.T) {
	userId, err := user.Current()
	username := userId.Username
	assert.Nil(t, err)
	filename := DefaultFileName()
	assert.NotNil(t, filename)
	assert.NotEmpty(t, filename)
	assert.Contains(t, filename, username)
}

func TestParseData(t *testing.T) {
	stoogeFile := filepath.Join(testdata, "stooges.ini")
	stooges, err := os.ReadFile(stoogeFile)
	assert.Nil(t, err)
	tests := []struct {
		name     string
		data     []byte
		expected map[string]map[string]string
	}{
		{"stooges", stooges, map[string]map[string]string{
			"Moe": {
				"rank":   "1",
				"saying": "Why, I oughta...",
			},
			"Larry": {
				"rank":   "2",
				"saying": "Hey, Moe!",
			},
			"Curly": {
				"rank":   "3",
				"saying": "Nyuk, nyuk, nyuk",
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sections, err := ParseData(tt.data)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, sections)
		})
	}
}

func TestIniRepository_GameList(t *testing.T) {
	tests := []struct {
		name          string
		filename      string
		expected      []string
		expectedError bool
	}{
		{"Normal file",
			filepath.Join(testdata, "aisleriot"),
			[]string{"spider", "freecell", "canfield", "klondike"},
			false},
		{"Different .ini",
			filepath.Join(testdata, "stooges.ini"),
			nil,
			false},
		{"Non-existent file",
			filepath.Join(testdata, "non-existent.ini"),
			nil,
			true},
		{"Malformed file",
			filepath.Join(testdata, "bogus.ini"),
			nil,
			true},
		{"Malformed file2",
			filepath.Join(testdata, "bogus2.ini"),
			nil,
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := NewIniRepository(tt.filename)
			if tt.expectedError {
				assert.NotNil(t, err, fmt.Sprint(err))
			} else {
				actual := repo.GameList()
				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}

func TestIniRepository_MostRecentGame(t *testing.T) {
	tests := []struct {
		name          string
		filename      string
		expected      string
		expectedError bool
	}{
		{"Normal file",
			filepath.Join(testdata, "aisleriot"),
			"spider",
			false},
		{"Different .ini",
			filepath.Join(testdata, "stooges.ini"),
			"",
			false},
		{"Non-existent file",
			filepath.Join(testdata, "non-existent.ini"),
			"",
			true},
		{"Malformed file",
			filepath.Join(testdata, "bogus.ini"),
			"",
			true},
		{"Malformed file2",
			filepath.Join(testdata, "bogus2.ini"),
			"",
			true},
		{"Good file 2",
			filepath.Join(testdata, "goodfile.ini"),
			"block-ten",
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := NewIniRepository(tt.filename)
			if tt.expectedError {
				assert.NotNil(t, err, fmt.Sprint(err))
			} else {
				actual := repo.MostRecentGame()
				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}
