package main

import (
	"crypto/tls"
	"log"
	"os"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"money-transfer-project-template-go/app"
)

// @@@SNIPSTART money-transfer-project-template-go-worker
func main() {

	// Connect to Temporal Cloud using the gRPC endpoint, namespace, and API key
	// supplied via environment variables, with TLS enabled.
	c, err := client.Dial(client.Options{
		HostPort:    os.Getenv("TEMPORAL_ADDRESS"),
		Namespace:   os.Getenv("TEMPORAL_NAMESPACE"),
		Credentials: client.NewAPIKeyStaticCredentials(os.Getenv("TEMPORAL_API_KEY")),
		ConnectionOptions: client.ConnectionOptions{
			TLS: &tls.Config{},
		},
	})
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
