package aisleriot

import "strings"

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

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
		sName = strings.ReplaceAll(sName, "-", "_")
		sName += ".scm"
	}
	return sName
}
