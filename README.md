# Temporal Go Project Template

This is a simple project for demonstrating Temporal with the Go SDK.

The full 10-minute tutorial is here: https://learn.temporal.io/getting_started/go/first_program_in_go/

## Basic instructions

### Step 0: Temporal Server

Make sure [Temporal Server is running](https://docs.temporal.io/docs/server/quick-install/) first:

```bash
git clone https://github.com/temporalio/docker-compose.git
cd  docker-compose
docker-compose up
```

### Step 1: Clone this Repository

In another terminal instance, clone this repo and run this application.

```bash
git clone https://github.com/temporalio/money-transfer-project-template-go
cd money-transfer-project-template-go
```

### Step 2: Run the Workflow

```bash
go run start/main.go
```

Observe that Temporal Web reflects the workflow, but it is still in "Running" status. This is because there is no Workflow or Activity Worker yet listening to the `TRANSFER_MONEY_TASK_QUEUE` task queue to process this work.

### Step 3: Run the Worker

In YET ANOTHER terminal instance, run the worker. Notice that this worker hosts both Workflow and Activity functions.

```bash
go run worker/main.go
```

Now you can see the workflow run to completion. You can also see the worker polling for workflows and activities in the task queue at [http://localhost:8080/namespaces/default/task-queues/TRANSFER_MONEY_TASK_QUEUE](http://localhost:8080/namespaces/default/task-queues/TRANSFER_MONEY_TASK_QUEUE).

## What Next?

You can run the Workflow code a few more times with `go run start/main.go` to understand how it interacts with the Worker and Temporal Server.

Please [read the tutorial](https://learn.temporal.io/getting_started/go/first_program_in_go/) for more details.
