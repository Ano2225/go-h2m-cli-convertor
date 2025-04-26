package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the console",
	Long: `Command to clear the console screen.This command is useful for cleaning up the console 
	completely.`,
	Run: func(cmd *cobra.Command, args []string) {
		var c *exec.Cmd
		// Determine the operating system and set the command accordingly
		switch os := os.Getenv("OS"); os {
		case "Windows_NT":
			// Windows
			// Use "cls" in cmd consle for Windows
			c = exec.Command("cmd", "/c", "cls")
		default:
			// Unix/Linux/Mac
			// Use "clear" in terminal for Unix/Linux/Mac
			// Use "clear" in powershell for Windows
			c = exec.Command("clear")
		}
		// Set the command's output to the console
		// This will clear the console screen
		c.Stdout = os.Stdout
		err := c.Run()
		// Check for errors
		// If there is an error, print it to the console
		if err != nil {
			fmt.Println("Error clearing the console:", err)
			return
		}
	},
}
