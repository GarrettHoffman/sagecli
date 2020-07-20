package cmd

import (
	"sort"

	"github.com/garretthoffman/sagecli/sagemaker"
	"github.com/spf13/cobra"
)

type notebookListOperation struct {
	sagemaker sagemaker.Client
	output    Output
}

func (o notebookListOperation) execute() {
	notebookInstances, err := o.find()

	if err != nil {
		o.output.Fatal(err, "Could not list notebook instances")
		return
	}

	if len(notebookInstances) == 0 {
		o.output.Info("No notebook instances found")
		return
	}

	rows := [][]string{
		[]string{"NAME", "STATUS", "INSTANCE TYPE", "CREATED AT", "URL"},
	}

	sort.Slice(notebookInstances, func(i, j int) bool {
		return notebookInstances[i].NotebookInstanceName < notebookInstances[j].NotebookInstanceName
	})

	for _, notebookInstance := range notebookInstances {
		rows = append(rows,
			[]string{
				notebookInstance.NotebookInstanceName,
				notebookInstance.NotebookInstanceStatus,
				notebookInstance.InstanceType,
				notebookInstance.CreationTime.Format("2006-01-02 15:04:05"),
				notebookInstance.Url,
			},
		)
	}

	o.output.Table("", rows)
}

func (o notebookListOperation) find() (sagemaker.NotebookInstances, error) {
	o.output.Debug("Listing Notebook Instances [API=sagemaker Action=ListNotebookInstances]")
	notebookInstances, err := o.sagemaker.ListNotebookInstances()

	if err != nil {
		return sagemaker.NotebookInstances{}, err
	}

	return notebookInstances, nil
}

var notebookListCmd = &cobra.Command{
	Use:   "list",
	Short: "List sagemaker notebook instances",
	Run: func(cmd *cobra.Command, args []string) {
		notebookListOperation{
			sagemaker: sagemaker.New(sess),
			output:    output,
		}.execute()
	},
}

func init() {
	notebookCmd.AddCommand(notebookListCmd)
}
