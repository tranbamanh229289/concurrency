package main

import (
	"fmt"
	"time"
)

type IJob interface {
	Process()
}

type Job struct {
	JobId int
	Name string
}

func (job *Job) Process() {

	fmt.Println("Do ", job.Name)

}

type Worker struct {
	Id         int
	JobRunning chan Job
	IsDone     chan bool
}

func NewWorker(id int, jobRunning chan Job) *Worker {
	return &Worker{
		Id:         id,
		JobRunning: jobRunning,
		IsDone:     make(chan bool),
	}
}

func (w *Worker) Run() {
	fmt.Println("running workerId: ", w.Id)
	go func (){
		for{
			select {
			case job:= <- w.JobRunning:
				fmt.Println("Job ", job.JobId,"(", job.Name,") is running by worker: ", w.Id)
				job.Process()
			case <- w.IsDone:
				fmt.Println("Job is complete! Stop worker ", w.Id)
				return
			}
		}
	}()
}

func (w *Worker) Stop(){
	w.IsDone <- true
}

type JobQueue struct {
	Workers []*Worker
	JobRunning chan Job
	IsDone chan bool
}

func NewJobQueue(numWorker int) *JobQueue {
	workers := make([]*Worker, numWorker, numWorker)
	jobRunning := make(chan Job)

	for i := 0; i < numWorker; i++ {
		workers[i] = NewWorker(i, jobRunning)
	}
	return &JobQueue{
		Workers: workers,
		JobRunning: jobRunning,
		IsDone: make(chan bool),
	}
}

func (jobQueue *JobQueue) Push(job Job) {
	jobQueue.JobRunning <- job
}

func (jobQueue *JobQueue) StopJobQueue() {
	jobQueue.IsDone <- true
}

func (jobQueue *JobQueue) Start() {
	go func() {
		for i := 0; i < len(jobQueue.Workers); i++ {
			jobQueue.Workers[i].Run()
			
		}
	}()

	go func(){
		for {
			select{
			case <- jobQueue.IsDone:
				for i := 0; i < len(jobQueue.Workers); i++ {
					jobQueue.Workers[i].Stop()
				}
				return
			}
		}
	}()
}



func main() {
	jobs := []string {
		"get",
		"post",
		"update",
		"delete",
		"authentication",
		"authorization",
	}

	jobQueue := NewJobQueue(4)
	jobQueue.Start()
	for index, value := range jobs {
		sender := Job{
			JobId: index,
			Name: value,
		}
		jobQueue.Push(sender)
	}

	time.AfterFunc(time.Second*2, func() {
		jobQueue.StopJobQueue()
	})

	time.Sleep(6*time.Second)
}