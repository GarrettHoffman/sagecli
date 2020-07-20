package cmd

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/garretthoffman/sagecli/console"
	"github.com/garretthoffman/sagecli/sagemaker"
)

type notebookDescribeOperation struct {
	sagemaker            sagemaker.Client
	notebookInstanceName string
}

func (o notebookDescribeOperation) execute() {
	console.Debug("Describing Notebook Instance: %s [API=sagemaker Action=DescribeNotebookInstance]", o.notebookInstanceName)
	notebookInstance, err := o.sagemaker.DescribeNotebookInstance(o.notebookInstanceName)

	if err != nil {
		console.Error(err, "Could not describe notebook instance %s", o.notebookInstanceName)
		return
	}

	console.Header("Summary")
	console.KeyValue("Notebook Instance Name", "%s\n", notebookInstance.NotebookInstanceName)
	console.KeyValue("Notebook Instance ARN", "%s\n", notebookInstance.NotebookInstanceArn)
	console.KeyValue("Notebook Url", "%s\n", notebookInstance.Url)
	console.KeyValue("Creation Time", "%s\n", notebookInstance.CreationTime.Format("2006-01-02 15:04:05"))
	console.KeyValue("Last Modified Time", "%s\n", notebookInstance.LastModifiedTime.Format("2006-01-02 15:04:05"))
	console.KeyValue("IAM Role ARN", "%s\n", notebookInstance.RoleArn)
	console.Header("Status")
	console.KeyValue("Notebook Status", "%s\n", notebookInstance.NotebookInstanceStatus)

	if notebookInstance.FailureReason != "" {
		console.KeyValue("Failure Reason", "%s\n", notebookInstance.FailureReason)
	}
	console.Header("Hardware")
	console.KeyValue("Instance Type", "%s\n", notebookInstance.InstanceType)
	if strings.Join(notebookInstance.AcceleratorTypes, ", ") != "" {
		console.KeyValue("Accelerator Types", "%s\n", strings.Join(notebookInstance.AcceleratorTypes, ", "))
	}
	console.KeyValue("Volume Size (GB)", "%d\n", notebookInstance.VolumeSizeInGB)

	if notebookInstance.DefaultCodeRepository != "" {
		console.Header("Code Repositories")
		console.KeyValue("Default Repository", "%s\n", notebookInstance.DefaultCodeRepository)
		if strings.Join(notebookInstance.AdditionalCodeRepositories, ", ") != "" {
			console.KeyValue("Additional Repositories", "%s\n", strings.Join(notebookInstance.AdditionalCodeRepositories, ", "))
		}
	}

	if notebookInstance.NotebookInstanceLifecycleConfigName != "" {
		console.Header("Lifecycle Policy")
		console.KeyValue("Lifecycle Policy Name", "%s\n", notebookInstance.NotebookInstanceLifecycleConfigName)
	}

	console.Header("Networking")
	console.KeyValue("Direct Internet Access", "%s\n", notebookInstance.DirectInternetAccess)
	if notebookInstance.NetworkInterfaceId != "" {
		console.KeyValue("Network Interface ID", "%s\n", notebookInstance.NetworkInterfaceId)
	}
	if notebookInstance.SubnetId != "" {
		console.KeyValue("Subnet ID", "%s\n", notebookInstance.SubnetId)
	}
	if strings.Join(notebookInstance.SecurityGroups, ", ") != "" {
		console.KeyValue("Security Groups", "%s\n", strings.Join(notebookInstance.SecurityGroups, ", "))
	}

}

var notebookDescribeCmd = &cobra.Command{
	Use:   "describe <notebook-instance-name>",
	Short: "Inspect notebook instance",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		notebookDescribeOperation{
			sagemaker:            sagemaker.New(sess),
			notebookInstanceName: args[0],
		}.execute()
	},
}

func init() {
	notebookCmd.AddCommand(notebookDescribeCmd)
}
