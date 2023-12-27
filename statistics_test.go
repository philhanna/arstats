package aisleriot

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccessors(t *testing.T) {
	ps := NewStatistics(45, 241, 479, 907)
	assert.Equal(t, 45, ps.Wins())
	assert.Equal(t, 196, ps.Losses())
	assert.Equal(t, 241, ps.Total())
	assert.Equal(t, 479, ps.Best())
	assert.Equal(t, 693, ps.Average())
	assert.Equal(t, 907, ps.Worst())
	assert.Equal(t, 3, ps.WinsToNextHigher())
	assert.Equal(t, 3, ps.LossesToNextLower())

	ps = NewStatistics(0, 0, 0, 0)
	assert.Equal(t, 0, ps.Wins())
	assert.Equal(t, 0, ps.Losses())
	assert.Equal(t, 0, ps.Total())
	assert.Equal(t, 0, ps.Best())
	assert.Equal(t, 0, ps.Average())
	assert.Equal(t, 0, ps.Worst())
	assert.Equal(t, -1, ps.WinsToNextHigher())
	assert.Equal(t, -1, ps.LossesToNextLower())
}

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

func TestStatistics_WinsToNextHigher(t *testing.T) {

	tests := []struct {
		name       string
		statString string
		want       int
	}{
		{"0 win 0 loss", "0;0;0;0;", -1},
		{"1 win 0 loss", "1;1;0;0;", -1},
		{"1 win 2 loss", "1;2;0;0;", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps, err := NewStatisticsFromString(tt.statString)
			assert.Nil(t, err)
			want := tt.want
			have := ps.WinsToNextHigher()
			assert.Equal(t, want, have)
		})
	}
}
