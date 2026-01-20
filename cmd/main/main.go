package main

import (
	"fmt"
	"log"
	"github.com/ABAlosaimi/DiFi/internal/combiner"
)

func main() {
	
	filePath := "/Users/abdulrahman/Documents/DiFi/cmd/main/hi.txt"
	filteringStartLineNum := 0 
	projectDir := "/Users/abdulrahman/Documents" 

	file, fileName, err := combiner.OpenFeatureFile(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	fmt.Printf("Opened file: %s\n", fileName)

	filteredFile, err := combiner.FilterFeatureFile(file, filteringStartLineNum)
	if err != nil {
		log.Fatalf("Failed to filter file: %v", err)
	}
	fmt.Printf("Filtered file, got %d lines\n", len(filteredFile))

	success, err := combiner.WriteFeatureFile(filteredFile, projectDir, fileName)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	if success {
		fmt.Printf("Successfully wrote filtered file to %s\n", projectDir)
	}
}