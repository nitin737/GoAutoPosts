package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nitin737/GoAutoPosts/go-daily/internal/model"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run validate_data.go <path-to-libraries.json>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	// Read file
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Parse JSON
	var libraries []model.Library
	if err := json.Unmarshal(data, &libraries); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Validate each library
	errors := 0
	for i, lib := range libraries {
		if lib.Name == "" {
			fmt.Printf("Library %d: missing name\n", i)
			errors++
		}
		if lib.Description == "" {
			fmt.Printf("Library %d (%s): missing description\n", i, lib.Name)
			errors++
		}
		if lib.URL == "" {
			fmt.Printf("Library %d (%s): missing URL\n", i, lib.Name)
			errors++
		}
		if lib.Category == "" {
			fmt.Printf("Library %d (%s): missing category\n", i, lib.Name)
			errors++
		}
		if len(lib.Tags) == 0 {
			fmt.Printf("Library %d (%s): no tags\n", i, lib.Name)
			errors++
		}
	}

	if errors > 0 {
		fmt.Printf("\n❌ Validation failed with %d errors\n", errors)
		os.Exit(1)
	}

	fmt.Printf("✅ Validation successful! %d libraries validated.\n", len(libraries))
}
