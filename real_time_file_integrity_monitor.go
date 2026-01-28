package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
	"github.com/fsnotify/fsnotify"
)

// Structure to store file state
type FileState struct {
	// Stores the files hash
	Hash string
}

// Stores the job information
type Job struct {
	Path string // Stores the file path
	Op   fsnotify.Op // Stores the operation of the fsnotify
}

const (
	watchDir = "./testdata"
	workers  = 4
)

// Function to hash files
func hashFile(path string) (string, error) {
	// Opens the file from the path and gets the error
	file, err := os.Open(path)
	// Checks if the error is not nil
	if err != nil {
		// returns a empty strng and the error
		return "", err
	}
	// Closes the file when done
	defer file.Close()

	// Creates new hash
	hash := sha256.New()
	// Checks if the io copy is hash and a file and checks if error isnnt ill
	if _, err := io.Copy(hash, file); err != nil {
		// returns a empty strng and the error
		return "", err
	}
	// Returns the hex value of the file hash
	return hex.EncodeToString(hash.Sum(nil)), nil
}