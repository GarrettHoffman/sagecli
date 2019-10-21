package cmd

import (
	"github.com/spf13/cobra"
)

var endpointCmd = &cobra.Command{
	Use:	"endpoint",
	Short:	"Manage model endpoints",

	//TODO: add long description
	Long: 	`Manage model enpoints`,
}

func init() {
	rootCmd.AddCommand(endpointCmd)
}