// @@@SNIPSTART transfer-money-project-template-go-worker
package main

import (
  "log"

  "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

  "transfer-money-project-template-go/app"
)

func main() {
  c, err := client.NewClient(client.Options{})
  if err != nil {
    log.Fatalln("unable to create Temporal client", err)
  }
  defer c.Close()

  w := worker.New(c, app.TransferMoneyTaskQueue, worker.Options{})
  w.RegisterWorkflow(app.TransferMoney)
  w.RegisterActivity(app.Withdraw)
  w.RegisterActivity(app.Deposit)

  err = w.Run(worker.InterruptCh())
  if err != nil {
    log.Fatalln("unable to start Worker", err)
  }
}
// @@@SNIPEND
