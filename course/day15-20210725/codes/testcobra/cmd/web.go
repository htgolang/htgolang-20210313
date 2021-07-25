package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addr string

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "web cmdb",
	Long:  "web cmdb",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("web run", addr, verbose)
		return nil
	},
}

func init() {
	webCmd.Flags().StringVarP(&addr, "addr", "s", "localhost:8888", "cmdb web server addr")
	rootCmd.AddCommand(webCmd)
}
