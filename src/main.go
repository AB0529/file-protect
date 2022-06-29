package main

import (
	"fmt"
)

func main() {
	var input int
	operations := []string{
		"1. Encode a file",
		"2. Decode a file",
	}
	// Get the operation to perform
	fmt.Println("What operation do you want to perform?")
	for _, opStr := range operations {
		fmt.Println(opStr)
	}
	
	fmt.Scanln(&input)

	switch input {
	case 1:
		// Get the file path
		var filepath string
		fmt.Println("Enter the filepath to the file to encode.")
		fmt.Scanln(&filepath)
		// Encode the file
		outputPath, err := EncodeOperation(filepath)
		
		if err != nil {
			panic(err)
		}
		
		fmt.Printf("Success, file now encoded as: %s", outputPath)
	case 2:
		// Get the file path
		var filepath string
		fmt.Println("Enter the filepath to the file to decode.")
		fmt.Scanln(&filepath)
		// Encode the file
		outputPath, err := DecodeOperation(filepath)
		
		if err != nil {
			panic(err)
		}
		
		fmt.Printf("Success, file now decoded as: %s", outputPath)
	default:
		fmt.Println("Unknown operation")
	}
}