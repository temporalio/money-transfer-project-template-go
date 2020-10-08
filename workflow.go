// @@@SNIPSTART transfer-money-project-template-go-workflow
package app

import (
  "fmt"
  "time"

  "go.temporal.io/sdk/workflow"
)

type TransferDetails struct {
  Amount      float32
  FromAccount string
  ToAccount   string
  ReferenceID string
}

func TransferMoney(ctx workflow.Context, transferDetails TransferDetails) error {

  options := workflow.ActivityOptions {
    ScheduleToCloseTimeout: time.Minute,
  }
  ctx = workflow.WithActivityOptions(ctx, options)

  err := workflow.ExecuteActivity(ctx, Withdraw, transferDetails).Get(ctx, nil)
  if err != nil {
    fmt.Println(err.Error())
  }

  err = workflow.ExecuteActivity(ctx, Deposit, transferDetails).Get(ctx, nil)
  if err != nil {
    fmt.Println(err.Error())
  }

  return nil
}
// @@@SNIPEND
