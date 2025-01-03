package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if err := Run(os.Args[1:], &os.File{}); err != nil {
		if _, err := fmt.Fprintln(os.Stderr, err); err != nil {
			panic(err)
		}

		os.Exit(1)
	}
}

func Run(ss []string, writer io.Writer) error {
	args, err := GetArguments(ss)

	if err != nil {
		return err
	}

	return ConvertFile(args.File, writer)
}
