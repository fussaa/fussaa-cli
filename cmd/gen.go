// cmd/gen.go
package cmd

import (
	"fmt"

	"github.com/fussaa/fussaa-cli/pkg/parser"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen [file]",
	Short: "Generate methods from struct existing struct(s)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the Go file path.")
			return
		}
		filePath := args[0]
		parser.ProcessFile(filePath)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}
