package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// WriteJsonInFile write given struct as Json in given file (create it if needed)
func WriteJsonInFile(s interface{}, path string) {
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("Can't parse JSON for file `%s`", path)
		os.Exit(1)
	}

	err = ioutil.WriteFile(path, b, os.ModePerm)
	if err != nil {
		fmt.Printf("Can't write JSON file `%s`", path)
		os.Exit(1)
	}
}

// ReadJson is used to read Json from given path
func ReadJson(path string) []byte {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Can't read JSON file `%s`", path)
		os.Exit(1)
	}
	return b
}
