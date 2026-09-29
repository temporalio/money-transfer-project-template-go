package app

import "os"

// @@@SNIPSTART money-transfer-project-template-go-shared-task-queue
// MoneyTransferTaskQueueName is the Task Queue the Worker polls and the Workflow is
// started on. Reads TEMPORAL_TASK_QUEUE when set — e.g. a test harness isolating each
// run on its own queue — otherwise the shared default, so a copy-paste user is unaffected.
var MoneyTransferTaskQueueName = EnvOr("TEMPORAL_TASK_QUEUE", "TRANSFER_MONEY_TASK_QUEUE")

// @@@SNIPEND

// EnvOr returns the value of environment variable key, or fallback when it is unset
// or empty. Used to keep the default behavior while letting a test harness override.
func EnvOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// @@@SNIPSTART money-transfer-project-template-go-transferdetails
type PaymentDetails struct {
	SourceAccount string
	TargetAccount string
	Amount        int
	ReferenceID   string
}

// @@@SNIPEND
