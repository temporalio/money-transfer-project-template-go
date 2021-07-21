# Temporal Go Project Template

This is a simple project for demonstrating Temporal with the Go SDK.

The full 20 minute guide is here: https://docs.temporal.io/docs/go/run-your-first-app-tutorial

## Basic instructions

### Step 0: Temporal Server

Make sure [Temporal Server is running](https://docs.temporal.io/docs/server/quick-install/) first:

```bash
git clone https://github.com/temporalio/docker-compose.git
cd  docker-compose
docker-compose up
```

Leave it running. You can see Temporal Web at [localhost:8088](localhost:8088). There should be no workflows visible in the dashboard right now.

### Step 1: Clone this repo

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

Now you can see the workflow run to completion. You can also see the worker polling for workflows and activities in the task queue at [http://localhost:8088/namespaces/default/task-queues/TRANSFER_MONEY_TASK_QUEUE](http://localhost:8088/namespaces/default/task-queues/TRANSFER_MONEY_TASK_QUEUE).

<img width="882" alt="CleanShot 2021-07-20 at 17 48 45@2x" src="https://user-images.githubusercontent.com/6764957/126413160-18663430-bb7a-4d3a-874e-80598e1fa07d.png">

## What Next?

You can run the Workflow code a few more times with `go run start/main.go` to understand how it interacts with the Worker and Temporal Server.

Please [read the tutorial](https://docs.temporal.io/docs/go/run-your-first-app-tutorial) for more details.
