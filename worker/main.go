package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"money-transfer-project-template-go/app"
)

// @@@SNIPSTART money-transfer-project-template-go-worker
func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, app.MoneyTransferTaskQueueName, worker.Options{})

	// This worker hosts both Workflow and Activity functions
	w.RegisterWorkflow(app.MoneyTransfer)
	w.RegisterActivity(app.Withdraw)
	w.RegisterActivity(app.Deposit)

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}

// @@@SNIPEND
