package main

import (
	"context"
	"flag"
	"fmt"
	openapi "github.com/care0717/digdag-worker"
	"time"
)

func main() {
	var (
		maxTime  int
		capacity int
	)
	flag.IntVar(&maxTime, "workMaxTime", 4800, "work max time [sec]")
	flag.IntVar(&capacity, "capacity", 15, "worker capacity")
	flag.Parse()
	cfg := openapi.NewConfiguration()
	c := openapi.NewAPIClient(cfg)

	ctx := context.Background()
	manager := openapi.NewWorkerManager(capacity)
	for i := 0; i < maxTime; i++ {
		created := openapi.ToCreatedFormat(time.Duration(i) * time.Second)
		req := c.DefaultApi.JobsGet(ctx).Created(created)
		jobs, _, err := c.DefaultApi.JobsGetExecute(req)
		if err != nil {
			panic(err)
		}
		for _, j := range jobs {
			worker := openapi.NewTaskWorker(j.Priority, j.Tasks)
			manager.Add(worker)
		}
		point := manager.Run()
		fmt.Println(point)
	}
}
