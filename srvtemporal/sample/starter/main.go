package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/anhgeeky/labs-temporal-go/sample"
	"go.temporal.io/sdk/client"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.NewLazyClient(client.Options{
		HostPort:  "localhost:7233",
		Namespace: "staging",
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowID := "BANK_TRANSFER-" + fmt.Sprintf("%d", time.Now().Unix())

	workflowOptions := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "taskQueueSample",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, sample.Workflow, "Temporal")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	// Synchronously wait for the workflow completion.
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}
