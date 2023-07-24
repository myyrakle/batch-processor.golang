package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var processor string

var rootCmd = &cobra.Command{
	Use:   "batch",
	Short: "run batch processes",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(processor)
		// Do Stuff Here
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&processor, "processor", "p", "", "Processor to use")
	rootCmd.MarkFlagRequired("processor")
}
