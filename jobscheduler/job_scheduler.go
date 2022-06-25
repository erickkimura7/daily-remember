package jobscheduler

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	Message  string
	DateTime time.Time
}

func PoolJob() {

	pool := []*Job{}

	pool = append(pool, &Job{
		Message:  "Teste",
		DateTime: time.Now().Add(time.Duration(-1) * time.Minute),
	})

	var wg sync.WaitGroup

	wg.Add(len(pool))

	for _, each := range pool {
		go func(job *Job) {
			difference := time.Since(job.DateTime)
			time.Sleep(difference)
			fmt.Println(job.Message)
			wg.Done()
		}(each)
	}

	wg.Wait()
}

func StartJob() {

}
