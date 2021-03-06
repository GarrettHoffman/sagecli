package cmd

import (
	"time"

	"github.com/garretthoffman/sagecli/console"
	"github.com/garretthoffman/sagecli/sagemaker"
	"github.com/spf13/cobra"
)

type notebookStartOperation struct {
	sagemaker            sagemaker.Client
	notebookInstanceName string
}

func (o notebookStartOperation) execute() {
	console.Debug("Describing Notebook Instance: %s [API=sagemaker Action=DescribeNotebookInstance]", o.notebookInstanceName)
	notebookInstance, err := o.sagemaker.DescribeNotebookInstance(o.notebookInstanceName)

	if err != nil {
		console.Error(err, "No notebook instance %s", o.notebookInstanceName)
		return
	}

	if notebookInstance.NotebookInstanceStatus != "Stopped" {
		console.Info("Notebook %s is not currently stopped", o.notebookInstanceName)
		return
	}

	console.Debug("Describing Notebook Instance: %s [API=sagemaker Action=StartNotebookInstance]", o.notebookInstanceName)
	err = o.sagemaker.StartNotebookInstance(o.notebookInstanceName)

	if err != nil {
		console.Error(err, "Error starting notebook instance %s", o.notebookInstanceName)
		return
	}

	console.Info("Starting notebook instance %s", o.notebookInstanceName)

	notebookStatus := "Pending"
	for notebookStatus != "InService" {
		time.Sleep(5000000000)
		print(".")

		console.Debug("Describing Notebook Instance: %s [API=sagemaker Action=DescribeNotebookInstance]", o.notebookInstanceName)
		notebookInstance, err = o.sagemaker.DescribeNotebookInstance(o.notebookInstanceName)

		if err != nil {
			console.Error(err, "Error fetching notebook instance status")
			return
		}

		notebookStatus = notebookInstance.NotebookInstanceStatus
	}

	print("\n")
	console.Info("Notebook instance %s started", o.notebookInstanceName)
}

var notebookStartCmd = &cobra.Command{
	Use:   "start <notebook-instance-name>",
	Short: "Start notebook instance",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		notebookStartOperation{
			sagemaker:            sagemaker.New(sess),
			notebookInstanceName: args[0],
		}.execute()
	},
}

func init() {
	notebookCmd.AddCommand(notebookStartCmd)
}
