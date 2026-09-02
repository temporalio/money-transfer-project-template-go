package app

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/contrib/envconfig"
)

// LoadClientOptions loads Temporal client configuration from the standard
// configuration file and environment variables.
func LoadClientOptions() (client.Options, error) {
	options, err := envconfig.LoadDefaultClientOptions()
	if err != nil {
		return client.Options{}, err
	}
	if options.HostPort == "" {
		options.HostPort = client.DefaultHostPort
	}
	if options.Namespace == "" {
		options.Namespace = client.DefaultNamespace
	}
	return options, nil
}
