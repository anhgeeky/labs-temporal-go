package main

import (
	"context"
	"log"

	srvtemporal "github.com/anhgeeky/labs-temporal-go"
	"github.com/anhgeeky/labs-temporal-go/sample"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

type NewRegisterer struct {
}

func (r NewRegisterer) Register(register worker.Registry) {
	register.RegisterWorkflow(sample.Workflow)
	register.RegisterActivity(sample.Activity)
}

func main() {
	ctx := context.Background()

	c, err := client.NewLazyClient(client.Options{
		HostPort:  "localhost:7233",
		Namespace: "staging",
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	log.Println("Temporal client connected")
	defer c.Close()

	cfg := srvtemporal.PlatformConfig{}
	var f srvtemporal.Registerer = NewRegisterer{}
	w, _ := srvtemporal.NewWorker(f, cfg,
		srvtemporal.WithClient(c),
		srvtemporal.WithName("name"),
		srvtemporal.WithTaskQueue("taskQueueSample"),
	)
	w.Run(ctx)
}
