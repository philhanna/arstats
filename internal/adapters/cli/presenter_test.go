package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PadParts(t *testing.T) {
	tests := []struct {
		name  string
		parts []string
		want  []string
	}{
		{"Simple", []string{"Larry", "Curly", "Moe"}, []string{"Larry", "Curly", "Moe  "}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, PadParts(tt.parts))
		})
	}
}

func Test_SecondsToTime(t *testing.T) {
	tests := []struct {
		seconds int
		want    string
	}{
		{0, "N/A"},
		{90, "01:30"},
		{3661, "61:01"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			assert.Equal(t, tt.want, SecondsToTime(tt.seconds))
		})
	}
}
