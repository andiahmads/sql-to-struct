package cmd

import (
	"fmt"
	"os"
	"sqlStruct/parser"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert CREATE TABLE SQL to GO struct",
	Run: func(cmd *cobra.Command, args []string) {
		validate := func(input string) error {
			if len(input) == 0 {
				return fmt.Errorf("file cannot be empty")
			}
			return nil
		}

		prompt := promptui.Prompt{
			Label:     color.GreenString("SQL File Path"),
			Validate:  validate,
			AllowEdit: true,
		}

		file, err := prompt.Run()
		if err != nil {
			color.Red("Prompt Failed:%v", err)
			return
		}

		sqlBytes, err := os.ReadFile(file)
		if err != nil {
			color.Red("Cannot read file: %v", err)
			return
		}
		color.Yellow("PARSING SQL.....")

		result, err := parser.ConvertSQLToStruct(string(sqlBytes), "")
		if err != nil {
			color.Red("Error: %v", err)
			return
		}
		color.Green("\n====Generated Struct ====")
		fmt.Println(color.BlueString(result))

	},
}

func init() {
	AddCommand(convertCmd)
}
