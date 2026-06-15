package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/contrib/envconfig"
	"go.temporal.io/sdk/worker"

	"money-transfer-project-template-go/app"
)

// @@@SNIPSTART money-transfer-project-template-go-worker
func main() {

	// Connect to Temporal Cloud by loading the "cloud-setup" profile from the
	// shared Temporal client config (temporal.toml), which supplies the Cloud
	// address, namespace, TLS settings, and API key.
	opts, err := envconfig.LoadClientOptions(envconfig.LoadClientOptionsRequest{
		ConfigFileProfile: "cloud-setup",
	})
	if err != nil {
		log.Fatalln("Unable to load Temporal client options.", err)
	}

	c, err := client.Dial(opts)
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	defer c.Close()

	w := worker.New(c, app.MoneyTransferTaskQueueName, worker.Options{})

	// This worker hosts both Workflow and Activity functions.
	w.RegisterWorkflow(app.MoneyTransfer)
	w.RegisterActivity(app.Withdraw)
	w.RegisterActivity(app.Deposit)
	w.RegisterActivity(app.Refund)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}

// @@@SNIPEND
