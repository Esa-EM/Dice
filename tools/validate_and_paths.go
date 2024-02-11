package tools

import (
	"bufio"
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

func RemoveEmptyLines() error {
	// Open the dicefile.txt file for reading and writing
	file, err := os.OpenFile(Path("dicefile.txt"), os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create a temporary file to store the non-empty lines
	tempFile, err := os.CreateTemp("", "temp")
	if err != nil {
		return fmt.Errorf("error creating temporary file: %v", err)
	}
	defer tempFile.Close()

	// Create a scanner to read from the original file
	scanner := bufio.NewScanner(file)

	// Iterate over each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line is empty
		if strings.TrimSpace(line) != "" {
			// Write non-empty lines to the temporary file
			_, err := fmt.Fprintln(tempFile, line)
			if err != nil {
				return fmt.Errorf("error writing to temporary file: %v", err)
			}
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error scanning file: %v", err)
	}

	// Close the original file
	if err := file.Close(); err != nil {
		return fmt.Errorf("error closing original file: %v", err)
	}

	// Remove the original file
	if err := os.Remove(Path("dicefile.txt")); err != nil {
		return fmt.Errorf("error removing original file: %v", err)
	}

	// Rename the temporary file to the original dicefile.txt
	if err := os.Rename(tempFile.Name(), Path("dicefile.txt")); err != nil {
		return fmt.Errorf("error renaming temporary file: %v", err)
	}

	return nil
}
