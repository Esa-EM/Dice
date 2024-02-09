package game

import (
	"os"
	"os/exec"
	"runtime"
)

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
