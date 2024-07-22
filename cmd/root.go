package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tf-simply",
	Short: "Improve your TF files",
	Long:  `Try to simplify your TF files removing default values.`,
	Run:   Run,
}

func Execute() error {
	return rootCmd.Execute()
}

const (
	EC_EMPTY = 0
	EC_ERROR = 1
	EC_DIFF  = 2
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("Hello, World!")
}
