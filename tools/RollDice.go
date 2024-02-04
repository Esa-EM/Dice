package tools

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func RollDice() int {
	var index string
	var err error
	var result int
	var i int
	index, err = ReadDiceFromFile(1)
	if err != nil {
		fmt.Println("Error:", err)
	}
	i, err = strconv.Atoi(index)
	if err != nil {
		fmt.Println("Error:", err)
	}
	dice, err := ReadDiceFromFile(i)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Split the string by spaces to get individual substrings
	strValues := strings.Split(dice, " ")

	// Create a slice to store the integers
	intValues := make([]int, len(strValues))

	// Convert each substring to an integer
	for i, s := range strValues {
		// Convert string to int
		val, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Error converting to integer:", err)
		}
		intValues[i] = val
	}

	result = Roll(intValues)

	return result
}

func Roll(dice []int) int {
	var number int

	rand.Seed(time.Now().UnixNano())

	// Generates a random integer between 0 and count of sides on dice
	steps := rand.Intn(len(dice))

	number = dice[steps]

	return number

}

func ReadDiceFromFile(index int) (string, error) {
	file, err := os.Open(diceFile)
	if err != nil {
		return "", fmt.Errorf("error opening dice file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		if lineNumber == index {
			return scanner.Text(), nil
		}
		lineNumber++
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error scanning dice file: %v", err)
	}

	return "", fmt.Errorf("line %d not found in dice file", index)
}

func LastLineIndex() (int, error) {
	// Open the dice.txt file for reading
	file, err := os.Open(diceFile)
	if err != nil {
		return 0, fmt.Errorf("error opening dice file: %v", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	var lastIndex int

	// Read each line until the end of the file
	for scanner.Scan() {
		lastIndex++
	}

	// Check for any errors that may have occurred during scanning
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error scanning dice file: %v", err)
	}

	// Return the index of the last line
	return lastIndex, nil
}
