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

// Function to run inital scan of the files
func initialScan(root string) map[string]FileState {
	// Makes a empty map using filestate
	state := make(map[string]FileState)

	//  Goes over each file in the root and gets the path, info, error
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// Checks for error and info about the directory
		if err != nil || info.IsDir() {
			// returns nil
			return nil
		}
		// gets the hash from teh hash file function
		hash, err := hashFile(path)
		// Checks if the error is nil
		if err == nil {
			// Stores the file stat in the map
			state[path] = FileState{Hash: hash}
		}
		// returns nil
		return nil
	})

	// Returns the state
	return state
}

