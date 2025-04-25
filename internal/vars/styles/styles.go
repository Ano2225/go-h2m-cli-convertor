package styles

import "github.com/fatih/color"

var (
	Title   = color.New(color.FgCyan, color.Bold) // Title style
	Command = color.New(color.FgGreen)            // Command style
	Error   = color.New(color.FgRed, color.Bold)  // Error style
	Help    = color.New(color.FgYellow)           // Help style
)
