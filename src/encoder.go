package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// EncrpytFile encrypts a file with AES
func EncryptFile(path string, key []byte) (string, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
    key, salt, err := NewKey(key, nil)
    if err != nil {
        return "", err
    }
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    nonce := make([]byte, gcm.NonceSize())
    if _, err = rand.Read(nonce); err != nil {
        return "", err
    }

    ciphertext := gcm.Seal(nonce, nonce, f, nil)
    ciphertext = append(ciphertext, salt...)

	name := filepath.Base(path)

    // Create a new file for saving the encrypted data.
    f2, err := os.Create(name)
    if err != nil {
        panic(err.Error())
    }
    _, err = io.Copy(f2, bytes.NewReader(ciphertext))
    if err != nil {
		return "", err
    }
	
	return name, nil
}

// UnencryptFile decodes a AES encrypted file
func UnencryptFile(path string, pw []byte) (string, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
    salt, f := f[len(f) - 32:], f[:len(f) - 32]
    key, _, err := NewKey(pw, salt)
    if err != nil {
        return "", err
    }
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    nonce, ciphertext := f[:gcm.NonceSize()], f[gcm.NonceSize():]

    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return "", err
    }

	name := filepath.Base(path)

    // Create a new file for saving the decrypted data.
    f2, err := os.Create(name)
    if err != nil {
        panic(err.Error())
    }
    _, err = io.Copy(f2, bytes.NewReader(plaintext))
    if err != nil {
		return "", err
    }
	
	return name, nil
}