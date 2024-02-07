package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Validate(input string) string {
	// Split the input string by spaces
	values := strings.Fields(input)

	// Create a slice to store valid dice values
	var validValues []string

	// Validate each value
	for _, value := range values {
		// Attempt to convert the value to an integer
		_, err := strconv.Atoi(value)
		if err != nil {
			// If conversion fails, the value is not a number
			fmt.Printf("Invalid value: %s\n", value)
			continue
		}
		// If conversion succeeds, add the value to the list of valid values
		validValues = append(validValues, value)
	}

	// Join the valid values into a single string separated by spaces
	validatedInput := strings.Join(validValues, " ")

	return validatedInput
}
func Path(fileName string) string {
	dirName := "files"
	var path string
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	if strings.HasPrefix(exPath, "/var/") {
		err = os.MkdirAll(dirName, 0755) // Create directories if they don't exist
		if err != nil {
			panic(err)
		}
		path = filepath.Join(dirName, fileName)
	} else {
		path = filepath.Join(exPath, dirName, fileName)
		err = os.MkdirAll(filepath.Join(exPath, dirName), 0755) // Create directories if they don't exist
		if err != nil {
			panic(err)
		}
	}
	return path
}

func InitializeDiceFile() error { //works
	// Check if the dice.txt file exists
	if _, err := os.Stat(Path("dicefile.txt")); os.IsNotExist(err) {
		// If it doesn't exist, create the file
		file, err := os.Create(Path("dicefile.txt"))
		if err != nil {
			return fmt.Errorf("error creating dice file: %v", err)
		}
		defer file.Close()

		// Fill the file with initial dice values
		initialContent := `2
1 2 3 4 5 6
1 2 3 4 5 6 7 8 9 10
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20
`
		_, err = file.WriteString(initialContent)
		if err != nil {
			return fmt.Errorf("error writing to dice file: %v", err)
		}
	}
	return nil
}
