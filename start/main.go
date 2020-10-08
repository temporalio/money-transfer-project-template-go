// @@@SNIPSTART transfer-money-project-template-go-start-workflow
package main

import (
  "context"
  "log"

  "github.com/google/uuid"
  "go.temporal.io/sdk/client"

  "transfer-money-project-template-go/app"
)

func main(){
  c, err := client.NewClient(client.Options{})
  if err != nil {
    log.Fatalln("unable to create Temporal client", err)
  }
  defer c.Close()

  options := client.StartWorkflowOptions{
    ID: "transfer-money-workflow",
    TaskQueue: app.TransferMoneyTaskQueue,
  }

  transferDetails := app.TransferDetails{
    Amount: 54.99,
    FromAccount: "001-001",
    ToAccount: "002-002",
    ReferenceID: uuid.New().String(),
  }

  we, _ := c.ExecuteWorkflow(context.Background(), options, app.TransferMoney, transferDetails)

  log.Println("Started workflow", "WorkflowID:", we.GetID(), "RunID:", we.GetRunID())
}
// @@@SNIPEND
