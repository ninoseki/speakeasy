package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version string

var rootCmd = &cobra.Command{
	Use: "speakeasy",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("42")
		return nil
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Speakeasy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
