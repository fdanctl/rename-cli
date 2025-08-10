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

// replaceCmd represents the replace command
var replaceCmd = &cobra.Command{
	Use:   "replace [flags] <pattern> <replacement> [files...]",
	Args:  cobra.MinimumNArgs(3),
	Short: "Replace patterns in filenames with optional numbering and date formatting",
	Long: `Replace patterns in filenames with optional numbering and date formatting

Usage Example:
rename-cli replace -n 1 'old_pattern' 'new_pattern_%0n' *.txt

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

		pattern := args[0]
		replacement := args[1]
		files := args[2:]

		fmt.Printf("%d files selected\n", len(files))

		re := regexp.MustCompile(pattern)
		for _, file := range files {
			trailingPath, filename, extension := renamer.ParseFileName(file)

			replaced := re.ReplaceAllString(filename, replacement)

			replaced = renamer.Enumerate(replaced, start)
			start++

			replaced = renamer.InsertDate(replaced)

			renamedFile := trailingPath + replaced + extension

			_, err := os.Stat(renamedFile)
			if file != renamedFile && !os.IsNotExist(err) {
				fmt.Printf("%s can't be renamed to %s. File already exist\n", file, renamedFile)
				continue
			}

			if dry {
				fmt.Printf("%s -> %s\n", file, renamedFile)
			} else {
				os.Rename(file, renamedFile)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(replaceCmd)
	replaceCmd.Flags().IntP("start", "n", 1, "Starting value for %n numbering")
	replaceCmd.Flags().Bool("dry-run", false, "Show what would be renamed without making changes")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// replaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// replaceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
