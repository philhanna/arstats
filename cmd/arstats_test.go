package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// ar "github.com/philhanna/aisleriot"
)

func Test_getDataProvider(t *testing.T) {
	pdp, err := getDataProvider()
	assert.Nil(t, err)
	assert.NotNil(t, pdp)
}
