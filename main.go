package main

import (
	"dice/history"
	"dice/tools"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/eiannone/keyboard"
)

func main() { //works for everything that is implemented for now

	var invalid bool
	var result int
	var uf bool
	mainmenu := `

	1. Roll dice
	2. Change dice
	3. Dice game vs real AI (upcoming)
	4. History
	5. Exit.

	`

	if len(os.Args) != 1 {
		tools.ClearScreen()
		tools.Helper()
	} else if len(os.Args) == 1 {
		tools.InitializeDiceFile()
		//menu printing

		err := keyboard.Open()
		if err != nil {
			fmt.Println("Error opening keyboard:", err)
			os.Exit(1)
		}
		defer keyboard.Close()
		for {
			tools.ClearScreen()
			if invalid {
				fmt.Println("Invalid choice. Please use valid number.")
				invalid = false
				fmt.Println(mainmenu)
			} else if uf {
				fmt.Println("You selected upcoming feature. it isnt there yet, wait for next version")
				uf = false
				fmt.Println(mainmenu)

			} else {
				fmt.Println("Select with number:")

				fmt.Println(mainmenu)
				char, key, err := keyboard.GetKey()
				if err != nil {
					fmt.Println("Error reading key:", err)
					break
				}
				switch {
				case char == '1':
					result = tools.RollDice()
					Result(result)

				case char == '2':

					ChangeDice()
				case char == '3':
					uf = true
					continue
				case char == '4':
					history.History()
					continue
				case char == '5' || key == keyboard.KeyEsc:
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
}

func Result(result int) { //works
	var invalid bool
	resultmenu := `
	1. Roll again!
	2. Back
	3. Exit
	`
	tools.ClearScreen() //lets print roll result out!
	go history.AddToHistory(result)
	for i := 0; i < 1; i++ {
		fmt.Println("Rolling")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(`Rolling
	`)
	time.Sleep(1 * time.Second)
	fmt.Println(result)
	time.Sleep(1 * time.Second)
	fmt.Println(resultmenu)

	for {

		if invalid {
			fmt.Println("Invalid choice. Please use valid number.")
			fmt.Println(`Rolled:
			`)
			invalid = false
			fmt.Println(result)
			fmt.Println(resultmenu)
		}
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error reading key:", err)
			break
		}

		switch {
		case char == '1':
			var NewResult int
			NewResult = tools.RollDice()
			Result(NewResult)

		case char == '2':
			main()
		case char == '3' || key == keyboard.KeyEsc:
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
	var reset bool
	var invalid bool
	changemenu := `
	
	1. Select dice
	2. Make new dice
	3. Reset dices (restores default dice set)
	4. Back

	
	
	`
	err := keyboard.Open()
	if err != nil {
		fmt.Println("Error opening keyboard:", err)
		os.Exit(1)
	}
	defer keyboard.Close()
	for {
		tools.ClearScreen()
		if invalid {
			fmt.Println("Invalid choice. Please use valid number.")
			invalid = false
		} else if reset {
			fmt.Println("Dice set is now back to default")
			fmt.Println("Select with number:")
			reset = false
		} else {
			fmt.Println("Select with number:")
		}
		fmt.Println(changemenu)
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error reading key:", err)
			break
		}

		switch {
		case char == '1':
			selDices()
		case char == '2':
			tools.AddDice()
		case char == '3':
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
				char, key, err := keyboard.GetKey()
				if err != nil {
					fmt.Println("Error reading key:", err)
					break
				}

				if char == '1' {
					if tools.ResetDices() {
						reset = true
					}
					tools.InitializeDiceFile()
				} else if char == '2' || key == keyboard.KeyEsc {
					main()
				} else {
					invalid = true
					continue
				}

			}

		case char == '4' || key == keyboard.KeyEsc:
			main()
		default:
			invalid = true
			continue

		}
	}
}

func selDices() []int { // works
	var invalid bool
	keyboard.Close()
	for {
		tools.ClearScreen()
		tools.ShowDices()
		fmt.Println("")
		if invalid {
			fmt.Println("Invalid choice. Please use valid number.")
			invalid = false
		} else {
			fmt.Println("Select dice by giving its number and pressing enter:")
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
		keyboard.Open()
		ChangeDice()
	}
}
