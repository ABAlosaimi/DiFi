package combiner

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func OpenFeatureFile(filePath string) (os.File, string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("opening file error: " + err.Error())
	}

	return *file, filepath.Base(file.Name()), nil
}

func removeNlines(featureFile []string, filteringStartLineNum int) ([]string, error) { 
	var filteredFile []string

	for i := filteringStartLineNum; i < len(featureFile); i++ {
		filteredFile = append(filteredFile, featureFile[i])
	}

	return filteredFile, nil
}

func FilterFeatureFile(featureFile os.File, filteringStartLineNum int) ([]string, error){
	var featureFileArray []string
	scanner := bufio.NewScanner(&featureFile)

	for scanner.Scan() {
		featureFileArray = append(featureFileArray, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file: ", err)
		return nil, err
    }

	filteredFile, err:= removeNlines(featureFileArray, filteringStartLineNum)
	if  err != nil {
		fmt.Println("feature file filtring error: " + err.Error())
	}

	return filteredFile, nil
}

func WriteFeatureFile(filteredFile []string, projectDir string, featureFileName string) (bool, error) {
	targetPath := filepath.Join(projectDir, featureFileName)

	newFeatureFile, err := os.Create(targetPath)
	if err != nil {
		fmt.Println("filtered file creation error: " + err.Error())
	}
	defer newFeatureFile.Close()

	for i := 0; i < len(filteredFile); i++ {
		if _, err := newFeatureFile.Write([]byte(filteredFile[i] + "\n")); err != nil {
			fmt.Println("writing file to the new/other project error: " + err.Error())
		}
	}

	return true, nil
}