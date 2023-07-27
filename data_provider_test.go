package aisleriot

import (
	"os/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultFileName(t *testing.T) {
	userId, err := user.Current()
	username := userId.Username
	assert.Nil(t, err)
	filename := DefaultFileName()
	assert.NotNil(t, filename)
	assert.NotEmpty(t, filename)
	assert.Contains(t, filename, username)
}

func TestParseData(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    map[string][]string
	}{
		{"stooges", []byte(`
[Larry]
rank=2
saying=Hey, Moe!

[Curly]
rank=3
saying=Nyuk, nyuk, nyuk

[Moe]
rank=1
saying=Why, I oughta...
`), map[string][]string{
	"Moe": {
		"rank=1",
		"saying=Why, I oughta...",
	},
	"Larry": {
		"rank=2",
		"saying=Hey, Moe!",
	},
	"Curly": {
		"rank=3",
		"saying=Nyuk, nyuk, nyuk",
	},
}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sections := ParseData(tt.data)
			assert.Equal(t, tt.want, sections)
		})
	}
}
