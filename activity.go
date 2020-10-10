package app

import (
	"context"
	"fmt"
)

func Withdraw(ctx context.Context, transferDetails TransferDetails) error {
	fmt.Printf(
		"\nWithdrawing $%f from account %s. ReferenceId: %s\n",
		transferDetails.Amount,
		transferDetails.FromAccount,
		transferDetails.ReferenceID,
	)
	return nil
}

// @@@SNIPSTART money-transfer-project-template-go-activity
func Deposit(ctx context.Context, transferDetails TransferDetails) error {
	fmt.Printf(
		"\nDepositing $%f into account %s. ReferenceId: %s\n",
		transferDetails.Amount,
		transferDetails.ToAccount,
		transferDetails.ReferenceID,
	)
	// Switch out comments on the return statements to simulate an error
	//return fmt.Errorf("deposit did not occur due to an issue")
	return nil
}
// @@@SNIPEND"
