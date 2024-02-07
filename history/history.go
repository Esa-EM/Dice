package history

import (
	"dice/tools"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// testing if I can make functions from packages print and make their menus and access tools functions

func History() { //works
	var invalid bool
	InitializeHistoryFile()

	//menu printing
	for {
		tools.ClearScreen()
		if invalid {
			fmt.Println("Invalid choice. Please use valid number.")
			invalid = false
		} else {
			fmt.Println("Select with number:")
		}
		fmt.Println(`

	1. Show history
	2. Delete history
	3. Back
	4. Exit.

	`)
		var choice int
		var choice2 string
		fmt.Scan(&choice)
		switch choice {
		case 1:
			tools.ClearScreen()
			ShowHistory()
			fmt.Println(`input anything and hit enter to go back`)
			fmt.Scan(&choice2)
			if choice2 != "gaomvapomvewoivamkl" {
				History()
			}

		case 2:
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
				fmt.Scan(&choice)

				switch choice {
				case 1:
					ClearHistory()
					History()
				case 2:
					History()
				default:
					invalid = true
					continue
				}

			}
		case 3:
			return
		case 4:
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
			return fmt.Errorf("error writing to dice file: %v", err)
		}
	}
	return nil
}
