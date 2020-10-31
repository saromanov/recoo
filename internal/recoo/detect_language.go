package recoo

import "github.com/go-enry/go-enry/v2"

// detectLanguage provides detecting of programming language
func detectLanguage(path string) error {
	d, err := readFile(path)
	if err != nil {
		return fmt.Errorf("unable to read file: %v", err)
	}
	lang, _ := enry.GetLanguageByContent("", d)
	fmt.Println(lang)
	return nil
}

func readFile(path string)([]byte, error) {
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