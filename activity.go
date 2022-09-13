package app

import (
	"context"
)

// @@@SNIPSTART money-transfer-project-template-go-activity-withdraw
func Withdraw(ctx context.Context, data PaymentDetails) (string, error) {
	bank := BankingService{"bank-api.example.com"}
	confirmation, err := bank.Withdraw(data.SourceAccount, data.Amount)
	if err != nil {
		return "", err
	}

	return confirmation, nil
}

// @@@SNIPEND

// @@@SNIPSTART money-transfer-project-template-go-activity-deposit
func Deposit(ctx context.Context, data PaymentDetails) (string, error) {
	bank := BankingService{"bank-api.example.com"}
	// Uncomment the next line and comment the one after that to simulate failure
	// confirmation, err := bank.DepositThatFails(data.TargetAccount, data.Amount)
	confirmation, err := bank.Deposit(data.TargetAccount, data.Amount)
	if err != nil {
		return "", err
	}

	return confirmation, nil
}

// @@@SNIPEND
