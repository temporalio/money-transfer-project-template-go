package app

import (
	"log"
	"os"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/contrib/envconfig"
)

// @@@SNIPSTART money-transfer-project-template-go-shared-task-queue
const MoneyTransferTaskQueueName = "TRANSFER_MONEY_TASK_QUEUE"

// @@@SNIPEND

// @@@SNIPSTART money-transfer-project-template-go-transferdetails
type PaymentDetails struct {
	SourceAccount string
	TargetAccount string
	Amount        int
	ReferenceID   string
}

// @@@SNIPEND

// CreateClientOptionsFromEnv creates and returns an instance of
// client.Options instance, based on two environment variables:
//
//	TEMPORAL_CONFIG_PATH: Path to the TOML file that defines the profile
//	TEMPORAL_PROFILE_NAME: Name of a specific profile in the TOML file to use
func CreateClientOptionsFromEnv() (client.Options, error) {
	clientOpts := client.Options{}

	configFilePath := os.Getenv("TEMPORAL_CONFIG_PATH")
	profileName := os.Getenv("TEMPORAL_PROFILE_NAME")
	if configFilePath != "" && profileName != "" {
		var err error
		clientOpts, err = envconfig.LoadClientOptions(envconfig.LoadClientOptionsRequest{
			ConfigFilePath:    configFilePath,
			ConfigFileProfile: profileName,
		})
		if err != nil {
			log.Fatalf("failed to load profile: %v", err)
		}
	}

	return clientOpts, nil
}
