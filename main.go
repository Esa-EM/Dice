package main

import (
	"dice/history"
	"dice/tools"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() { //works for everything that is implemented for now

	var invalid bool
	var result int
	var uf bool

	if len(os.Args) != 1 {
		tools.ClearScreen()
		tools.Helper()
	} else if len(os.Args) == 1 {
		tools.InitializeDiceFile()
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
				Result(result)

			case 2:

				ChangeDice()
			case 3:
				uf = true
				continue
			case 4:
				history.History()
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
}

func Result(result int) { //works
	var invalid bool
	tools.ClearScreen() //lets print roll result out!
	go history.AddToHistory(result)
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
	1. Roll again!
	2. Back
	3. Exit
	`)
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var NewResult int
			NewResult = tools.RollDice()
			Result(NewResult)

		case 2:
			main()
		case 3:
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
			tools.AddDice()
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
					tools.ResetDices()
					tools.InitializeDiceFile()
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
		fmt.Println(`input 0 to return

	`)

		var choice string

		fmt.Scan(&choice)
		if choice == "0" {
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
		}
		choiceint += 1
		choice = strconv.Itoa(choiceint)
		//Edit first line to be latest dice index
		tools.EditFirstLine(choice)
		ChangeDice()
	}
}

func showDices() { //works
	Dices, err := os.ReadFile(tools.Path("dicefile.txt"))
	if err != nil {
		fmt.Println("Error reading Dices:", err)
		return
	}
	fmt.Println("\nDices:")
	lines := strings.Split(string(Dices), "\n")
	for i, line := range lines {
		if i != 0 && line != "" {
			fmt.Printf("%0d - %s\n", i, line)
		}
	}
}
