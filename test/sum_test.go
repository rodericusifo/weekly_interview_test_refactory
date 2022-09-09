package test

import (
	"example/test/src"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	inputX := 12
	inputY := 2
	output := 14
	res := src.Sum(inputX, inputY)
	assert.Equal(t, output, res)
}