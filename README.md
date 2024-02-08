# Dice Rolling Project

This project is a simple dice rolling application written in Go. It allows users to roll dice, change dice configurations, view roll history, and more.

## Files Overview

- **main.go**: Contains the main logic for the application, including menu navigation and user input handling.
- **history.go**: Manages the history of dice rolls, including viewing, clearing, and adding rolls to the history file.
- **clearscreen.go**: Provides functions for clearing the terminal screen, including platform-specific implementations.
- **editing_dices.go**: Handles the editing of dice configurations, such as adding new dice or resetting to default configurations.
- **rolls_dice.go**: Contains functions for rolling dice based on the current configuration.
- **validate_and_paths.go**: Provides validation functions and file path management utilities.

## Functionality

- **Rolling Dice**: Users can roll dice configured with different numbers of sides.
- **Changing Dice Configuration**: Users can add new dice configurations or reset to default configurations.
- **Viewing Roll History**: Users can view a history of previous dice rolls.
- **Clearing Roll History**: Users can clear the history of dice rolls.

## Getting Started

1. Ensure you have Go installed on your system.
2. Clone or download this repository to your local machine.
3. Navigate to the project directory.
4. Run `go run main.go` to start the application.

 OR

1. Use exec (exe for windows) file provided in Exec directory. (.Exe is not tested. I have no idea if it works.)

## Usage

- Use the number keys to navigate through the menu options.
- Follow the on-screen instructions to perform actions such as rolling dice or managing configurations.
- Classic, "press any key to continue" is implemented in history view menu.

## Note

- This project is experimental and may contain bugs or incomplete features.
- Feel free to contribute or provide feedback to improve the project.
- This program will create new directory and some .txt files in it to save data. You can delete them any time you want.
- Directory will be made at same location where main.go, or built exec file is.

