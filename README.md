# gherkin2jira-ac

A command to convert Gherkin files into Markdown for Jira AC.

Loosely inspired by [gherkin2markdown](https://github.com/raviqqe/gherkin2markdown).

## Installation

```
go install github.com/alle/gherkin2jira-ac@latest
```

## Usage

```
gherkin2jira-ac <file>
```

## Example

**Given** a file named `math.feature` with:

```gherkin
Feature: Python
  Scenario: Hello, world!
    Given a file named "main.py" with:
    """python
    print("Hello, world!")
    """
    When I successfully run `python3 main.py`
    Then the stdout should contain exactly "Hello, world!"

  Scenario Outline: Add numbers
    Given a file named "main.py" with:
    """python
    print(<x> + <y>)
    """
    When I successfully run `python3 main.py`
    Then the stdout should contain exactly "<z>"

    Examples:
      | x | y | z |
      | 1 | 2 | 3 |
      | 4 | 5 | 9 |
```

**When** I successfully run `gherkin2jira-ac math.feature`

**Then** the stdout should contain exactly:

````markdown
**Feature**: *Python*

**Scenario**: *Hello, world!*
  **Given** a file named "main.py" with:

  ```python
print("Hello, world!")
  ```
  **When** I successfully run `python3 main.py`
  **Then** the stdout should contain exactly "Hello, world!"

**Scenario Outline**; *Add numbers*
  **Given** a file named "main.py" with:

  ```python
print(<x> + <y>)
  ```
  **When** I successfully run `python3 main.py`
  **Then** the stdout should contain exactly "<z>"

**Examples**:

| x   | y   | z   |
| --- | --- | --- |
| 1   | 2   | 3   |
| 4   | 5   | 9   |
````

## License

[MIT](LICENSE)
