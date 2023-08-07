package aisleriot

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStatisticsFromString(t *testing.T) {
	tests := []struct {
		name          string
		line          string
		expected      []int
		expectedError bool
		expectedWord  string
	}{
		{"current", "45;241;479;907;", []int{45, 241, 479, 907}, false, ""},
		{"no games", "0;0;0;0;", []int{0, 0, 0, 0}, false, ""},
		{"1 loss", "0;1;0;0;", []int{0, 1, 0, 0}, false, ""},
		{"too few", "0;1;300;", nil, true, "got 3"},
		{"too many", "0;1;300;45;241;479;907;", nil, true, "got 7"},
		{"bad wins", "bogus;1;1;1;", nil, true, "wins"},
		{"bad total", "1;bogus;1;1;", nil, true, "total"},
		{"bad best", "1;1;bogus;1;", nil, true, "best"},
		{"bad worst", "1;1;1;bogus;", nil, true, "worst"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := NewStatisticsFromString(tt.line)
			if tt.expectedError {
				errmsg := fmt.Sprintf("%v", err)
				assert.Contains(t, errmsg, tt.expectedWord)
			} else {
				assert.Nil(t, err, err)
				expected := tt.expected
				assert.Equal(t, expected[0], actual.wins)
				assert.Equal(t, expected[1], actual.total)
				assert.Equal(t, expected[2], actual.best)
				assert.Equal(t, expected[3], actual.worst)
			}
		})
	}
}
