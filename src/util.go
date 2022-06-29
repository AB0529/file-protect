package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"syscall"

	"golang.org/x/crypto/scrypt"
	"golang.org/x/term"
)

// Config config which holds password info
// TODO: Use PGP keys or hash password
type Config struct {
	Password string
}

// Folder the structure for a locked folder config
type Folder struct {
	Path string
}

var (
	ErrEmptyPassword = errors.New("password cannot be empty")
	ErrDuplicateFolder = errors.New("folder already added")
)

// GetPassword gets the user password
func GetPassword() string {
	fmt.Println("Enter password for the file.")
	pw, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}
	return string(pw)
}

// NewKey generates a 32byte key from password
func NewKey(password, salt []byte) ([]byte, []byte, error) {
    if salt == nil {
        salt = make([]byte, 32)
        if _, err := rand.Read(salt); err != nil {
            return nil, nil, err
        }
    }

    key, err := scrypt.Key(password, salt, 1048576, 8, 1, 32)
    if err != nil {
        return nil, nil, err
    }

	return key, salt, nil
}