package history

import (
	"dice/tools"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

// testing if I can make functions from packages print and make their menus and access tools functions

func History() { //works
	var invalid bool
	var deleted bool
	InitializeHistoryFile()

	//menu printing
	for {
		tools.ClearScreen()
		if invalid {
			fmt.Println("Invalid choice. Please use valid number.")
			invalid = false
		} else if deleted {
			deleted = false
			fmt.Println("---History deleted---")
			fmt.Println("Select with number:")
		}
		fmt.Println(`

	1. Show history
	2. Delete history
	3. Back
	4. Exit.

	`)
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error reading key:", err)
			break
		}
		switch {
		case char == '1':
			tools.ClearScreen()
			ShowHistory()
			fmt.Println(`Press any key to go back`)
			_, _, err := keyboard.GetKey()
			if err != nil {
				fmt.Println("Error reading key:", err)
				break
			}
			continue

		case char == '2':
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
				2 = I changed my mind...
				`)
				char, _, err := keyboard.GetKey()
				if err != nil {
					fmt.Println("Error reading key:", err)
					break
				}

				if char == '1' {
					ClearHistory()
					deleted = true
					break
				}
				if char == '2' {
					break
				} else {
					invalid = true
					continue
				}

			}
		case char == '3':
			return
		case char == '4' || key == keyboard.KeyEsc:
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

func AddToHistory(lastres int) { //works

	InitializeHistoryFile()
	//lets make int into string
	latestResult := strconv.Itoa(lastres)

	file, err := os.OpenFile(tools.Path("history.txt"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error adding history:", err)
		return
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, latestResult)
	if err != nil {
		fmt.Println("Error adding history:", err)
	}

}

func ShowHistory() { //works
	history, err := os.ReadFile(tools.Path("history.txt"))
	if err != nil {
		fmt.Println("Error reading history:", err)
		return
	}
	fmt.Println("\nHistory of dice rollings:")
	lines := strings.Split(string(history), "\n")
	for i, line := range lines {
		if i != 0 && line != "" {
			fmt.Printf("%s\n", line)
		}
	}
}

func ClearHistory() { //works
	err := os.Remove(tools.Path("history.txt"))
	if err != nil {
		fmt.Println("Error deleting history:", err)
	}
	InitializeHistoryFile()
}

func InitializeHistoryFile() error { //works
	// Check if the history file exists
	if _, err := os.Stat(tools.Path("history.txt")); os.IsNotExist(err) {
		// If it doesn't exist, create the file
		file, err := os.Create(tools.Path("history.txt"))
		if err != nil {
			return fmt.Errorf("error creating history file: %v", err)
		}
		defer file.Close()
		_, err = file.WriteString(`You should use app to read these
`)
		if err != nil {
			return fmt.Errorf("error writing to history file: %v", err)
		}
	}
	return nil
}
