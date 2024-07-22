package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go_uninstall [binaries...]",
	Short: "A tool to list and uninstall Go binaries",
	Args:  cobra.ArbitraryArgs, // 任意の引数を受け取る
	Run: func(cmd *cobra.Command, args []string) {
		gobinPath := getGoBinPath()
		for _, binary := range args {
			binaryPath := filepath.Join(gobinPath, binary)
			err := os.Remove(binaryPath)
			if err != nil {
				fmt.Printf("Error removing file %s: %v\n", binary, err)
			} else {
				fmt.Printf("Removed %s\n", binary)
			}
		}
	},
}

func getGoBinPath() string {
	gobinPath := os.Getenv("GOBIN")
	if gobinPath == "" {
		gobinPath = filepath.Join(os.Getenv("GOPATH"), "bin")
		if _, err := os.Stat(gobinPath); os.IsNotExist(err) {
			ex, err := os.Executable()
			if err != nil {
				fmt.Printf("Error getting executable path: %v\n", err)
				os.Exit(1)
			}
			gobinPath = filepath.Dir(ex)
		}
	}
	return gobinPath
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Here you will define your flags and configuration settings.
}
