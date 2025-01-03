package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetArguments(t *testing.T) {
	for _, c := range []struct {
		parameters []string
		arguments  Arguments
	}{
		{[]string{"file"}, Arguments{File: "file"}},
	} {
		args, err := GetArguments(c.parameters)

		assert.Nil(t, err)
		assert.Equal(t, c.arguments, args)
	}
}
