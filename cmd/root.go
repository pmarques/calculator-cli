package cmd

import (
	"fmt"
	"os"
	"syscall"

	"github.com/spf13/cobra"
)

var (
	GitCommit = ""
	Version   = "dev"
)

// Similar to UsageTemplate() from cobra command, but with additional "Available Plugins" list.
const usageTemplate = `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}{{$cmds := .Commands}}{{if eq (len .Groups) 0}}

Available Commands:{{range $cmds}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{else}}{{range $group := .Groups}}

{{.Title}}{{range $cmds}}{{if (and (eq .GroupID $group.ID) (or .IsAvailableCommand (eq .Name "help")))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if not .AllChildCommandsHaveGroup}}

Additional Commands:{{range $cmds}}{{if (and (eq .GroupID "") (or .IsAvailableCommand (eq .Name "help")))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

{{- if (and (not .HasParent) (gt (len pluginList) 0))}}

Available Plugins:
{{- range pluginList}}
  {{.}}
{{- end}}
{{- end}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

const versionTemplate = `
{{- with .Name}}{{printf "%s " .}}{{end}}{{printf "version: %s" .Version -}}

{{- if not (eq GitCommit "") }}
Git commit: {{ GitCommit }}
{{- end}}
`

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Used for (Cobra) flags.
	var Debug bool

	// rootCmd represents the base command when called without any subcommands.
	rootCmd := &cobra.Command{
		Use:     "calculator",
		Short:   "Calculator is a simple CLI calculator as a cobra example",
		Version: Version,
	}

	rootCmd.SetUsageTemplate(usageTemplate)

	cobra.AddTemplateFunc("GitCommit", func() string { return GitCommit })
	rootCmd.SetVersionTemplate(versionTemplate)

	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "debug output")

	// Testing 2 different initialisatons
	rootCmd.AddCommand(newSubCmd())
	newSumCmd(rootCmd)

	pluginlist := FindPlugins()

	cobra.AddTemplateFunc("pluginList", func() []string { return pluginlist })

	args := os.Args[1:]

	if _, _, err := rootCmd.Find(args); err != nil {
		exCmd, err := FindPlugin(os.Args[1])
		// if we can't find command then execute the normal rootCmd command.
		if err == nil {
			// if we have found the plugin then sysexec it by replacing current process.
			if err := syscall.Exec(exCmd, append([]string{exCmd}, os.Args[2:]...), os.Environ()); err != nil {
				fmt.Fprintf(os.Stderr, "Command finished with error: %v", err)
				os.Exit(int(syscall.ENOENT))
			}
		}
	}

	///nolint:errcheck
	rootCmd.Execute()
}
