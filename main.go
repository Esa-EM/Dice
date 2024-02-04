package main

import (
	"bufio"
	"dice/tools"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const dicefile = "dice.txt"

func main() {

	var invalid bool
	var result int
	var uf bool

	if len(os.Args) != 1 {
		tools.ClearScreen()
		tools.Helper()
	} else if len(os.Args) == 1 {
		InitializeDiceFile()
	out:
		//menu printing
		for {
			tools.ClearScreen()
			if invalid {
				fmt.Println("Invalid choice. Please use valid number.")
				invalid = false
			} else if uf {
				fmt.Println("You selected upcoming feature. it isnt there yet, wait for next version")
				uf = false
			} else {
				fmt.Println("Select with number:")
			}
			fmt.Println(`

	1. Roll dice
	2. Change dice
	3. Betting game (upcoming)
	4. History & statistics (upcoming)
	5. Exit.

	`)
			var choice int
			fmt.Scan(&choice)
			switch choice {
			case 1:
				result = tools.RollDice()
				break out

			case 2:

				ChangeDice()
			case 3:
				uf = true
				continue
			case 4:
				uf = true
				continue
			case 5:
				tools.ClearScreen()
				fmt.Println(`No more rolls. Goodbye!`)
				time.Sleep(2 * time.Second)
				tools.Clear()
				os.Exit(0)
			default:
				invalid = true
				continue

			}
		}

	}
	tools.ClearScreen() //lets print roll result out!
	for i := 0; i < 1; i++ {
		fmt.Println("Rolling")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(`Rolling
	`)
	time.Sleep(1 * time.Second)
	for {
		if invalid {
			fmt.Println("Invalid choice. Please use valid number.")
			fmt.Println(`Rolled:
			`)
			invalid = false
		}
		fmt.Println(result)
		time.Sleep(1 * time.Second)
		fmt.Println(`
	
	1. Back
	2. Exit
	`)
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			main()
		case 2:
			tools.ClearScreen()
			fmt.Println(`No more rolls. Goodbye!`)
			time.Sleep(2 * time.Second)
			tools.Clear()
			os.Exit(0)
		default:
			invalid = true
			continue
		}
	}

}

func ChangeDice() { //works
	var invalid bool
	for {
		tools.ClearScreen()
		if invalid {
			fmt.Println("Invalid choice. Please use valid number.")
			invalid = false
		} else {
			fmt.Println("Select with number:")
		}
		fmt.Println(`
	
	1. Select dice
	2. Make new dice
	3. Reset dices (restores default dice set)
	4. Back

	
	
	`)
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			selDices()
		case 2:
			AddDice()
		case 3:
			for {
				tools.ClearScreen()
				if invalid {
					fmt.Println("Invalid choice. Please use valid number.")
					invalid = false
				} else {
					fmt.Println(`Are you sure? This cannot be undone`)
				}
				fmt.Println(`
				1 = YES!
				2 = NONO ABORT MISSION!
				`)
				fmt.Scan(&choice)

				switch choice {
				case 1:
					resetDices()
				case 2:
					main()
				default:
					invalid = true
					continue
				}

			}

		case 4:
			main()
		default:
			invalid = true
			continue

		}
	}
}

func selDices() []int { // works
	var invalid bool
	for {
		tools.ClearScreen()
		showDices()
		fmt.Println("")
		if invalid {
			fmt.Println("Invalid choice. Please use valid number.")
			invalid = false
		} else {
			fmt.Println("Select dice by giving its number:")
		}
		fmt.Println(`input 1 to return

	`)

		var choice string

		fmt.Scan(&choice)
		if choice == "1" {
			ChangeDice()
		}
		lastI, err := tools.LastLineIndex()
		if err != nil {
			fmt.Println("invalid input: not a number")
		}
		choiceint, err := strconv.Atoi(choice)
		if err != nil {
			fmt.Println("invalid input: not a number")
		}

		if choiceint > lastI {
			choiceint = lastI
			choice = strconv.Itoa(choiceint)
		}

		//Edit first line to be latest dice index
		EditFirstLine(choice)
		ChangeDice()
	}
}

func showDices() { //works
	Dices, err := os.ReadFile(dicefile)
	if err != nil {
		fmt.Println("Error reading Dices:", err)
		return
	}
	fmt.Println("\nDices:")
	lines := strings.Split(string(Dices), "\n")
	for i, line := range lines {
		if i != 0 && line != "" {
			fmt.Printf("%0d - %s\n", i+1, line)
		}
	}
}

func AddDice() { //works
	tools.ClearScreen()
	fmt.Print(`
	Enter the dice values separated with space and hit enter. Example:
	1 2 3 4 5 6 
	only numbers will be accepted
	Input your values:`)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	toValidate := scanner.Text()
	newDice := tools.Validate(toValidate)
	if len(newDice) != len(toValidate) {
		tools.ClearScreen()
		fmt.Println("invalid input")
		AddDice()
	}

	file, err := os.OpenFile(dicefile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error adding dice:", err)
		return
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, newDice)
	if err != nil {
		fmt.Println("Error adding dice:", err)
	}
	// now we set new dice to be current one
	lastIndex, err := tools.LastLineIndex()
	if err != nil {
		fmt.Println("Error adding dice:", err)
	}
	// Convert the index to a string
	lastIndexStr := strconv.Itoa(lastIndex)
	//Edit first line to be latest dice index
	EditFirstLine(lastIndexStr)
}

func InitializeDiceFile() error { //works
	// Check if the dice.txt file exists
	if _, err := os.Stat(dicefile); os.IsNotExist(err) {
		// If it doesn't exist, create the file
		file, err := os.Create(dicefile)
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
func resetDices() { //works

	err := os.Remove(dicefile)
	if err != nil {
		fmt.Println("Error deleting empty file:", err)
	}
	main()
}

func EditFirstLine(newFirstLine string) error {
	// Open the dice.txt file for reading and writing
	file, err := os.OpenFile(dicefile, os.O_RDWR, 0644)
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
