package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init cmdb",
	Long:  "init cmdb db user",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("init cmdb", verbose)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
