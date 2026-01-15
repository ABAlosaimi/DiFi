package main

import (
	"bufio"
	"fmt"
	"os"
)


func main()  {
	// the below code was for getting hands dirty with the file I/O in go 

	// filepath := "/Users/abdulrahman/Documents/DiFi/cmd/hi.java"

	// data, err := os.Open(filepath)
	// if err != nil {
	// 	fmt.Printf("Error reading file: %v\n", err)
	// 	return
	// }
	// defer data.Close()

	// var filteredFile []string
	// scanner := bufio.NewScanner(data)

	// for scanner.Scan() {
	// 	filteredFile = append(filteredFile, scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
    //     fmt.Println("Error reading file:", err)
    // }

	// file := printNLines(filteredFile, 5)

	// for i := 0; i < len(file); i++ {
	// 	fmt.Println(file[i])
	// }

	// rdFile, err := os.Create("hi.txt") // it tested to create a file 

	// for i := 0; i < len(file); i++{
	// 	rdFile.Write([]byte(file[i] + "\n"))
	// }

	// if err != nil {
	// 	fmt.Println("erro withe the created file: " + err.Error())
	// }

	// final, err := os.ReadFile(rdFile.Name())

	// fmt.Println(string(final))
}