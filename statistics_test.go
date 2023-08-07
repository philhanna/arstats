package aisleriot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStatisticsFromString(t *testing.T) {
	tests := []struct {
		name          string
		line          string
		expected      []int
		expectedError bool
	}{
		{"current", "45;241;479;907;", []int{45, 241, 479, 907}, false},
		{"no games", "0;0;0;0;", []int{0, 0, 0, 0}, false},
		{"1 loss", "0;1;0;0;", []int{0, 1, 0, 0}, false},
		{"too few", "0;1;300;", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := NewStatisticsFromString(tt.line)
			if tt.expectedError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				expected := tt.expected
				assert.Equal(t, expected[0], actual.wins)
				assert.Equal(t, expected[1], actual.total)
				assert.Equal(t, expected[2], actual.best)
				assert.Equal(t, expected[3], actual.worst)
			}
		})
	}
}
