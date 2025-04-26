package constants

const Signature = "\n✨ Ano2225 ❤️ - Creator of H2M ✨\n" // Signature for the end of the program

// Template for the help message
const UsageTemplate = `{{.Long | trimTrailingWhitespaces}}

{{if .HasAvailableSubCommands}}{{styleLine "Commandes disponibles:"}}{{range .Commands}}{{if .IsAvailableCommand}}
  {{commandStyle .Name | printf "%-12s"}} {{.Short}}{{end}}{{end}}

{{end}}{{if .HasAvailableLocalFlags}}{{styleLine "Options globales:"}}
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}
{{end}}`
