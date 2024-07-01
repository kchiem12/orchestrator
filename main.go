package main

import (
	"fmt"
	"time"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/kchiem12/orchestrator/manager"
	"github.com/kchiem12/orchestrator/node"
	"github.com/kchiem12/orchestrator/task"
	"github.com/kchiem12/orchestrator/worker"
)

func main() {
	t := task.Task{
		ID:     uuid.New(),
		Name:   "task1",
		State:  task.Pending,
		Image:  "nginx",
		Memory: 1024,
		Disk:   1,
	}

	te := task.TaskEvent{
		ID:        uuid.New(),
		State:     task.Pending,
		Timestamp: time.Now(),
		Task:      t,
	}

	fmt.Printf("task: %v\n", t)
	fmt.Printf("task event: %v\n", te)

	w := worker.Worker{
		Queue: *queue.New(),
		Db:    make(map[uuid.UUID]task.Task),
	}

	fmt.Printf("worker: %v\n", w)
	w.CollectStats()
	w.RunTask()
	w.StartTask()
	w.StopTask()

	m := manager.Manager{
		Pending: *queue.New(),
		TaskDb:  make(map[string][]task.Task),
		EventDb: make(map[string][]task.TaskEvent),
		Workers: []string{w.Name},
	}

	fmt.Printf("manager: %v\n", m)
	m.SelectWorker()
	m.UpdateTasks()
	m.SendWork()

	n := node.Node{
		Name:            "node1",
		Ip:              "127.0.0.0",
		Cores:           4,
		Memory:          1024,
		MemoryAllocated: 512,
		Disk:            10,
		DiskAllocated:   5,
		Role:            "worker",
	}

	fmt.Printf("node: %v\n", n)
}
