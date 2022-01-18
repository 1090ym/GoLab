package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "cobra",
	Short: "command tool",
	Long: "cobra command tool, it's a great go lib",
	Run: runRoot,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func runRoot(cmd *cobra.Command, args []string) {
	fmt.Printf("execute %s args: %v\n", cmd.Name(), args)
}
