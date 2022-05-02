package build

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-enry/go-enry/v2"
)

type Language int

const (
	GO     Language = iota
	Python Language = 1
	Unknown
)

// detectLanguage provides detecting of programming language
func detectLanguage(path string) (Language, error) {
	d, err := readFile(path)
	if err != nil {
		return Unknown, fmt.Errorf("unable to read file: %v", err)
	}
	langs := enry.GetLanguagesByExtension(filepath.Base(path), d, nil)
	if len(langs) > 1 {
		return Unknown, fmt.Errorf("unable to detect one language: %v", langs)
	}
	switch langs[0] {
	case "Go":
		return GO, nil
	case "Python":
		return Python, nil
	}
	return Unknown, fmt.Errorf("unknown language")
}

func readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %v", err)
	}
	d, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("unable to read data from file: %v", err)
	}
	return d, nil
}
