package cmd

import (
	"fmt"
	"os"
	"os/exec"

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

	changes, err := hasChanges()
	if err != nil {
		fmt.Println("Error running terraform plan, tf-simplify can't continue: ", err)
		os.Exit(1)
	}

	fmt.Println("Has changes: ", changes)
}

func hasChanges() (bool, error) {
	if err := exec.Command("terraform", "plan", "-detailed-exitcode").Run(); err != nil {
		if err.(*exec.ExitError).ExitCode() == EC_DIFF {
			return true, nil
		}
		return false, err
	}
	return false, nil
}
