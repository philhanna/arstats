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
