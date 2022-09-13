package app

// This code simulates a client for a hypothetical banking service.
// It supports both withdrawals and deposits, and generates a
// pseudorandom transaction ID for each request. Tip: You can
// modify these functions to introduce delays or errors, allowing
// you to experiment with failures and timeouts.
import (
	"errors"
	"math/rand"
)

type BankingService struct {
	// the hostname is to make it more realistic. This code does not
	// actually make any network calls.
	Hostname string
}

func (client BankingService) Withdraw(accountNum string, amount int) (string, error) {
	return generateTranactionId("W", 10), nil
}

func (client BankingService) Deposit(accountNum string, amount int) (string, error) {
	return generateTranactionId("D", 10), nil
}

func (client BankingService) DepositThatFails(accountNum string, amount int) (string, error) {
	return "", errors.New("This deposit has failed.")
}

func generateTranactionId(prefix string, length int) string {
	randChars := make([]byte, length)
	for i := range randChars {
		allowedChars := "0123456789"
		randChars[i] = allowedChars[rand.Intn(len(allowedChars))]
	}
	return prefix + string(randChars)
}
