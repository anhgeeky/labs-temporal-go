package main

import (
	"context"

	srvtemporal "github.com/jamillosantos/server-temporal"
	"github.com/jamillosantos/server-temporal/sample"
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
	cfg := srvtemporal.PlatformConfig{}
	var f srvtemporal.Registerer = NewRegisterer{}
	w, _ := srvtemporal.NewWorker(f, cfg, srvtemporal.WithName("name"))
	w.Listen(ctx)
}
