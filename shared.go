package app

import (
	"fmt"
	"os"
	"path/filepath"

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
// client.Options. This uses the default settings, unless the
// TEMPORAL_PROFILE environment variable is set, in which case
// it configures the options as per the specified profile name.
func CreateClientOptionsFromEnv() (client.Options, error) {
	profileName := os.Getenv("TEMPORAL_PROFILE")
	if profileName == "" {
		return client.Options{}, nil
	}

	configFilePath, err := DefaultConfigFilePath()
	if err != nil {
		return client.Options{}, fmt.Errorf("failed to get config path: %w", err)
	}

	clientOpts, err := envconfig.LoadClientOptions(envconfig.LoadClientOptionsRequest{
		ConfigFilePath:    configFilePath,
		ConfigFileProfile: profileName,
	})
	if err != nil {
		return client.Options{}, fmt.Errorf("failed to load profile %q: %w", profileName, err)
	}

	return clientOpts, nil
}

// Returns the path representing the default location of the
// configuration file, based on the current operating system.
func DefaultConfigFilePath() (string, error) {
	userDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("failed getting user config dir: %w", err)
	}
	return filepath.Join(userDir, "temporalio", "temporal.toml"), nil
}
