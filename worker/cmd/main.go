package main

import (
	"context"
	"fmt"
	openapi "github.com/care0717/digdag-worker"
	"time"
)

func main() {
	cfg := openapi.NewConfiguration()
	c := openapi.NewAPIClient(cfg)

	ctx := context.Background()
	maxTime := 2*60*60
	var workers []openapi.TaskWorker
	for i := 0; i < maxTime; i++ {
		var currentExecutingPoint int
		created := openapi.ToCreatedFormat(time.Duration(i)*time.Second)
		req := c.DefaultApi.JobsGet(ctx).Created(created)
		jobs, _, err := c.DefaultApi.JobsGetExecute(req)
		if err != nil {
			panic(err)
		}
		for _, j := range jobs {
			workers = append(workers, openapi.NewTaskWorker(j.Tasks))
		}
		var nextWorkers []openapi.TaskWorker
		for _, w := range workers {
			currentExecutingPoint += int(w.ExecutingPoint())
			w.Work()
			if !w.Complete() {
				nextWorkers = append(nextWorkers, w)
			}
		}
		workers = nextWorkers
		fmt.Println(currentExecutingPoint)
	}
}
