package main

import (
	"github.com/docopt/docopt-go"
)

const usage = `Gherkin to Markdown for Jira AC converter

Usage:
	gherkin2jira-ac <file>

Options:
	-h, --help  Show this help.`

type Arguments struct {
	File string `docopt:"<file>"`
}

func GetArguments(ss []string) (Arguments, error) {
	args := Arguments{}
	err := parseArguments(usage, ss, &args)
	return args, err
}

func parseArguments(u string, ss []string, args interface{}) error {
	opts, err := docopt.ParseArgs(u, ss, "")

	if err != nil {
		return err
	}

	return opts.Bind(args)
}
