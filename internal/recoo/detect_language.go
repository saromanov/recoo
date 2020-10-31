package recoo

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-enry/go-enry/v2"
)

type Language string

var (
	GO      Language = iota
	Unknown Language
)

// detectLanguage provides detecting of programming language
func detectLanguage(path string) (Language, error) {
	d, err := readFile(path)
	if err != nil {
		return Unknown, fmt.Errorf("unable to read file: %v", err)
	}
	lang, _ := enry.GetLanguageByContent("", d)
	switch lang {
	case "Go":
		return GO, nil
	}
	return Unknown, fmt.Errorf("unknown language")
}

func readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open file")
	}
	d, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("unable to read data from file: %v", err)
	}
	return d, nil
}
