package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_padParts(t *testing.T) {

	tests := []struct {
		name string
		parts []string
		want []string
	}{
		{"Simple", []string{"Larry", "Curly", "Moe"}, []string{"Larry", "Curly", "Moe"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.parts)
		})
	}
}
