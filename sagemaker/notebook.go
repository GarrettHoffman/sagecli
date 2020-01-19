package sagemaker

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	awsSagemaker "github.com/aws/aws-sdk-go/service/sagemaker"
)

// Notebook represents a Sagemaker notebook instance.
type NotebookInstance struct {
	AcceleratorTypes 					[]string
	AdditionalCodeRepositories 			[]string
	CreationTime 						time.Time
	DefaultCodeRepository				string
	DirectInternetAccess				string
	FailureReason						string
	InstanceType						string
	KmsKeyId							string
	LastModifiedTime					time.Time
	NetworkInterfaceId					string
	NotebookInstanceArn					string
	NotebookInstanceLifecycleConfigName	string
	NotebookInstanceName 				string
	NotebookInstanceStatus				string
	RoleArn								string
	RootAccess							string
	SecurityGroups						[]string
	SubnetId							string
	Url									string
	VolumeSizeInGB						int64
}

// Notebooks is a collection of Sagemaker notebook instances.
type NotebookInstances []NotebookInstance

func (sagemaker SDKClient) ListNotebookInstances() (NotebookInstances, error) {
	return sagemaker.listNotebookInstances(&awsSagemaker.ListNotebookInstancesInput{})
}

func (sagemaker SDKClient) listNotebookInstances(i *awsSagemaker.ListNotebookInstancesInput) (NotebookInstances, error) {
	var notebookInstances NotebookInstances

	handler := func(resp *awsSagemaker.ListNotebookInstancesOutput, lasPage bool) bool {
		for _, notebookInstance := range resp.NotebookInstances {
			notebookInstances = append(notebookInstances, 
				NotebookInstance{
					AdditionalCodeRepositories: 			aws.StringValueSlice(notebookInstance.AdditionalCodeRepositories),
					CreationTime: 							aws.TimeValue(notebookInstance.CreationTime),
					DefaultCodeRepository:					aws.StringValue(notebookInstance.DefaultCodeRepository),
					InstanceType:							aws.StringValue(notebookInstance.InstanceType),
					LastModifiedTime:						aws.TimeValue(notebookInstance.LastModifiedTime),
					NotebookInstanceArn:					aws.StringValue(notebookInstance.NotebookInstanceArn),
					NotebookInstanceLifecycleConfigName:	aws.StringValue(notebookInstance.NotebookInstanceLifecycleConfigName),
					NotebookInstanceName: 					aws.StringValue(notebookInstance.NotebookInstanceName),
					NotebookInstanceStatus:					aws.StringValue(notebookInstance.NotebookInstanceStatus),
					Url:									aws.StringValue(notebookInstance.Url),
				},
			)
		}

		return true
	}

	err := sagemaker.client.ListNotebookInstancesPages(i, handler)

	return notebookInstances, err
}
