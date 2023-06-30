package main

import (
	"encoding/base64"
	"fmt"
	"github.com/atotto/clipboard"
	"os"
)

func main() {
	// Check if there are command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Please provide a base64-encoded string argument.")
		return
	}

	// Get the base64 string argument
	base64Str := os.Args[1]

	// Decode the base64 string
	decoded, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		fmt.Println("Decoding error:", err)
		return
	}

	// Save the decoded value to the clipboard
	err = clipboard.WriteAll(string(decoded))
	if err != nil {
		fmt.Println("Clipboard error:", err)
		return
	}

	fmt.Println(string(decoded),"copied to clipboard.")
}
