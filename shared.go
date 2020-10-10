package app

// @@@SNIPSTART money-transfer-project-template-go-shared-task-queue
const TransferMoneyTaskQueue = "TRANSFER_MONEY_TASK_QUEUE"
// @@@SNIPEND

type TransferDetails struct {
	Amount      float32
	FromAccount string
	ToAccount   string
	ReferenceID string
}
