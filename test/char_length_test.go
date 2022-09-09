package test

import (
	"testing"
  "github.com/stretchr/testify/assert"

	"example/test/src"
)

func TestCharLength(t *testing.T) {
	input := "hello"
	output := 5
	res := src.CharLength(input)
	assert.Equal(t, output, res, "Not Same")
}

