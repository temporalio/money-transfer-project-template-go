package app

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"strings"

	"go.temporal.io/sdk/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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

// CreateClientOptionsFromEnv creates a client.Options instance, configures
// it based on environment variables, and returns that instance. It
// supports the following environment variables:
//
//		TEMPORAL_ADDRESS: Host and port (formatted as host:port) of the endpoint (Temporal frontend)
//		TEMPORAL_NAMESPACE: Namespace to be used by the Client
//	 TEMPORAL_CLOUD_API_KEY: The API key to use for authentication in Temporal Cloud
//		TEMPORAL_TLS_CERT: Path to the x509 certificate
//		TEMPORAL_TLS_KEY: Path to the private certificate key
//
// This uses the SDK's default configuration for a given setting if the
// corresponding environment variable is not set.
func CreateClientOptionsFromEnv() (client.Options, error) {
	hostPort := os.Getenv("TEMPORAL_ADDRESS")
	namespaceName := os.Getenv("TEMPORAL_NAMESPACE")

	clientOpts := client.Options{
		HostPort:  hostPort,
		Namespace: namespaceName,
	}

	if apiKey := os.Getenv("TEMPORAL_CLOUD_API_KEY"); apiKey != "" {
		// Warn if the environment variable for an API key is defined, but
		// but the endpoint address is not valid for API key authentication.
		// The detail page for the Namespace in Temporal Cloud will show
		// which endpoint to use with API keys (has a temporal.io domain)
		// and which endpoint to use with mTLS (has a tmprl.cloud domain).
		if !strings.Contains(hostPort, ".temporal.io:") {
			log.Println("warning: using an API key, but not an API key endpoint")
		}

		clientOpts.Credentials = client.NewAPIKeyStaticCredentials(apiKey)
		clientOpts.ConnectionOptions = client.ConnectionOptions{
			TLS: &tls.Config{},
			DialOptions: []grpc.DialOption{
				grpc.WithUnaryInterceptor(
					func(ctx context.Context, method string, req any, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
						return invoker(
							metadata.AppendToOutgoingContext(ctx, "temporal-namespace", namespaceName),
							method,
							req,
							reply,
							cc,
							opts...,
						)
					},
				),
			},
		}
	}

	if certPath := os.Getenv("TEMPORAL_TLS_CERT"); certPath != "" {
		cert, err := tls.LoadX509KeyPair(certPath, os.Getenv("TEMPORAL_TLS_KEY"))
		if err != nil {
			return clientOpts, fmt.Errorf("failed loading key pair: %w", err)
		}

		clientOpts.ConnectionOptions.TLS = &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
	}

	return clientOpts, nil
}
