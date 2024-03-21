package app

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_SuccessfulTransferWorkflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	testDetails := PaymentDetails{
		SourceAccount: "85-150",
		TargetAccount: "43-812",
		Amount:        250,
		ReferenceID:   "12345",
	}

	// Mock activity implementation
	env.OnActivity(Withdraw, mock.Anything, testDetails).Return("", nil)
	env.OnActivity(Deposit, mock.Anything, testDetails).Return("", nil)

	env.ExecuteWorkflow(MoneyTransfer, testDetails)
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
}

func Test_DepositFailedWorkflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	testDetails := PaymentDetails{
		SourceAccount: "85-150",
		TargetAccount: "43-812",
		Amount:        250,
		ReferenceID:   "12345",
	}

	// Mock activity implementation
	env.OnActivity(Withdraw, mock.Anything, testDetails).Return("", nil)
	env.OnActivity(Deposit, mock.Anything, testDetails).Return("", errors.New("unable to deposit"))
	env.OnActivity(Refund, mock.Anything, testDetails).Return("", nil)

	env.ExecuteWorkflow(MoneyTransfer, testDetails)
	require.True(t, env.IsWorkflowCompleted())
	require.Error(t, env.GetWorkflowError())
}
