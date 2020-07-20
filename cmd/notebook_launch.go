package cmd

import (
	"github.com/pkg/browser"

	"github.com/garretthoffman/sagecli/console"
	"github.com/garretthoffman/sagecli/sagemaker"
	"github.com/spf13/cobra"
)

type notebookLaunchOperation struct {
	sagemaker            sagemaker.Client
	notebookInstanceName string
}

func (o notebookLaunchOperation) execute() {
	console.Debug("Describing Notebook Instance: %s [API=sagemaker Action=DescribeNotebookInstance]", o.notebookInstanceName)
	notebookInstance, err := o.sagemaker.DescribeNotebookInstance(o.notebookInstanceName)

	if err != nil {
		console.Error(err, "No notebook instance %s", o.notebookInstanceName)
		return
	}

	if notebookInstance.NotebookInstanceStatus != "InService" {
		console.Info("Notebook status must be InService to launch Jupyter, run sage notebook start %s", o.notebookInstanceName)
	}

	jupyterLab := "https://" + notebookInstance.Url + "/lab"
	browser.OpenURL(jupyterLab)
}

var notebookLaunchCmd = &cobra.Command{
	Use:   "launch <notebook-instance-name>",
	Short: "Launch jupyter lab for the notebook instance",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		notebookLaunchOperation{
			sagemaker:            sagemaker.New(sess),
			notebookInstanceName: args[0],
		}.execute()
	},
}

func init() {
	notebookCmd.AddCommand(notebookLaunchCmd)
}
