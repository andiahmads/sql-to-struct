package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List supported SQL → Go mappings",
	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan("Supported SQL → Go Types:")
		fmt.Println(`
varchar    → string
text       → string
longtext   → string
tinyint(1) → bool
int        → int
bigint     → int64
datetime   → *time.Time
timestamp  → *time.Time
enum       → string
`)
	},
}

func init() {
	AddCommand(listCmd)
}
