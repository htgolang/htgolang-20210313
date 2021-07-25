package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var verbose bool

var rootCmd = &cobra.Command{
	Use:   "cmdb",
	Short: "cmdb cmd",
	Long:  "cmdb command line",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("cmdb run", verbose)
		return nil
	},
}

func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
