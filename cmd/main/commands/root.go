package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "difi",
	Short: "DiFi - Feature file combiner and processor",
	Long: `DiFi is a CLI tool that helps you filter and combine feature files.
		   It allows you to process multiple feature files, filter content by line numbers,
		   and copy them to target project directories.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
}
