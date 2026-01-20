package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/ABAlosaimi/DiFi/internal/combiner"
)

func main() {

	if len(os.Args) < 4 {
		log.Fatalf("Usage: %s <filePath> <filteringStartLineNum> <projectDir>", os.Args[0])
	}

	filePath := os.Args[1]
	projectDir := os.Args[3]
	filteringStartLineNum, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid filteringStartLineNum: %v", err)
	} 

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