package main

import (
	"bufio"
	"fmt"
	"os"
)


func printNLines(lines []string, num int) []string {
    var printNLines []string
    for i := num; i < len(lines); i++ {
        printNLines = append(printNLines, lines[i])
    }
    return printNLines
}


func main()  {
	filepath := "/Users/abdulrahman/Documents/DiFi/cmd/hi.java"

	data, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	defer data.Close()

	var filteredFile []string
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		filteredFile = append(filteredFile, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }

	file := printNLines(filteredFile, 5)

	for i := 0; i < len(file); i++ {
		fmt.Println(file[i])
	}

}