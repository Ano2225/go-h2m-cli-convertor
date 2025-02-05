package utils

import "github.com/fatih/color"

//VARIABLES GLOBALES POUR LA CONFIGURATION
var (
	AppVersion = "1.0.0"
	OutputFile string
	InputFile string
)

//Constantes pour le style
const Logo = `

███    ███  ██████  ██   ██ 
████  ████ ██    ██ ██  ██  
██ ████ ██ ██    ██ █████   
██  ██  ██ ██    ██ ██  ██  
██      ██  ██████  ██   ██ 

    Markdown ⇄ HTML
`

var (
	TitleStyle = color.New(color.FgCyan, color.Bold)
	CommandStyle = color.New(color.FgGreen)
	ErrorStyle = color.New(color.FgRed, color.Bold)
	HelpStyle = color.New(color.FgYellow)
)


// Template d'aide personnalisé
const UsageTemplate = `{{.Long | trimTrailingWhitespaces}}

{{if .HasAvailableSubCommands}}{{styleLine "Commandes disponibles:"}}{{range .Commands}}{{if .IsAvailableCommand}}
  {{commandStyle .Name | printf "%-12s"}} {{.Short}}{{end}}{{end}}

{{end}}{{if .HasAvailableLocalFlags}}{{styleLine "Options globales:"}}
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}
{{end}}`

