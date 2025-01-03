package main

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	f, err := os.CreateTemp("", "")
	assert.Nil(t, err)

	_, err = f.WriteString("Feature: Foo")
	assert.Nil(t, err)

	assert.Nil(t, Run([]string{f.Name()}, io.Discard))

	os.Remove(f.Name())
}

func TestCommandWithNonExistentFile(t *testing.T) {
	assert.NotNil(t, Run([]string{"non-existent.feature"}, io.Discard))
}
