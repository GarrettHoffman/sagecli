package cmd

import (
	"github.com/spf13/cobra"
)

var labelCmd = &cobra.Command{
	Use:	"label",
	Short:	"Manage data labeling jobs",

	//TODO: add long description
	Long: 	`Manage data labeling jobs`,
}

func init() {
	rootCmd.AddCommand(labelCmd)
}