/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"regexp"
	"rename-cli/internal/renamer"

	"github.com/spf13/cobra"
)

// appendCmd represents the append command
var appendCmd = &cobra.Command{
	Use:   "append [flags] <apendix> [files...]",
	Args:  cobra.MinimumNArgs(2),
	Short: "Inserts to the end of filenames with optional numbering and date formatting",
	Long: `Inserts to the end of filenames with optional numbering and date formatting

Usage Example:
rename-cli append -n 1 --dry-run ' - %YYYY-%MM-%DD' *.jpg

Key Features:
  - %n: Replaced with a sequential number starting from the --start value (default 1). Example: %0n → 01, %n → 1, %00n → 001.
  - Date/Time Placeholders: %YYYY (4-digit year), %YY (2-digit year), %MM (2-digit month), %DD (2-digit day), %hh (2-digit hour), %mm (2-digit minute), %ss (2-digit second).
  - --dry-run: Simulates replacements without modifying files.

Explanation of % Placeholders:
  The % syntax allows dynamic replacements in the replacement string:
    1. Numbering: %n → Sequential number (e.g., %0n → 01, %n → 1). Leading zeros can be specified with %0n, %00n, %000n, etc.
    2. Date/Time: Use %YYYY, %MM, %DD, etc., to insert current date/time values. These are replaced with the current system time when the command runs. This makes the command flexible for batch renaming, versioning, or timestamping files."
`,
	Run: func(cmd *cobra.Command, args []string) {
		start, err := cmd.Flags().GetInt("start")
		if err != nil {
			fmt.Println(err)
		}

		dry, err := cmd.Flags().GetBool("dry-run")
		if err != nil {
			fmt.Println(err)
		}

		appendix := args[0]
		files := args[1:]

		fmt.Printf("%d files selected\n", len(files))

		re := regexp.MustCompile(`\..*$`)
		for _, file := range files {
			extension := re.FindString(file)
			replaced := re.ReplaceAllString(file, appendix) + extension

			if !re.MatchString(file) {
				replaced += appendix
			}

			replaced = renamer.Enumerate(replaced, start)
			start++

			replaced = renamer.InsertDate(replaced)

			_, err := os.Stat(replaced)
			if file != replaced && !os.IsNotExist(err) {
				fmt.Printf("%s can't be renamed to %s. File already exist\n", file, replaced)
				continue
			}

			if dry {
				fmt.Printf("%s -> %s\n", file, replaced)
			} else {
				os.Rename(file, replaced)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(appendCmd)
	appendCmd.Flags().IntP("start", "n", 1, "Starting value for %n numbering")
	appendCmd.Flags().Bool("dry-run", false, "Show what would be renamed without making changes")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// appendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// appendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
