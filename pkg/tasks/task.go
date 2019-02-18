package tasks

import (
	"log"
	"os/exec"
	"sync"
)

// Task struct stores the properties of a task
type Task struct {
	name string
	done bool
	cmd  []string
}

// Create creates and returns a Task
func Create(name string, cmd ...string) Task {
	return Task{
		name: name,
		done: false,
		cmd:  cmd,
	}
}

// Execute executes a process
func Execute(wg *sync.WaitGroup, name string, args ...string) {
	cmd := exec.Command(name, args...)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	defer wg.Done()
}

// Start begins executing the task
func (task *Task) Start() {
	var wg sync.WaitGroup

	wg.Add(1)
	go Execute(&wg, task.cmd[0], task.cmd[1:]...)
	wg.Wait()

	task.Complete()
}

// Complete sets the Task done
func (task *Task) Complete() {
	task.done = true
}
