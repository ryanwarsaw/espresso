package main

import (
	"log"
	"os"
)

func CreateDebugFile() *os.File {
	file, err := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Error creating debug file: %v", err)
	}
	return file
}