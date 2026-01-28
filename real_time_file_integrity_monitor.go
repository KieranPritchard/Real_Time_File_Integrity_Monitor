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