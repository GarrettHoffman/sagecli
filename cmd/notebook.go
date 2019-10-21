package cmd

import (
	"github.com/spf13/cobra"
)

var notebookCmd = &cobra.Command{
	Use:	"notebook",
	Short:	"Manage hosted notebook instances",
	Long: 	`Manage hosted notebook instance

Notebooks are jupyter environments hosted on AWS Sagemaker Notebook. Instances of a notebook 
will be kept running until you manually stop them either through AWS APIs, the AWS Management
Console, or sage notebook stop.`,
}

func init() {
	rootCmd.AddCommand(notebookCmd)
}