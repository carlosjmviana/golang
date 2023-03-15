package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUnmarshaled(t *testing.T) {
	omit := Omit{}
	newOmit, err := omit.GetUnmarshaled()

	assert.NoError(t, err)
	assert.NotNil(t, newOmit)
	assert.NotNil(t, newOmit.PersonShowed)
	assert.Nil(t, newOmit.PersonHidden)
}
