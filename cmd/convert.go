package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sqlStruct/parser"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func SelectFileFromPath(path string) (string, error) {
	// jika path adalah file -> selesai
	fi, err := os.Stat(path)
	if err == nil && !fi.IsDir() {
		return path, nil
	}

	// jika bukan folder → ambil folder-nya saja
	dir := path
	if err != nil || !fi.IsDir() {
		dir = filepath.Dir(path)
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	items := []string{}
	for _, f := range files {
		name := f.Name()
		if f.IsDir() {
			name += "/"
		}
		items = append(items, name)
	}

	if len(items) == 0 {
		return "", fmt.Errorf("folder kosong")
	}

	selectPrompt := promptui.Select{
		Label: color.CyanString("Choose file in %s", dir),
		Items: items,
		Size:  12,
	}

	_, chosen, err := selectPrompt.Run()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, chosen), nil
}

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
			Label:     color.GreenString("SQL File Path (press Enter to browse)"),
			Validate:  validate,
			AllowEdit: true,
		}

		file, err := prompt.Run()
		if err != nil {
			color.Red("Prompt Failed:%v", err)
			return
		}

		file, err = SelectFileFromPath(file)
		if err != nil {
			color.Red("File selection failed: %v", err)
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
