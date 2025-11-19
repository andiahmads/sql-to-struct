package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sqlstruct",
	Short: "SQL to Golang struct Generator",
	Long: color.CyanString(`
╔══════════════════════════════╗
║     SQL → Go Struct Tool     ║
║    With Colors + Autosuggest ║
╚══════════════════════════════╝
`),
	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow("Use a command. Example:")
		color.Green(" sqlstruct convert")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func AddCommand(c *cobra.Command) {
	rootCmd.AddCommand(c)

}
