package tools

import (
	"fmt"
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
