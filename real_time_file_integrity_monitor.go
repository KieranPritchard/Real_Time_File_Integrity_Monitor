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

// Worker function. takes id, jobs, and state as a function
func worker(id int, jobs <-chan Job, state map[string]FileState, mu *sync.Mutex) {
	// Loops over the range of jobs
	for job := range jobs {
		switch {
		// Handle file creation or modification events
		case job.Op&(fsnotify.Create|fsnotify.Write) != 0:
			// Calculate the new hash to see if the content actually changed
			hash, err := hashFile(job.Path)
			// Catchs the error
			if err != nil {
				// Goes to next iteration
				continue
			}

			// Lock the map before reading or writing to prevent race conditions
			mu.Lock()
			// Gets the old hash and whether the file exists
			old, exists := state[job.Path]
			// Checks if the file exists and the if the hashes dont match
			if exists && old.Hash != hash {
				// Outputs the file at the path is modified
				fmt.Println("[MODIFIED]", job.Path)
			// Checks if the file doesnt not exist
			} else if !exists {
				// Outputs the file at path has been created
				fmt.Println("[CREATED]", job.Path)
			}
			// Sets the new state from the file state data struct
			state[job.Path] = FileState{Hash: hash}
			// Unlocks the map
			mu.Unlock()
		
		// Handle file deletion events
		case job.Op&fsnotify.Remove != 0:
			// Locks the map
			mu.Lock()
			// Remove the file from our internal tracking map
			delete(state, job.Path)
			// Unlocks the map
			mu.Unlock()
			// Prints the file is is deleted
			fmt.Println("[DELETED]", job.Path)
		}
	}
}

// Main function
func main() {
	// Outputs the monitor is started
	fmt.Println("Real-time File Integrity Monitor starting...")

	// Gets the state from the inital scan function
	state := initialScan(watchDir)
	// Prints a baseline is established
	fmt.Println("Baseline established")

	// Creates a new watcher with fsnotfiy and new watcher
	watcher, err := fsnotify.NewWatcher()
	// Checks for error
	if err != nil {
		// Panics on error
		panic(err)
	}
	// Closes the watcher when done
	defer watcher.Close()

	// Checks for error from watcher checking watch directory
	if err := watcher.Add(watchDir); err != nil {
		// Panics on error
		panic(err)
	}

	// Creates a map of jobs
	jobs := make(chan Job, 100)
	// Adds sync mutex to the path
	var mu sync.Mutex

	// Start workers
	for i := 0; i < workers; i++ {
		// Creates go worker function
		go worker(i, jobs, state, &mu)
	}

	// Event dispatcher goroutine
	go func() {
		// Loops over each event in watchers events
		for event := range watcher.Events {
			// Adds the job to jobs
			jobs <- Job{
				Path: event.Name,
				Op:   event.Op,
			}
		}
	}()

	// Error handler goroutine
	go func() {
		// Loops over the errors
		for err := range watcher.Errors {
			// Prints the watcher error
			fmt.Println("Watcher error:", err)
		}
	}()

	// Keeps process alive
	for {
		// Sleeps for the second
		time.Sleep(time.Second)
	}
}