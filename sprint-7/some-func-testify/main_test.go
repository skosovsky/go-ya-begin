package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssertExample(t *testing.T) {
	expected := "42"
	result := SomeFunction()
	assert.NotEmpty(t, result)
	assert.Equal(t, result, expected)
}

func TestAssertExampleRequire(t *testing.T) {
	expected := "42"
	result := SomeFunction()
	require.NotEmpty(t, result)
	require.Equal(t, result, expected)
}
