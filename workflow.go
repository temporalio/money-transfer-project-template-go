package app

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// @@@SNIPSTART money-transfer-project-template-go-workflow
func MoneyTransfer(ctx workflow.Context, input PaymentDetails) (string, error) {
	// RetryPolicy specifies how to automatically handle retries if an Activity fails.
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    time.Minute,
		MaximumAttempts:    500,
	}
	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failures by default, this is just an example.
		RetryPolicy: retrypolicy,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var output1 string
	err := workflow.ExecuteActivity(ctx, Withdraw, input).Get(ctx, &output1)
	if err != nil {
		return "", err
	}

	var output2 string
	err = workflow.ExecuteActivity(ctx, Deposit, input).Get(ctx, &output2)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Transfer complete (transaction IDs: %s, %s)", output1, output2)
	return result, nil
}

// @@@SNIPEND
