package main

import (
	"context"
	"log"
	"os"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/contrib/envconfig"

	"money-transfer-project-template-go/app"
)

// @@@SNIPSTART money-transfer-project-template-go-start-workflow
func main() {
	// Create the client object just once per process.
	// Connect to Temporal Cloud by loading the "cloud-setup" profile from the
	// shared Temporal client config (temporal.toml), which supplies the Cloud
	// address, namespace, TLS settings, and API key.
	opts, err := envconfig.LoadClientOptions(envconfig.LoadClientOptionsRequest{
		ConfigFileProfile: "cloud-setup",
	})
	if err != nil {
		log.Fatalln("Unable to load Temporal client options:", err)
	}

	c, err := client.Dial(opts)

	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}

	defer c.Close()

	input := app.PaymentDetails{
		SourceAccount: "85-150",
		TargetAccount: "43-812",
		Amount:        250,
		ReferenceID:   "12345",
	}

	// Use the WORKFLOW_ID from the environment if set (the cloud-setup flow uses
	// distinct names for the clean run and the failure-and-recovery run), and
	// otherwise fall back to the default demo name.
	workflowID := os.Getenv("WORKFLOW_ID")
	if workflowID == "" {
		workflowID = "money-transfer-demo"
	}

	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: app.MoneyTransferTaskQueueName,
	}

	log.Printf("Starting transfer from account %s to account %s for %d", input.SourceAccount, input.TargetAccount, input.Amount)

	we, err := c.ExecuteWorkflow(context.Background(), options, app.MoneyTransfer, input)
	if err != nil {
		log.Fatalln("Unable to start the Workflow:", err)
	}

	log.Printf("WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())

	var result string

	err = we.Get(context.Background(), &result)

	if err != nil {
		log.Fatalln("Unable to get Workflow result:", err)
	}

	log.Println(result)
}

// @@@SNIPEND
