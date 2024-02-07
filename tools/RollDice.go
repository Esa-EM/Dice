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

func RollDice() int { //works
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

func Roll(dice []int) int { //works
	var number int

	rand.Seed(time.Now().UnixNano())

	// Generates a random integer between 0 and count of sides on dice
	steps := rand.Intn(len(dice))

	number = dice[steps]

	return number

}

func ReadDiceFromFile(index int) (string, error) { //works
	file, err := os.Open(Path("dicefile.txt"))
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

func LastLineIndex() (int, error) { //works
	// Open the dice.txt file for reading
	file, err := os.Open(Path("dicefile.txt"))
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

func DiceArt(value int) string { //not used
	switch value {
	case 1:
		return `
  ________
 |        |
 |   ●    |
 |        |
 |________|`
	case 2:
		return `
  ________
 | ●      |
 |        |
 |      ● |
 |________|`
	case 3:
		return `
  ________
 | ●      |
 |   ●    |
 |      ● |
 |________|`
	case 4:
		return `
  ________
 | ●    ● |
 |        |
 | ●    ● |
 |________|`
	case 5:
		return `
  ________
 | ●    ● |
 |   ●    |
 | ●    ● |
 |________|`
	case 6:
		return `
  ________
 | ●    ● |
 | ●    ● |
 | ●    ● |
 |________|`
	default:
		return "Invalid value, must be between 1 and 6"
	}
}
