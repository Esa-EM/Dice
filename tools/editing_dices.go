package tools

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AddDice() { //works, but doesnt work for windows?
	ClearScreen()
	fmt.Print(`
    Enter the dice values separated with space and hit enter. Example:
    1 2 3 4 5 6 
    only numbers will be accepted
    Input your values:`)
	reader := bufio.NewReader(os.Stdin)
	toValidate, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	// Replace both "\r" and "\n" with empty strings. Lets hope this fixes windows problem.
	toValidate = strings.ReplaceAll(toValidate, "\r", "")
	toValidate = strings.ReplaceAll(toValidate, "\n", "")

	newDice := Validate(toValidate)
	if len(newDice) != len(toValidate) {
		ClearScreen()
		fmt.Println("invalid input")
		AddDice()
	}

	file, err := os.OpenFile(Path("dicefile.txt"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error adding dice:", err)
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, newDice)
	if err != nil {
		fmt.Println("Error adding dice:", err)
	}
	// now we set new dice to be current one
	lastIndex, err := LastLineIndex()
	if err != nil {
		fmt.Println("Error adding dice:", err)
	}
	// Convert the index to a string
	lastIndexStr := strconv.Itoa(lastIndex)
	//Edit first line to be latest dice index
	EditFirstLine(lastIndexStr)
}

func ResetDices() { //works

	err := os.Remove(Path("dicefile.txt"))
	if err != nil {
		fmt.Println("Error deleting dices:", err)
	}
}

func EditFirstLine(newFirstLine string) error { //works
	// Open the dice.txt file for reading and writing
	file, err := os.OpenFile(Path("dicefile.txt"), os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("error opening dice file: %v", err)
	}
	defer file.Close()

	// Seek back to the beginning of the file
	_, err = file.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("error seeking to the beginning of the dice file: %v", err)
	}

	// Write the modified first line back to the file
	_, err = fmt.Fprintf(file, "%s\n", newFirstLine)
	if err != nil {
		return fmt.Errorf("error writing to dice file: %v", err)
	}

	return nil
}
