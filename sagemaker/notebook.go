package sagemaker

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	awsSagemaker "github.com/aws/aws-sdk-go/service/sagemaker"
	"github.com/garretthoffman/sagecli/console"
)

// Notebook represents a Sagemaker notebook instance.
type Notebook struct {
	ARN						string
	Name					string
	Status					string
	Url						string
	InstanceType			string
	RoleARN					string
	CreatedAt				float32
	ModifiedAt				float32
	VolumeSize				int16
	RootAccess				string
	DirectInternetAccess	string
	SubnetId				string
	SecurityGroups			[]string
	AcceleratorTypes		[]string
	CodeRepo				string
	AdditionalCodeRepos		[]string
}

// Notebooks is a collection of Sagemaker notebook instances.
type Notebooks []Notebook