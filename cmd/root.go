package cmd

import (
	"fmt"
	"os"
	"time"

	internal "batch/internal"

	"github.com/spf13/cobra"
)

var processorName string

var rootCmd = &cobra.Command{
	Use:   "batch",
	Short: "run batch processes",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()

		fmt.Println("# Processor: ", processorName)
		fmt.Println("--------------------")

		processor := internal.Processors[processorName]

		if processor == nil {
			fmt.Println("Processor not found")
			os.Exit(1)
		}

		processor.Run()

		elapsed := time.Since(start)

		fmt.Println("--------------------")
		fmt.Println("# Elapsed Time: ", elapsed)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&processorName, "processor", "p", "", "Processor to use")
	rootCmd.MarkFlagRequired("processor")
}
