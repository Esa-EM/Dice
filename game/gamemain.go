package game

import (
	"dice/tools"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/eiannone/keyboard"
)

func StartGame() {
	//add options for new game, remember to make variant out of that filename! we need it!
	//continue saved game
	//quit game
	Clear()
	time.Sleep(1 * time.Second)
	message := "You are about to face brutal real AI in GAME OF DICE!"
	fmt.Printf("\033[1m%s\033[0m", message)
	time.Sleep(3 * time.Second)
	fmt.Println(`

	There will be no mercy or instructions (use arrow keys and enter to select stuff)
	`)

	time.Sleep(2 * time.Second)
	fmt.Println("Mostly coz I am too lazy to write everything...")
	time.Sleep(2 * time.Second)
	fmt.Println(`
	Hit esc to exit game whenever you want. 
	It might not work all the time, but try again in second.`)
	time.Sleep(2 * time.Second)
	for {
		fmt.Println(`
	Sometimes use enter to proceed. Like now.`)
		_, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error reading key:", err)
			time.Sleep(1 * time.Second)
		}
		switch key {
		case keyboard.KeyEnter:
			maingame()
		case keyboard.KeyEsc:
			Clear()
			fmt.Println(`press esc again to exit, anything else to go back`)
			_, key, err := keyboard.GetKey()
			if err != nil {
				fmt.Println("Error reading key:", err)
				time.Sleep(1 * time.Second)
			}
			if key != keyboard.KeyEsc {
				continue
			} else {
				return
			}
		}

	}
}

func maingame() {

	Clear()

	optio1 := "Continue saved game"
	optio2 := "New game (Deletes old save file!)"
	optio3 := "I dont know what you want, press esc to quit..."

	var i int

	for {
		if i == 0 {
			Clear()
			fmt.Printf("\n\033[1mGAME OF DICE!\033[0m\n\n")
			fmt.Printf("\033[1m%s\033[0m\n%s\n%s\n", optio1, optio2, optio3)
		}
		if i == 1 {
			Clear()
			fmt.Printf("\n\033[1mGAME OF DICE!\033[0m\n\n")
			fmt.Printf("%s\n\033[1m%s\033[0m\n%s\n", optio1, optio2, optio3)
		}
		if i == 2 {
			Clear()
			fmt.Printf("\n\033[1mGAME OF DICE!\033[0m\n\n")
			fmt.Printf("%s\n%s\n\033[1m%s\033[0m\n", optio1, optio2, optio3)
		}
		_, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error reading key:", err)
			time.Sleep(1 * time.Second)
		}
		switch key {
		case keyboard.KeyArrowUp:
			if i > 0 {
				i--
			}
		case keyboard.KeyArrowDown:
			if i < 2 {
				i++
			}
		case keyboard.KeyEnter:
			switch {
			case i == 0:
				if _, err := os.Stat(tools.Path("Save1.txt")); os.IsNotExist(err) {
					Clear()
					fmt.Println("There is no save file to load. Starting new game")
					go InitializeSaveFile("Save1")
					time.Sleep(2 * time.Second)
					NewGame()

				} else {
					fmt.Println("this is not implemented yet.")
					time.Sleep(2 * time.Second)
					//LoadGame()
					//implement load savefile and figure rest
				}
			case i == 1:
				Clear()
				fmt.Println("Loading game")
				go InitializeSaveFile("Save1")
				time.Sleep(1 * time.Second)
				NewGame()
			case i == 2:
				Clear()
				fmt.Println(`press esc to exit, anything else to go back`)
				_, key, err := keyboard.GetKey()
				if err != nil {
					fmt.Println("Error reading key:", err)
					time.Sleep(1 * time.Second)
				}
				if key != keyboard.KeyEsc {
					continue
				} else {
					Clear()
					os.Exit(0)
				}
			}
		case keyboard.KeyEsc:
			Clear()
			fmt.Println(`press esc again to exit, anything else to go back`)
			_, key, err := keyboard.GetKey()
			if err != nil {
				fmt.Println("Error reading key:", err)
				time.Sleep(1 * time.Second)
			}
			if key != keyboard.KeyEsc {
				continue
			} else {
				os.Exit(0)
			}

		}
	}

}

func TheGame(money int, owned string, equipped string, shop string) {

	var dice []int
	var Weighted6 bool
	var Weighted5 bool
	var Weighted4 bool
	var SlightlyWeighted6 bool
	var SlightlyWeighted5 bool
	var SlightlyWeighted4 bool
	var Charmlvl1 bool
	var Charmlvl2 bool
	var Charmlvl3 bool
	var User string
	var AI string
	var Userint int
	var AIint int
	bet := money / 10

	dice = defaultDice
	AIdice := dice

	if Weighted6 {
		dice = append(dice, 6, 6)
	}
	if Weighted5 {
		dice = append(dice, 5, 5)
	}
	if Weighted4 {
		dice = append(dice, 4, 4)
	}
	if SlightlyWeighted6 {
		dice = append(dice, 6)
	}
	if SlightlyWeighted5 {
		dice = append(dice, 5)
	}
	if SlightlyWeighted4 {
		dice = append(dice, 4)
	}
	if Charmlvl1 {
		AIdice = append(AIdice, 3)
	}
	if Charmlvl2 {
		AIdice = append(AIdice, 2)
	}
	if Charmlvl3 {
		AIdice = append(AIdice, 1)
	}

	Clear()

	optio1 := "Roll!"
	optio2 := "Set bet"
	optio3 := "Items"
	optio4 := "Shop"

	var x int
	var y int

	for {
		if money > 900000000 {
			money = 900000000
		}
		if x == 0 && y == 0 {
			Clear()
			fmt.Printf("\n\033[1mGAME OF DICE!\033[0m\nMoney:%d€     Bet:%d€\n\n\n", money, bet)
			fmt.Printf("\033[1m%s\033[0m     %s\n%s     %s\n", optio1, optio2, optio3, optio4)
		}
		if x == 1 && y == 0 {
			Clear()
			fmt.Printf("\n\033[1mGAME OF DICE!\033[0m\nMoney:%d€     Bet:%d€\n\n\n", money, bet)
			fmt.Printf("%s     \033[1m%s\033[0m\n%s     %s\n", optio1, optio2, optio3, optio4)
		}
		if x == 0 && y == 1 {
			Clear()
			fmt.Printf("\n\033[1mGAME OF DICE!\033[0m\nMoney:%d€     Bet:%d€\n\n\n", money, bet)
			fmt.Printf("%s     %s\n\033[1m%s\033[0m     %s\n", optio1, optio2, optio3, optio4)
		}
		if x == 1 && y == 1 {
			Clear()
			fmt.Printf("\n\033[1mGAME OF DICE!\033[0m\nMoney:%d€     Bet:%d€\n\n\n", money, bet)
			fmt.Printf("%s     %s\n%s     \033[1m%s\033[0m\n", optio1, optio2, optio3, optio4)
		}
		_, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error reading key:", err)
			time.Sleep(1 * time.Second)
		}
		switch key {
		case keyboard.KeyArrowUp:
			if y > 0 {
				y--
			}
		case keyboard.KeyArrowDown:
			if y < 1 {
				y++
			}
		case keyboard.KeyArrowLeft:
			if x > 0 {
				x--
			}
		case keyboard.KeyArrowRight:
			if x < 1 {
				x++
			}
		case keyboard.KeyEnter:
			switch {
			case x == 0 && y == 0:
				User = strconv.Itoa(tools.Roll(dice))
				AI = strconv.Itoa(tools.Roll(AIdice))
				Clear()
				fmt.Println("You rolled:")
				fmt.Println(DiceArt(User))
				time.Sleep(1 * time.Second)
				fmt.Println("AI rolled:")
				fmt.Println(DiceArt(AI))
				time.Sleep(1 * time.Second)
				Userint, err = strconv.Atoi(User)
				if err != nil {
					fmt.Print("error getting your dice value")
					time.Sleep(1 * time.Second)
				}
				AIint, err = strconv.Atoi(AI)
				if err != nil {
					fmt.Print("error getting AI dice value")
					time.Sleep(1 * time.Second)
				}
				if Userint > AIint && money <= 100 {
					fmt.Printf("\nYou win %d€!\n", bet)
					money += bet
					time.Sleep(3 * time.Second)
				} else if Userint > AIint {
					fmt.Printf("\nYou win %d€!\n", bet)
					money += bet
					time.Sleep(3 * time.Second)
				} else {
					fmt.Printf("\nAI wins! You lose %d€\n", bet)
					money -= bet
					time.Sleep(3 * time.Second)
					if money <= 0 {
						Clear()
						go InitializeSaveFile("save1")
						fmt.Println("You ran out of money. You lose")
						time.Sleep(1 * time.Second)
						fmt.Println(`Press any key to go back`)
						_, _, err := keyboard.GetKey()
						if err != nil {
							fmt.Println("Error reading key:", err)
							break
						}
						maingame()

					}
				}

			case x == 1 && y == 0:
				bet = Bet(money)
				time.Sleep(200 * time.Millisecond)
				continue
			case x == 0 && y == 1:
				fmt.Println("items")
				time.Sleep(1 * time.Second)
				continue
			case x == 1 && y == 1:
				fmt.Println("shop")
				time.Sleep(1 * time.Second)
				continue
			}
		case keyboard.KeyEsc:
			Clear()
			fmt.Println(`press esc again to exit, anything else to go back`)
			_, key, err := keyboard.GetKey()
			if err != nil {
				fmt.Println("Error reading key:", err)
				time.Sleep(1 * time.Second)
			}
			if key != keyboard.KeyEsc {
				continue
			} else {
				os.Exit(0)
			}

		}
	}

}

func Bet(money int) int {
	Clear()

	optio1 := strconv.Itoa(money / 10)
	optio2 := strconv.Itoa(money / 5)
	optio3 := strconv.Itoa(money / 2)
	optio4 := "All IN!"
	var bet int
	var x int
	var y int

	for {
		if x == 0 && y == 0 {
			Clear()
			fmt.Printf("\n\033[1mSet your Bet!\033[0m (Press esc to go back)\n\n")
			fmt.Printf("\033[1m%s€\033[0m     %s€\n%s€     %s\n", optio1, optio2, optio3, optio4)
		}
		if x == 1 && y == 0 {
			Clear()
			fmt.Printf("\n\033[1mSet your Bet!\033[0m (Press esc to go back)\n\n")
			fmt.Printf("%s€     \033[1m%s€\033[0m\n%s€     %s\n", optio1, optio2, optio3, optio4)
		}
		if x == 0 && y == 1 {
			Clear()
			fmt.Printf("\n\033[1mSet your Bet!\033[0m (Press esc to go back)\n\n")
			fmt.Printf("%s€     %s€\n\033[1m%s€\033[0m     %s\n", optio1, optio2, optio3, optio4)
		}
		if x == 1 && y == 1 {
			Clear()
			fmt.Printf("\n\033[1mSet your Bet!\033[0m (Press esc to go back)\n\n")
			fmt.Printf("%s€     %s€\n%s€     \033[1m%s\033[0m\n", optio1, optio2, optio3, optio4)
		}
		_, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error reading key:", err)
			time.Sleep(1 * time.Second)
		}
		switch key {
		case keyboard.KeyArrowUp:
			if y > 0 {
				y--
			}
		case keyboard.KeyArrowDown:
			if y < 1 {
				y++
			}
		case keyboard.KeyArrowLeft:
			if x > 0 {
				x--
			}
		case keyboard.KeyArrowRight:
			if x < 1 {
				x++
			}
		case keyboard.KeyEnter:
			switch {
			case x == 0 && y == 0:
				bet = money / 10
				return bet
			case x == 1 && y == 0:
				bet = money / 5
				return bet
			case x == 0 && y == 1:
				bet = money / 2
				return bet
			case x == 1 && y == 1:
				bet = money
				return bet
			}
		case keyboard.KeyEsc:
			return bet

		}
	}
}
