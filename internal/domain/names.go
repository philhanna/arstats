package domain

import "strings"

// ToDisplayName converts a raw game name (as found in the config file's
// Recent list) into a human-readable display name.
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

// ToSectionName converts a display game name to its INI section name.
// Spaces and hyphens are converted to underscores and ".scm" is appended.
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

// titleCase makes the first character uppercase and the remainder lowercase.
func titleCase(name string) string {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return name
	}
	if len(name) == 1 {
		return strings.ToUpper(name)
	}
	return strings.ToUpper(name[:1]) + strings.ToLower(name[1:])
}
