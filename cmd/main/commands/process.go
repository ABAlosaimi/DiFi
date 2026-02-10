package commands

import (
	"fmt"
	"log"
	"github.com/ABAlosaimi/DiFi/internal/combiner"
	"github.com/spf13/cobra"
)

var (
	filterLine int
	projectDir string
)

var processCmd = &cobra.Command {
	Use:   "process [files...]",
	Short: "Process one or more feature files",
	Long:  `Process feature files by filtering content starting from a specific line number
		    and copying them to a target project directory.

			Examples:
  				 difi process -f 5 -p /path/to/project file1.feature
 				 difi process -f 10 -p /path/to/project file1.feature file2.feature file3.feature
 				 difi process --filter 3 --project ./myproject *.feature`,
    Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Validate required flags
		if projectDir == "" {
			log.Fatal("Project directory is required (use -p or --project)")
		}

		fmt.Printf("Processing %d file(s) with filter starting at line %d\n", len(args), filterLine)
		fmt.Printf("Target project directory: %s\n\n", projectDir)

		// processing singel file 
		if len(args) == 1 {
			file, fileName, err := combiner.OpenFeatureFile(args[0])
			if err != nil {
				fmt.Printf("something went worng with opening your file: %v", err)
			}

			filteredFile, err:= combiner.FilterFeatureFile(file, filterLine)

			if _ , err := combiner.WriteFeatureFile(filteredFile, projectDir, fileName); err != nil {
				fmt.Print("the feature file copied sucessfully")
			}
		}

		// Process all files in case of multi file processing 
		err := combiner.ProcessMultipleFeatureFiles(args, filterLine, projectDir)
		if err != nil {
			log.Fatalf("Error processing files: %v", err)
		}

		fmt.Printf("\nSuccessfully processed all %d file(s)\n", len(args))
	},
}

func init() {
	processCmd.Flags().IntVarP(&filterLine, "filter", "f", 0, "Starting line number for filtering (0 means no filtering)")
	processCmd.Flags().StringVarP(&projectDir, "project", "p", "", "Target project directory (required)")
	processCmd.MarkFlagRequired("project")

	
	rootCmd.AddCommand(processCmd)
}
