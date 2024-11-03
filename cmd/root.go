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
