package main

import (
	"fmt"
	"example.com/monorepo/internal/files"
)

func main() {
	filesList := files.ListFiles()
	fmt.Printf("✅ Files service: %d files found\n", len(filesList))
	for _, f := range filesList {
		fmt.Println("-", f)
	}
}
