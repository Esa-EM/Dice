package tools

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const diceFile = "dice.txt"

func ClearScreen() {
	// Clearing the screen depends on the operating system
	switch runtime.GOOS {
	case "linux", "darwin": // Unix-like systems
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		diceimg() //adding dice logo for unix.
	case "windows": // Windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		diceimg() //adding dice logo for windows too
	}

}

func diceimg() {
	dice := `
      -------
    / *   * /|
   / *   * / |   
    -------  |    
  |   *   | *|     
  |       | /     
  |   *   |/
   -------
  
 
   `
	fmt.Println(dice)
}
func Clear() {
	// Clearing the screen depends on the operating system
	switch runtime.GOOS {
	case "linux", "darwin": // Unix-like systems
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows": // Windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

}

func Helper() {
	fmt.Println(`Usage: ./Dice
	This is a dice. It will give you numbers based on dice setup
	You can either make your own dice, or use default one
	Commands:
	1. Roll dice
	2. Change dice
	3. Settings
	4. Exit.
	`)
}
