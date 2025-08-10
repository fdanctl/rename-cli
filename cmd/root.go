/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rename-cli",
	Short: "Batch file renaming utility with append, prepend, and replace options",
	Long: `Batch file renaming utility with append, prepend, and replace options

Key Features:
  - %n: Replaced with a sequential number starting from the --start value (default 1). Example: %0n → 01, %n → 1, %00n → 001.
  - Date/Time Placeholders: %YYYY (4-digit year), %YY (2-digit year), %MM (2-digit month), %DD (2-digit day), %hh (2-digit hour), %mm (2-digit minute), %ss (2-digit second).
  - --dry-run: Simulates replacements without modifying files.

Explanation of % Placeholders:
  The % syntax allows dynamic replacements in the replacement string:
    1. Numbering: %n → Sequential number (e.g., %0n → 01, %n → 1). Leading zeros can be specified with %0n, %00n, %000n, etc.
    2. Date/Time: Use %YYYY, %MM, %DD, etc., to insert current date/time values. These are replaced with the current system time when the command runs. This makes the command flexible for batch renaming, versioning, or timestamping files."
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rename-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
