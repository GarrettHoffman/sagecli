package cmd

import (
	"github.com/spf13/cobra"
)

var trainCmd = &cobra.Command{
	Use:	"train",
	Short:	"Manage model training jobs",

	//TODO: add long description
	Long: 	`Manage model training jobs`,
}

func init() {
	rootCmd.AddCommand(trainCmd)
}