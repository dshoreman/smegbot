package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// FileExists checks whether a path exists and is not a dir
func FileExists(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !f.IsDir()
}

// ReadJSON reads and unmarshals a JSON file
func ReadJSON(path string, v interface{}) error {
	return json.Unmarshal(ReadFile(path), &v)
}

// ReadString reads a file and casts the byte slice to a string
func ReadString(path string) string {
	return string(ReadFile(path))
}

// ReadFile attempts to read a file into a byte slice
func ReadFile(path string) []byte {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Couldn't read file"+path, err)
	}
	return b
}

// WriteJSON converts a struct to JSON before calling WriteFile
func WriteJSON(file string, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println("\nError: Failed encoding JSON", err)
		return err
	}
	return WriteFile(file, b)
}

// WriteFile takes a JSON byte represenation and writes it to a file
func WriteFile(file string, content []byte) error {
	if err := os.MkdirAll(filepath.Dir(file), 0700); err != nil {
		fmt.Println("\nError: Couldn't create directory "+filepath.Dir(file), err)
		return err
	}
	if err := ioutil.WriteFile(file, content, 0644); err != nil {
		fmt.Println("\nError: Couldn't write file "+file, err)
		return err
	}
	return nil
}
