package app

const MoneyTransferTaskQueueName = "TRANSFER_MONEY_TASK_QUEUE"

type PaymentDetails struct {
	SourceAccount string
	TargetAccount string
	Amount        int
	ReferenceID   string
}
