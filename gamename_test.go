package aisleriot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestToSectionName(t *testing.T) {
	tests := []struct {
		name     string
		gameName string
		expected string
	}{
		{"simple", "freecell", "freecell.scm"},
		{"with hyphen", "auld-lang-syne", "auld_lang_syne.scm"},
		{"empty", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ToSectionName(tt.gameName)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestToDisplayName(t *testing.T) {
	tests := []struct {
		name     string
		gameName string
		expected string
	}{
		{"simple", "freecell", "Freecell"},
		{"with suffix", "freecell.scm", "Freecell"},
		{"with hyphen", "auld-lang-syne", "Auld Lang Syne"},
		{"empty", "", ""},
		{"single letters", "a-short-name.scm", "A Short Name"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ToDisplayName(tt.gameName)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

