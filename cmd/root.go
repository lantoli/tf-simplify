package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

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
	DIR      = "./"
	//	DIR      = "../example-tf/local/"
	FILE = DIR + "main.tf"
)

func Run(cmd *cobra.Command, args []string) {
	changes, err := hasChanges()
	if err != nil {
		fmt.Println("Error running terraform plan: ", err)
		os.Exit(1)
	}
	if changes {
		fmt.Println("There are plan changes, tf-simplify can't start")
		os.Exit(1)
	}
	contentByte, err := os.ReadFile(FILE)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}
	contentStr := string(contentByte)
	finalContent := ""
	lines := strings.Split(contentStr, "\n")
	for i, line := range lines {
		keep := true
		if strings.Contains(line, "=") {
			newLines := append([]string(nil), lines...)
			newLines[i] = "#" + newLines[i]
			newContent := strings.Join(newLines, "\n") + "\n"
			if err := os.WriteFile(FILE, []byte(newContent), 0644); err != nil {
				fmt.Println("Error writing file: ", err)
				os.Exit(1)
			}
			changes, err := hasChanges()
			if err == nil && !changes {
				fmt.Print(i+1, ", ")
				keep = false
			}
		}
		if keep {
			finalContent += line + "\n"
		}
	}
	fmt.Println()

	finalContent = removeEmptyBlocks(finalContent)

	if err := os.WriteFile(FILE, []byte(finalContent), 0644); err != nil {
		fmt.Println("Error setting final file: ", err)
		os.Exit(1)
	}
	changes, err = hasChanges()
	if err == nil && !changes {
		fmt.Println("Final file applied successfully")
	} else {
		fmt.Println("Error, Final file has changes")
	}
}

func removeEmptyBlocks(content string) string {
	lines := strings.Split(content, "\n")
	finalContent := make([]string, 0, len(lines))
	i := 0
	for i < len(lines) {
		if i+1 < len(lines) &&
			strings.ContainsRune(lines[i], '{') &&
			strings.ContainsRune(lines[i+1], '}') {
			i += 2
		} else {
			finalContent = append(finalContent, lines[i])
			i++
		}
	}
	return strings.Join(finalContent, "\n")
}

func hasChanges() (bool, error) {
	cmd := exec.Command("terraform", "plan", "-detailed-exitcode")
	cmd.Dir = DIR
	if err := cmd.Run(); err != nil {
		if err.(*exec.ExitError).ExitCode() == EC_DIFF {
			return true, nil
		}
		return false, err
	}
	return false, nil
}
