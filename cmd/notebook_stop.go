package cmd

import (
	"time"

	"github.com/garretthoffman/sagecli/console"
	"github.com/garretthoffman/sagecli/sagemaker"
	"github.com/spf13/cobra"
)

type notebookStopOperation struct {
	sagemaker            sagemaker.Client
	notebookInstanceName string
}

func (o notebookStopOperation) execute() {
	console.Debug("Describing Notebook Instance: %s [API=sagemaker Action=DescribeNotebookInstance]", o.notebookInstanceName)
	notebookInstance, err := o.sagemaker.DescribeNotebookInstance(o.notebookInstanceName)

	if err != nil {
		console.Error(err, "No notebook instance %s", o.notebookInstanceName)
		return
	}

	if notebookInstance.NotebookInstanceStatus != "InService" {
		console.Info("Notebook %s is not currently in service", o.notebookInstanceName)
		return
	}

	console.Debug("Describing Notebook Instance: %s [API=sagemaker Action=StopNotebookInstance]", o.notebookInstanceName)
	err = o.sagemaker.StopNotebookInstance(o.notebookInstanceName)

	if err != nil {
		console.Error(err, "Error stopping notebook instance %s", o.notebookInstanceName)
		return
	}

	console.Info("Stopping notebook instance %s", o.notebookInstanceName)

	notebookStatus := "Stopping"
	for notebookStatus != "Stopped" {
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
	console.Info("Notebook instance %s stopped", o.notebookInstanceName)
}

var notebookStopCmd = &cobra.Command{
	Use:   "stop <notebook-instance-name>",
	Short: "Stop notebook instance",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		notebookStopOperation{
			sagemaker:            sagemaker.New(sess),
			notebookInstanceName: args[0],
		}.execute()
	},
}

func init() {
	notebookCmd.AddCommand(notebookStopCmd)
}
