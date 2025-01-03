package main

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertFile(t *testing.T) {
	f, err := os.CreateTemp("", "")
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	_, err = f.Write([]byte("Feature: Foo"))
	assert.Nil(t, err)

	assert.Nil(t, ConvertFile(f.Name(), io.Discard))
}

func TestConvertFileError(t *testing.T) {
	f, err := os.CreateTemp("", "")
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	_, err = f.Write([]byte("Feature"))
	assert.Nil(t, err)

	assert.NotNil(t, ConvertFile(f.Name(), io.Discard))
}
