package app

import (
	"context"
	"log"
)

// @@@SNIPSTART money-transfer-project-template-go-activity-withdraw
func Withdraw(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf(
		"Withdrawing $%d from account %s.\n\n",
		data.Amount,
		data.SourceAccount,
	)
	bank := BankingService{"bank-api.example.com"}
	confirmation, err := bank.Withdraw(data.SourceAccount, data.Amount)
	return confirmation, err
}

// @@@SNIPEND

// @@@SNIPSTART money-transfer-project-template-go-activity-deposit
func Deposit(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf(
		"Depositing $%d into account %s.\n\n",
		data.Amount,
		data.TargetAccount,
	)
	bank := BankingService{"bank-api.example.com"}
	// Uncomment the next line and comment the one after that to simulate failure
	// confirmation, err := bank.DepositThatFails(data.TargetAccount, data.Amount)
	confirmation, err := bank.Deposit(data.TargetAccount, data.Amount)
	return confirmation, err
}

// @@@SNIPEND
