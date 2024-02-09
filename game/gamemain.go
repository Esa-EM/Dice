package game

import (
	"fmt"
	"os"
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
	optio2 := "New game"
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
			case i == 0: //I will fill these later
			case i == 1:
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
