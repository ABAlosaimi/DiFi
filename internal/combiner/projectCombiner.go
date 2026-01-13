package combiner

import (
	"fmt"
	"os"
)

func addPreviousFeature(featureFileName string, featureDirName string, projectDir string, projectSubDir string) bool {
    file, err := os.Open(featureDirName + "/" + featureFileName)
    if err != nil {
        return false
    }
    defer file.Close()

    return true
}