package main

import (
	"fmt"
	"time"
)

type ControlMsg int

const (
	DoExit = iota
	ExitOk
)

type Job struct {
	data int
}

type Result struct {
	result int
	job    Job
}

// job is a "read only", or "receiver" channel
// results is a "write only" channel
// control is bidirectional
func doubler(jobs <-chan Job, results chan<- Result, control chan ControlMsg) {
	for {
		select {
		case msg := <-control:
			switch msg {
			case DoExit:
				fmt.Println("Exit go routine")
				control <- ExitOk

			default:
				panic("Unhandled control message")
			}
		case job := <-jobs:
			results <- Result{result: job.data * 2, job: job}
		}
	}
}

func main() {

	// jobs channel
	jobs := make(chan Job, 50)
	// results channel
	results := make(chan Result, 50)
	// control channel
	control := make(chan ControlMsg)

	go doubler(jobs, results, control)

	for i := 0; i < 30; i++ {
		jobs <- Job{i}
	}

	for {
		select {
		case result := <-results:
			fmt.Printf("Job: %d, Result: %d\n", result.job.data, result.result)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Timeout")
			control <- DoExit
			<-control
			fmt.Println("Exit main")
			return
		}
	}
}
