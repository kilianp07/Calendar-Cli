package main

import (
	"fmt"
	"os"

	"github.com/kilianp07/Calendar-Cli/app"
	"github.com/spf13/cobra"
)

var (
	month  int
	year   int
	format string
	output string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "Calendar CLI",
		Short: "A CLI to generate calendars",
		Run: func(cmd *cobra.Command, args []string) {
			// Check if all flags are set
			if month == 0 || year == 0 || output == "" || format == "" {
				fmt.Println("All flags (month, year, output, format) must be set.")
				cmd.Help()
				os.Exit(1)
			}
			app.Start(month, year, format, output)
		},
	}

	// Define flags for the root command
	rootCmd.Flags().IntVarP(&month, "month", "m", 0, "Specify the month")
	rootCmd.Flags().IntVarP(&year, "year", "y", 0, "Specify the year")
	rootCmd.Flags().StringVarP(&format, "format", "f", "", "Specify the output format")
	rootCmd.Flags().StringVarP(&output, "output", "o", "", "Specify the output file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
