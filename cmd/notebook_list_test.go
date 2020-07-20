package cmd

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/garretthoffman/sagecli/cmd/mock"
	"github.com/garretthoffman/sagecli/sagemaker"
	sagemakerMockClient "github.com/garretthoffman/sagecli/sagemaker/mock/client"
)

func TestListNotebookInstanceOperation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := sagemakerMockClient.NewMockClient(mockCtrl)
	mockOutput := &mock.Output{}

	notebookInstance1 := sagemaker.NotebookInstance{
		NotebookInstanceName:   "notebook-test-1",
		NotebookInstanceStatus: "Stopped",
		InstanceType:           "ml.t2.medium",
		CreationTime:           time.Now(),
		Url:                    "notebook-test-1.notebook.us-east-1.sagemaker.aws",
	}

	notebookInstance2 := sagemaker.NotebookInstance{
		NotebookInstanceName:   "notebook-test-2",
		NotebookInstanceStatus: "Running",
		InstanceType:           "ml.m5.medium",
		CreationTime:           time.Now(),
		Url:                    "notebook-test-2.notebook.us-east-1.sagemaker.aws",
	}

	notebookInstances := sagemaker.NotebookInstances{notebookInstance1, notebookInstance2}

	mockClient.EXPECT().ListNotebookInstances().Return(notebookInstances, nil)

	notebookListOperation{
		sagemaker: mockClient,
		output:    mockOutput,
	}.execute()

	if len(mockOutput.Tables) == 0 {
		t.Fatalf("expected table, got none")
	}

	if len(mockOutput.Tables[0].Rows) != 3 {
		t.Errorf("expected table with 3 rows, got %d", len(mockOutput.Tables[0].Rows))
	}

	if expected, got := []string{"NAME", "STATUS", "INSTANCE TYPE", "CREATED AT", "URL"}, mockOutput.Tables[0].Rows[0]; !reflect.DeepEqual(expected, got) {
		t.Errorf("expected column headers: %v, got: %v", expected, got)
	}

	row1 := mockOutput.Tables[0].Rows[1]

	if row1[0] != notebookInstance1.NotebookInstanceName {
		t.Errorf("expected name: %s, got: %s", notebookInstance1.NotebookInstanceName, row1[0])
	}

	if row1[1] != notebookInstance1.NotebookInstanceStatus {
		t.Errorf("expected name: %s, got: %s", notebookInstance1.NotebookInstanceStatus, row1[1])
	}

	if row1[2] != notebookInstance1.InstanceType {
		t.Errorf("expected name: %s, got: %s", notebookInstance1.InstanceType, row1[2])
	}

	if expected := notebookInstance1.CreationTime.Format("2006-01-02 15:04:05"); row1[3] != expected {
		t.Errorf("expected name: %s, got: %s", expected, row1[3])
	}

	if row1[4] != notebookInstance1.Url {
		t.Errorf("expected name: %s, got: %s", notebookInstance1.Url, row1[4])
	}

	row2 := mockOutput.Tables[0].Rows[2]

	if row2[0] != notebookInstance2.NotebookInstanceName {
		t.Errorf("expected name: %s, got: %s", notebookInstance2.NotebookInstanceName, row2[0])
	}

	if row2[1] != notebookInstance2.NotebookInstanceStatus {
		t.Errorf("expected name: %s, got: %s", notebookInstance2.NotebookInstanceStatus, row2[1])
	}

	if row2[2] != notebookInstance2.InstanceType {
		t.Errorf("expected name: %s, got: %s", notebookInstance2.InstanceType, row2[2])
	}

	if expected := notebookInstance2.CreationTime.Format("2006-01-02 15:04:05"); row2[3] != expected {
		t.Errorf("expected name: %s, got: %s", expected, row2[3])
	}

	if row2[4] != notebookInstance2.Url {
		t.Errorf("expected name: %s, got: %s", notebookInstance2.Url, row2[4])
	}
}

func TestListNotebookInstanceOperationError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := sagemakerMockClient.NewMockClient(mockCtrl)
	mockOutput := &mock.Output{}

	mockClient.EXPECT().ListNotebookInstances().Return(sagemaker.NotebookInstances{}, errors.New("sdk error!"))

	notebookListOperation{
		sagemaker: mockClient,
		output:    mockOutput,
	}.execute()

	if len(mockOutput.FatalMsgs) == 0 {
		t.Fatalf("expected fatal output, got none")
	}

	if expected, got := "Could not list notebook instances", mockOutput.FatalMsgs[0].Msg; got != expected {
		t.Errorf("expected fatal output: %s, got: %s", expected, got)
	}
}

func TestListNotebookInstanceOperationNoneFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := sagemakerMockClient.NewMockClient(mockCtrl)
	mockOutput := &mock.Output{}

	mockClient.EXPECT().ListNotebookInstances().Return(sagemaker.NotebookInstances{}, nil)

	notebookListOperation{
		sagemaker: mockClient,
		output:    mockOutput,
	}.execute()

	if len(mockOutput.InfoMsgs) == 0 {
		t.Fatalf("expected info output, got none")
	}

	if expected, got := "No notebook instances found", mockOutput.InfoMsgs[0]; got != expected {
		t.Errorf("expected info output: %s, got: %s", expected, got)
	}
}
