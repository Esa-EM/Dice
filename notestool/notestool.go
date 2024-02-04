package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "help" {
		fmt.Println("Usage: ./notestool [COLLECTION]")
		fmt.Println("Tool for managing short single-line notes.")
		fmt.Println("If COLLECTION does not exist, it will be created.")
		fmt.Println("Commands:")
		fmt.Println("1. Show notes.")
		fmt.Println("2. Add a note.")
		fmt.Println("3. Delete a note.")
		fmt.Println("4. Exit.")
		return
	}

	collection := os.Args[1]
	filename := collection + ".txt"

	fmt.Println("Welcome to the notes tool!")

	for {
		fmt.Println("\nSelect operation:")
		fmt.Println("1. Show notes.")
		fmt.Println("2. Add a note.")
		fmt.Println("3. Delete a note.")
		fmt.Println("4. Exit.")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			showNotes(filename)
		case 2:
			addNote(filename)
		case 3:
			deleteNote(filename)
		case 4:
			fmt.Println("Exiting the program. Goodbye!")

			// Check for empty file and delete if needed
			checkAndDeleteEmptyFile(filename)

			return
		default:
			fmt.Println("Invalid choice. Please choose a valid operation.")
		}
	}
}

func showNotes(filename string) {
	notes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading notes:", err)
		return
	}

	fmt.Println("\nNotes:")
	lines := strings.Split(string(notes), "\n")
	for i, line := range lines {
		if line != "" {
			fmt.Printf("%03d - %s\n", i+1, line)
		}
	}
}

func addNote(filename string) {
	fmt.Print("\nEnter the note text: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	newNote := scanner.Text()

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error adding note:", err)
		return
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, newNote)
	if err != nil {
		fmt.Println("Error adding note:", err)
	}
}

func deleteNote(filename string) {
	showNotes(filename)
	fmt.Print("\nEnter the number of note to remove or 0 to cancel: ")

	var choice int
	fmt.Scan(&choice)

	if choice == 0 {
		return
	}

	notes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading notes:", err)
		return
	}

	lines := strings.Split(string(notes), "\n")

	if choice < 1 || choice > len(lines) {
		fmt.Println("Invalid note number.")
		return
	}

	lines = append(lines[:choice-1], lines[choice:]...)
	newContent := strings.Join(lines, "\n")

	err = os.WriteFile(filename, []byte(newContent), 0644)
	if err != nil {
		fmt.Println("Error deleting note:", err)
		return
	}

	// Check if the file is empty, and delete it if needed
	if len(lines) == 0 {
		err = os.Remove(filename)
		if err != nil {
			fmt.Println("Error deleting file:", err)
		}
	}
}

func checkAndDeleteEmptyFile(filename string) {
	notes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading notes:", err)
		return
	}

	if len(strings.TrimSpace(string(notes))) == 0 {
		err = os.Remove(filename)
		if err != nil {
			fmt.Println("Error deleting empty file:", err)
		}
	}
}
