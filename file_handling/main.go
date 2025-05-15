package main

import (
	"fmt"
	"os"
)

func main() {

	f, err1 := os.Open("c:/Users/Administrator/Desktop/go_lang/file_handling/example.txt")

	if err1 != nil {
		panic(fmt.Sprintf("Error opening file: %v", err1))
	}

	defer f.Close()

	fileInfo, err2 := f.Stat()

	if err2 != nil {
		panic(fmt.Sprintf("Error getting file info: %v", err2))
	}

	buffer := make([]byte, fileInfo.Size())

	_, err3 := f.Read(buffer)

	if err3 != nil {
		panic(fmt.Sprintf("Error reading file content: %v", err3))
	}
	fmt.Println("File Content:", string(buffer))
	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("File Size:", fileInfo.Size())
	fmt.Println("File Mode:", fileInfo.Mode())
	fmt.Println("Last Modified:", fileInfo.ModTime())
	fmt.Println("Is Directory:", fileInfo.IsDir())
	fmt.Println("Sys:", fileInfo.Sys())

	fmt.Println("--------------------------------------------------------------")
	data, err4 := os.ReadFile("c:/Users/Administrator/Desktop/go_lang/file_handling/example.txt")
	if err4 != nil {
		panic(fmt.Sprintf("Error reading file: %v", err4))
	}
	fmt.Println("File Content using ReadFile:", string(data))

}
