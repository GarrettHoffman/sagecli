package cmd

import (
	"github.com/spf13/cobra"
)

var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "Manage trained models",

	//TODO: add long description
	Long: `Manage trained models`,
}

func init() {
	rootCmd.AddCommand(modelCmd)
}
