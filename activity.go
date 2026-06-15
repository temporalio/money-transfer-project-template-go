package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/temporal"
)

// @@@SNIPSTART money-transfer-project-template-go-activity-withdraw
func Withdraw(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf("Withdrawing $%d from account %s.\n\n",
		data.Amount,
		data.SourceAccount,
	)

	referenceID := fmt.Sprintf("%s-withdrawal", data.ReferenceID)
	bank := BankingService{"bank-api.example.com"}
	confirmation, err := bank.Withdraw(data.SourceAccount, data.Amount, referenceID)
	return confirmation, err
}

// @@@SNIPEND

// @@@SNIPSTART money-transfer-project-template-go-activity-deposit
func Deposit(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf("Depositing $%d into account %s.\n\n",
		data.Amount,
		data.TargetAccount,
	)

	referenceID := fmt.Sprintf("%s-deposit", data.ReferenceID)
	bank := BankingService{"bank-api.example.com"}

	// Demo-only failure injection, driven by the DEMO_FAILURE env var on the
	// Worker. Unset/"off" leaves behavior unchanged.
	switch strings.ToLower(os.Getenv("DEMO_FAILURE")) {
	case "transient":
		// Reuse the always-failing banking path for the first two attempts; the
		// error is retryable, so Temporal retries and the activity succeeds on
		// attempt 3 -> the Workflow recovers and COMPLETEs.
		if activity.GetInfo(ctx).Attempt < 3 {
			return bank.DepositThatFails(data.TargetAccount, data.Amount, referenceID)
		}
	case "permanent":
		// Reuse the always-failing banking path, but make it non-retryable so the
		// Workflow's refund compensation (saga rollback) runs instead of retrying.
		_, err := bank.DepositThatFails(data.TargetAccount, data.Amount, referenceID)
		return "", temporal.NewNonRetryableApplicationError("deposit failed", "DepositFailure", err)
	}

	confirmation, err := bank.Deposit(data.TargetAccount, data.Amount, referenceID)
	return confirmation, err
}

// @@@SNIPEND

// @@@SNIPSTART money-transfer-project-template-go-activity-refund
func Refund(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf("Refunding $%v back into account %v.\n\n",
		data.Amount,
		data.SourceAccount,
	)

	referenceID := fmt.Sprintf("%s-refund", data.ReferenceID)
	bank := BankingService{"bank-api.example.com"}
	confirmation, err := bank.Deposit(data.SourceAccount, data.Amount, referenceID)
	return confirmation, err
}

// @@@SNIPEND
