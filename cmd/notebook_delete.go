package cmd

import (
	"time"

	"github.com/garretthoffman/sagecli/console"
	"github.com/garretthoffman/sagecli/sagemaker"
	"github.com/spf13/cobra"
)

type notebookDeleteOperation struct {
	sagemaker            sagemaker.Client
	notebookInstanceName string
}

func (o notebookDeleteOperation) execute() {
	console.Debug("Describing Notebook Instance: %s [API=sagemaker Action=DescribeNotebookInstance]", o.notebookInstanceName)
	notebookInstance, err := o.sagemaker.DescribeNotebookInstance(o.notebookInstanceName)

	if err != nil {
		console.Error(err, "No notebook instance %s", o.notebookInstanceName)
		return
	}

	if notebookInstance.NotebookInstanceStatus != "Stopped" {
		console.Info("Notebook instance status must be Stopped to delete status, run sage notebook stop %s", o.notebookInstanceName)
		return
	}

	console.Debug("Describing Notebook Instance: %s [API=sagemaker Action=DeleteNotebookInstance]", o.notebookInstanceName)
	err = o.sagemaker.DeleteNotebookInstance(o.notebookInstanceName)

	if err != nil {
		console.Error(err, "Error deleting notebook instance %s", o.notebookInstanceName)
		return
	}

	console.Info("Deleting notebook instance %s", o.notebookInstanceName)
	for {
		time.Sleep(5000000000)
		print(".")

		console.Debug("Describing Notebook Instance: %s [API=sagemaker Action=DescribeNotebookInstance]", o.notebookInstanceName)
		notebookInstance, err = o.sagemaker.DescribeNotebookInstance(o.notebookInstanceName)

		if err != nil {
			print("\n")
			console.Info("Notebook instance %s deleted", o.notebookInstanceName)
			return
		}
	}
}

var notebookDeleteCmd = &cobra.Command{
	Use:   "delete <notebook-instance-name>",
	Short: "Delete notebook instance",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		notebookDeleteOperation{
			sagemaker:            sagemaker.New(sess),
			notebookInstanceName: args[0],
		}.execute()
	},
}

func init() {
	notebookCmd.AddCommand(notebookDeleteCmd)
}
