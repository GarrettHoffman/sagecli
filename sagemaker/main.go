// Package sagemaker is a client for AWS Sagemaker.
package sagemaker


//go:generate mockgen -package client -destination=mock/client/client.go github.com/garretthoffman/sagecli/sagemaker Client
//go:generate mockgen -package sdk -source ../vendor/github.com/aws/aws-sdk-go/service/sagemaker/sagemakeriface/interface.go -destination=mock/sdk/sagemakeriface.go github.com/aws/aws-sdk-go/service/sagemaker/sagemakeriface SageMakerAPI

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sagemaker"
	"github.com/aws/aws-sdk-go/service/sagemaker/sagemakeriface"
)

// Client represnts a method for accessing Sagemaker.
type Client interface {
	ListNotebookInstances() (NotebookInstances, error)
}

// SDKClient implements access to Sagemaker via the AWS SDK.
type SDKClient struct {
	client sagemakeriface.SageMakerAPI
}

// New returns an SDKClient configured with the given session.
func New(sess *session.Session) SDKClient {
	return SDKClient{
		client: sagemaker.New(sess),
	}
}