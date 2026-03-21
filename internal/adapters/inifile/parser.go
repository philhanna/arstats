package inifile

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

var (
	reSection = regexp.MustCompile(`\[(.*)\]`)
	reItem    = regexp.MustCompile(`([^=]+)=([^=]+)`)
)

// ParseData reads raw INI file bytes and returns a map of section names
// to their key-value pairs.
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
			continue
		}
		lineNumber++

		if sectionName == "" {
			sectionName, err = parseSection(line)
			if err != nil {
				return nil, err
			}
			sm[sectionName] = make(map[string]string)
			continue
		}

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
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
	_ = lineNumber
	return sm, nil
}

func parseSection(line string) (string, error) {
	group := reSection.FindStringSubmatch(line)
	if group == nil {
		return "", fmt.Errorf("invalid section header: %q", line)
	}
	return group[1], nil
}

func parseItem(line string) (string, string, error) {
	group := reItem.FindStringSubmatch(line)
	if group == nil {
		return "", "", fmt.Errorf("invalid item: %q", line)
	}
	return group[1], group[2], nil
}
