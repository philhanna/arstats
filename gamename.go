package aisleriot

import "strings"

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// ToDisplayName converts a game name as found in the [AisleRiot Config]
// Recent=name1;name2;name3 item into a name suitable for display.
func ToDisplayName(gameName string) string {
	return "-1"
}

// ToSectionName converts a game name to the corresponding section name.
// Hyphens are converted to underscores and ".scm" is appended.
func ToSectionName(gameName string) string {
	sName := strings.TrimSpace(gameName)
	if sName != "" {
		sName = strings.ReplaceAll(sName, "-", "_")
		sName += ".scm"
	}
	return sName
}

