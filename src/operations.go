package main

import (
	"path/filepath"
)

// EncodeOperation will encode the filepath provided
func EncodeOperation(path string) (string, error) {
	pw := GetPassword()
	// Make sure password is not empty
	if pw == "" {
		panic("password cannot be empty")
	}
	
	// Convert to abs path
	path, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	output, err := EncryptFile(path, []byte(pw))
	
	return output, err

}

// DecodeOperation will decode the filepath provided
func DecodeOperation(path string) (string, error) {
	pw := GetPassword()
	// Make sure password is not empty
	if pw == "" {
		panic("password cannot be empty")
	}

	// Convert to abs path
	path, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	output, err := UnencryptFile(path, []byte(pw))
		
	return output, err
}