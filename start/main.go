package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"

	"money-transfer-project-template-go/app"
)

// @@@SNIPSTART money-transfer-project-template-go-start-workflow
func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	options := client.StartWorkflowOptions{
		ID:        "transfer-money-workflow",
		TaskQueue: app.TransferMoneyTaskQueue,
	}
	transferDetails := app.TransferDetails{
		Amount:      54.99,
		FromAccount: "001-001",
		ToAccount:   "002-002",
		ReferenceID: uuid.New().String(),
	}
	we, _ := c.ExecuteWorkflow(context.Background(), options, app.TransferMoney, transferDetails)
	printResults(transferDetails, we.GetID(), we.GetRunID())
}
// @@@SNIPEND

func printResults(transferDetails app.TransferDetails, workflowID, runID string) {
	fmt.Printf(
		"\nTransfer of $%f from account %s to account %s is processing. ReferenceID: %s\n",
		transferDetails.Amount,
		transferDetails.FromAccount,
		transferDetails.ToAccount,
		transferDetails.ReferenceID,
	)
	fmt.Printf(
		"\nWorkflowID: %s RunID: %s\n",
		workflowID,
		runID,
	)
}
