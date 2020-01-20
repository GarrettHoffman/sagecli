package sagemaker

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	awsSagemaker "github.com/aws/aws-sdk-go/service/sagemaker"
)

// Notebook represents a Sagemaker notebook instance.
type NotebookInstance struct {
	AcceleratorTypes                    []string
	AdditionalCodeRepositories          []string
	CreationTime                        time.Time
	DefaultCodeRepository               string
	DirectInternetAccess                string
	FailureReason                       string
	InstanceType                        string
	KmsKeyId                            string
	LastModifiedTime                    time.Time
	NetworkInterfaceId                  string
	NotebookInstanceArn                 string
	NotebookInstanceLifecycleConfigName string
	NotebookInstanceName                string
	NotebookInstanceStatus              string
	RoleArn                             string
	RootAccess                          string
	SecurityGroups                      []string
	SubnetId                            string
	Url                                 string
	VolumeSizeInGB                      int64
}

// Notebooks is a collection of Sagemaker notebook instances.
type NotebookInstances []NotebookInstance

func (sagemaker SDKClient) ListNotebookInstances() (NotebookInstances, error) {
	return sagemaker.listNotebookInstances(&awsSagemaker.ListNotebookInstancesInput{})
}

func (sagemaker SDKClient) DescribeNotebookInstance(name string) (NotebookInstance, error) {
	return sagemaker.describeNotebookInstance(
		&awsSagemaker.DescribeNotebookInstanceInput{
			NotebookInstanceName: &name,
		},
	)
}

func (sagemaker SDKClient) listNotebookInstances(i *awsSagemaker.ListNotebookInstancesInput) (NotebookInstances, error) {
	var notebookInstances NotebookInstances

	handler := func(resp *awsSagemaker.ListNotebookInstancesOutput, lasPage bool) bool {
		for _, notebookInstance := range resp.NotebookInstances {
			notebookInstances = append(notebookInstances,
				NotebookInstance{
					AdditionalCodeRepositories:          aws.StringValueSlice(notebookInstance.AdditionalCodeRepositories),
					CreationTime:                        aws.TimeValue(notebookInstance.CreationTime),
					DefaultCodeRepository:               aws.StringValue(notebookInstance.DefaultCodeRepository),
					InstanceType:                        aws.StringValue(notebookInstance.InstanceType),
					LastModifiedTime:                    aws.TimeValue(notebookInstance.LastModifiedTime),
					NotebookInstanceArn:                 aws.StringValue(notebookInstance.NotebookInstanceArn),
					NotebookInstanceLifecycleConfigName: aws.StringValue(notebookInstance.NotebookInstanceLifecycleConfigName),
					NotebookInstanceName:                aws.StringValue(notebookInstance.NotebookInstanceName),
					NotebookInstanceStatus:              aws.StringValue(notebookInstance.NotebookInstanceStatus),
					Url:                                 aws.StringValue(notebookInstance.Url),
				},
			)
		}

		return true
	}

	err := sagemaker.client.ListNotebookInstancesPages(i, handler)

	return notebookInstances, err
}

func (sagemaker SDKClient) describeNotebookInstance(i *awsSagemaker.DescribeNotebookInstanceInput) (NotebookInstance, error) {
	describeNotebookOutput, err := sagemaker.client.DescribeNotebookInstance(i)

	notebookInstance := NotebookInstance{
		AcceleratorTypes:                    aws.StringValueSlice(describeNotebookOutput.AcceleratorTypes),
		AdditionalCodeRepositories:          aws.StringValueSlice(describeNotebookOutput.AdditionalCodeRepositories),
		CreationTime:                        aws.TimeValue(describeNotebookOutput.CreationTime),
		DefaultCodeRepository:               aws.StringValue(describeNotebookOutput.DefaultCodeRepository),
		DirectInternetAccess:                aws.StringValue(describeNotebookOutput.DirectInternetAccess),
		FailureReason:                       aws.StringValue(describeNotebookOutput.FailureReason),
		InstanceType:                        aws.StringValue(describeNotebookOutput.InstanceType),
		KmsKeyId:                            aws.StringValue(describeNotebookOutput.KmsKeyId),
		LastModifiedTime:                    aws.TimeValue(describeNotebookOutput.LastModifiedTime),
		NetworkInterfaceId:                  aws.StringValue(describeNotebookOutput.NetworkInterfaceId),
		NotebookInstanceArn:                 aws.StringValue(describeNotebookOutput.NotebookInstanceArn),
		NotebookInstanceLifecycleConfigName: aws.StringValue(describeNotebookOutput.NotebookInstanceLifecycleConfigName),
		NotebookInstanceName:                aws.StringValue(describeNotebookOutput.NotebookInstanceName),
		NotebookInstanceStatus:              aws.StringValue(describeNotebookOutput.NotebookInstanceStatus),
		RoleArn:                             aws.StringValue(describeNotebookOutput.RoleArn),
		RootAccess:                          aws.StringValue(describeNotebookOutput.RootAccess),
		SecurityGroups:                      aws.StringValueSlice(describeNotebookOutput.SecurityGroups),
		SubnetId:                            aws.StringValue(describeNotebookOutput.SubnetId),
		Url:                                 aws.StringValue(describeNotebookOutput.Url),
		VolumeSizeInGB:                      aws.Int64Value(describeNotebookOutput.VolumeSizeInGB),
	}

	return notebookInstance, err
}
