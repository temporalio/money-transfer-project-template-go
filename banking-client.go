package app

// This code simulates a client for a hypothetical banking service.
// It supports both withdrawals and deposits, and generates a
// pseudorandom transaction ID for each request.
//
// Tip: You can modify these functions to introduce delays or errors, allowing
// you to experiment with failures and timeouts.
import (
	"errors"
	"math/rand"
)

type account struct {
	AccountNumber string
	Balance       int64
}

type bank struct {
	Accounts []account
}

func (b bank) findAccount(accountNumber string) (account, error) {

	for _, v := range b.Accounts {
		if v.AccountNumber == accountNumber {
			return v, nil
		}
	}

	return account{}, errors.New("account not found")
}

// InsufficientFundsError is raised when the account doesn't have enough money.
type InsufficientFundsError struct{}

func (m *InsufficientFundsError) Error() string {
	return "Insufficient Funds"
}

// InvalidAccountError is raised when the account number is invalid
type InvalidAccountError struct{}

func (m *InvalidAccountError) Error() string {
	return "Account number supplied is invalid"
}

// our mock bank
var mockBank = &bank{
	Accounts: []account{
		{AccountNumber: "85-150", Balance: 2000},
		{AccountNumber: "43-812", Balance: 0},
	},
}

// BankingService mocks interaction with a bank API. It supports withdrawals
// and deposits
type BankingService struct {
	// the hostname is to make it more realistic. This code does not
	// actually make any network calls.
	Hostname string
}

// Withdraw simulates a Withdrawal from a bank.
// Accepts the account number (string), amount (int), and a reference ID (string)
// for idempotent transaction tracking.
// Returns a transaction id when successful
// Returns various errors based on amount and account number.
func (client BankingService) Withdraw(accountNumber string, amount int, referenceID string) (string, error) {
	acct, err := mockBank.findAccount(accountNumber)

	if err != nil {
		return "", &InvalidAccountError{}
	}

	if amount > int(acct.Balance) {
		return "", &InsufficientFundsError{}
	}

	return generateTransactionID("W", 10), nil
}

// Deposit simulates a Deposit into a bank.
// Accepts the account number (string), amount (int), and a reference ID (string)
// for idempotent transaction tracking.
// Returns a transaction id when successful
// Returns InvalidAccountError if the account is invalid
func (client BankingService) Deposit(accountNumber string, amount int, referenceID string) (string, error) {

	_, err := mockBank.findAccount(accountNumber)
	if err != nil {
		return "", &InvalidAccountError{}
	}

	return generateTransactionID("D", 10), nil
}

// DepositThatFails simulates an unknown error.
func (client BankingService) DepositThatFails(accountNumber string, amount int, referenceID string) (string, error) {
	return "", errors.New("This deposit has failed.")
}

func generateTransactionID(prefix string, length int) string {
	randChars := make([]byte, length)
	for i := range randChars {
		allowedChars := "0123456789"
		randChars[i] = allowedChars[rand.Intn(len(allowedChars))]
	}
	return prefix + string(randChars)
}
