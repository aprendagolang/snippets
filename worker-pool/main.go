package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id  int
	num int
}

type Result struct {
	job   Job
	total int
}

var (
	jobs    = make(chan Job, 10)
	results = make(chan Result, 10)
)

func sum(number int) (total int) {
	no := number
	for no != 0 {
		digit := no % 10
		total += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, sum(job.num)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocateJob() {
	for i := 0; i < 300; i++ {
		num := rand.Intn(999)
		job := Job{i, num}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, número %d , soma dos dígitos %d\n", result.job.id, result.job.num, result.total)
	}
	done <- true
}

func main() {
	go allocateJob()
	done := make(chan bool)
	go result(done)
	createWorkerPool()
	<-done
}
