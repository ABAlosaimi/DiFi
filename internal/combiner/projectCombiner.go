package combiner

import (
	"bufio"
	"fmt"
	"os"
)

func OpenFeatureFile(filePath string) os.File {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("opening file error: " + err.Error())
	}
	defer file.Close()

	return *file
}

func removeNlines(featureFile []string, num int) *os.File {
	 newFeatureFile, err := os.Create("x.java")
	if err != nil {
		fmt.Println("file creation error: " + err.Error())
	}

	for i := num; i < len(featureFile); i++ {
		_, err := newFeatureFile.Write([]byte(featureFile[i]))
		if err != nil {
			fmt.Println("somthing went wrong with the filteration process: " + err.Error())
			break
		}
	}

	return newFeatureFile
}

func FilterFeatureFile(featureFile os.File, filteringStartNum int) *os.File {
	var featureFileArray []string
	scanner := bufio.NewScanner(&featureFile)

	for scanner.Scan() {
		featureFileArray = append(featureFileArray, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }

	filteredFile := removeNlines(featureFileArray, filteringStartNum)

	return filteredFile
}

func AddFeatureFile(featureFile os.File, projectDir string) bool {
	
	// adding the filtered file to the new prject diroctory 
	return true
}