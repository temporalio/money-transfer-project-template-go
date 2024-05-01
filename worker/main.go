package main

// @@@SNIPSTART money-transfer-project-template-go-worker-cloud-import
import (
	"crypto/tls"
	"log"
	"os"

	"github.com/joho/godotenv"
	// @@@SNIPEND
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"money-transfer-project-template-go/app"
)

func main() {
	// @@@SNIPSTART money-transfer-project-template-go-worker-env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Unable to load environment variables from file", err)
	}
	// @@@SNIPEND
	// @@@SNIPSTART money-transfer-project-template-go-worker-key-pair
	clientKeyPath := os.Getenv("TEMPORAL_MTLS_TLS_KEY")
	clientCertPath := os.Getenv("TEMPORAL_MTLS_TLS_CERT")
	cert, err := tls.LoadX509KeyPair(clientCertPath, clientKeyPath)
	if err != nil {
		log.Fatalln("Unable to load cert and key pair.", err)
	}
	// @@@SNIPEND
	// @@@SNIPSTART money-transfer-project-template-go-worker-options
	namespace := os.Getenv("TEMPORAL_NAMESPACE")
	hostPort := os.Getenv("TEMPORAL_HOST_URL")
	c, err := client.Dial(client.Options{
		HostPort:  hostPort,
		Namespace: namespace,
		ConnectionOptions: client.ConnectionOptions{
			TLS: &tls.Config{Certificates: []tls.Certificate{cert}},
		},
	})
	// @@@SNIPEND
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
