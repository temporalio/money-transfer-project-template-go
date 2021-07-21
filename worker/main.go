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
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	// This worker hosts both Workflow and Activity functions
	w := worker.New(c, app.TransferMoneyTaskQueue, worker.Options{})
	w.RegisterWorkflow(app.TransferMoney)
	w.RegisterActivity(app.Withdraw)
	w.RegisterActivity(app.Deposit)
	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
// @@@SNIPEND
