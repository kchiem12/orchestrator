package worker

import (
	"fmt"

	"github.com/kchiem12/orchestrator/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]task.Task
	Taskcount int
}

func (w *Worker) CollectStats() {
	fmt.Println("Collecting stats")
}

func (w *Worker) RunTask() {
	fmt.Println("Running task")
}

func (w *Worker) StartTask() {
	fmt.Println("Starting task")
}

func (w *Worker) StopTask() {
	fmt.Println("Stopping task")
}
