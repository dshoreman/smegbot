package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// WriteJSON converts a struct to JSON before calling WriteFile
func WriteJSON(file string, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return WriteFile(file, b)
}

// WriteFile takes a JSON byte represenation and writes it to a file
func WriteFile(file string, content []byte) error {
	err := os.MkdirAll(filepath.Dir(file), 0700)
	if err != nil {
		fmt.Println("\nError: Couldn't create directory "+file, err)
		return err
	}
	return ioutil.WriteFile(file, content, 0644)
}
