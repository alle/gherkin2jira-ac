package main

import (
	"fmt"
	"strings"
	"testing"

	gherkin "github.com/cucumber/gherkin/go/v27"
	"github.com/stretchr/testify/assert"
)

func TestNewRenderer(t *testing.T) {
	NewRenderer()
}

func TestRendererRender(t *testing.T) {
	for _, ss := range [][2]string{
		{
			"Feature: Foo",
			"**Feature: Foo**\n\n",
		},
		{`
Feature: Foo
  Scenario: Bar
    Given that
    When I do something
    Then something happens`, `
**Feature: Foo**

**Scenario: Bar**

**Given** that
**When** I do something
**Then** something happens`,
		},
		{`
Feature: Foo
  Scenario: Bar
    When I do something:
    """sh
    foo
    """`, fmt.Sprintf(`
**Feature: Foo**

**Scenario: Bar**

**When** I do something:

%[1]ssh
foo
%[1]s`, "```"),
		},
		{`
Feature: Foo
  bar`, `
**Feature: Foo**

bar`,
		},
		{`
Feature: Foo
  Scenario: Bar

    baz`, `
**Feature: Foo**

**Scenario: Bar**

baz`,
		},
		{`
Feature: Foo
  Background: Bar
    When I do something`, `
**Feature: Foo**

**Background: Bar**

**When** I do something`,
		},
		{`
Feature: Foo
  Background: Bar
  Given Baz:
    | foo |
    | bar |`, `
**Feature: Foo**

**Background: Bar**

**Given** Baz:

| foo |
| bar |`,
		},
		{`
Feature: Foo
  Scenario Outline: Bar
    When <someone> does <something>
    Examples:
      | someone | something |
      | I       | cooking   |
      | You     | coding    |`, `
**Feature: Foo**

**Scenario Outline: Bar**

**When** <someone> does <something>

**Examples**:

| someone | something |
|---------|-----------|
| I       | cooking   |
| You     | coding    |`},
		{`
Feature: Foo
  Scenario Outline: Bar
    When <someone> does <something>
    Examples: Baz
      | someone | something |
      | I       | cooking   |
      | You     | coding    |`, `
**Feature: Foo**

**Scenario Outline: Bar**

**When** <someone> does <something>

**Examples**:

*Baz*

| someone | something |
|---------|-----------|
| I       | cooking   |
| You     | coding    |`},
		{`
Feature: Foo
  Scenario Outline: Bar
    When <someone> does <something>
    Examples: Baz
      foo bar baz

      | someone | something |
      | I       | cooking   |
      | You     | coding    |`, `
**Feature: Foo**

**Scenario Outline: Bar**

**When** <someone> does <something>

**Examples**:

*Baz*

foo bar baz

| someone | something |
|---------|-----------|
| I       | cooking   |
| You     | coding    |`},
		{`
Feature: Foo
  Scenario Outline: Bar
    When <someone> does <something>
    Examples: Baz
      | someone |
      | I       |
      | You     |
    Examples: Blah
      | something |
      | cooking   |
      | coding    |`, `
**Feature: Foo**

**Scenario Outline: Bar**

**When** <someone> does <something>

**Examples**:

*Baz*

| someone |
|---------|
| I       |
| You     |

*Blah*

| something |
|-----------|
| cooking   |
| coding    |`},
	} {
		d, err := gherkin.ParseGherkinDocument(strings.NewReader(ss[0]), func() string { return "" })

		assert.Nil(t, err)
		assert.Equal(t, strings.TrimSpace(ss[1])+"\n", NewRenderer().Render(d))
	}
}
