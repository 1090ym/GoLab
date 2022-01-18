package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use: "hello",
	Short: "hello children command",
	Long: "this is a hello children command",
	Args: cobra.MinimumNArgs(1),
	Run: runHello,
}

func init() {
	rootCmd.AddCommand(helloCmd)
}

func runHello(cmd *cobra.Command, args []string)  {
	fmt.Println("hello:", args[0])
}
