package game

import (
	"dice/tools"
	"fmt"
	"os"
	"strconv"
)

func InitializeSaveFile(name string) (bool, error) {
	// Check if the history file exists
	name += ".txt"
	if _, err := os.Stat(tools.Path(name)); os.IsNotExist(err) {
		// If it doesn't exist, create the file
		file, err := os.Create(tools.Path(name))
		if err != nil {
			return false, err
		}
		defer file.Close()
		return true, nil
	}

	return true, nil
}

func OwnedAndEquippedItemsEncoding(owned string, equipped string) string { //owned string will only include items that are not equipped

	var nextItem string
	var newOwned string
	var combined string

	for i := 3; i <= 93 || len(equipped) != 0; i += 3 {
		if i%2 != 0 {
			nextItem = equipped[:3]
			equipped = equipped[3:]
			combined += nextItem
			newOwned += strconv.Itoa(i)
			if len(newOwned) < 2 {
				newOwned = "0" + newOwned
			}
			combined += newOwned
			newOwned = ""
		}
	}

	for i := 2; i <= 98 || len(owned) != 0; i += 2 {
		if i%3 != 0 {
			nextItem = owned[:3]
			owned = owned[3:]
			combined += nextItem
			newOwned += strconv.Itoa(i)
			if len(newOwned) < 2 {
				newOwned = "0" + newOwned
			}
			combined += newOwned
			newOwned = ""
		}
	}

	return combined
}

func Savefile(money int, owned string, savefile string) bool {

	var toSave string
	var moneyLenStr string
	var moneyString string

	money = money + 584239

	moneyString = strconv.Itoa(money)
	moneylenght := len(moneyString)
	moneyLenStr = strconv.Itoa(moneylenght)
	if len(moneyLenStr) == 1 {
		moneyLenStr = "0" + moneyLenStr
	}

	toSave += moneyLenStr + moneyString + owned

	file, err := os.OpenFile(tools.Path(savefile), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening savefile", err)
		return false
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, toSave)
	if err != nil {
		fmt.Println("Error saving data", err)
		return false
	}

	return true
}
