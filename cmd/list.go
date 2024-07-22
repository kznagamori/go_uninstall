package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all installed Go binaries",
	Run: func(cmd *cobra.Command, args []string) {
		gobinPath := getGoBinPath()
		files, err := os.ReadDir(gobinPath)
		if err != nil {
			fmt.Printf("Error reading directory: %v\n", err)
			return
		}

		for _, file := range files {
			if !file.IsDir() {
				fmt.Println(file.Name())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
