package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/cucumber/gherkin/go/v27"
)

const featureFileExtension = ".feature"
const markdownFileExtension = ".md"

func ConvertFile(sourceFile string, writer io.Writer) error {
	file, err := os.Open(sourceFile)

	if err != nil {
		return err
	}

	document, err := gherkin.ParseGherkinDocument(file, func() string { return sourceFile })

	if err != nil {
		return err
	}

	convertedDocument := NewRenderer().Render(document)

	destFile, err := openDestFile(sourceFile)

	if err != nil {
		return err
	}
	_, err = fmt.Fprint(destFile, convertedDocument)

	return err
}

func openDestFile(sourceFile string) (*os.File, error) {
	path := strings.TrimSuffix(sourceFile, featureFileExtension) + markdownFileExtension

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		return nil, err
	}

	err = file.Truncate(0)

	if err != nil {
		return nil, err
	}

	return file, nil
}
