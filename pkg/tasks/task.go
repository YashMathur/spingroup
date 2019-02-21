package tasks

import (
	"log"
	"os/exec"
	"sync"
)

// Task struct stores the properties of a task
type Task struct {
	Name string
	Done bool
	cmd  []string
}

// Create creates and returns a Task
func Create(name string, cmd ...string) Task {
	return Task{
		Name: name,
		Done: false,
		cmd:  cmd,
	}
}

// execute executes a process
func execute(wg *sync.WaitGroup, name string, args ...string) {
	cmd := exec.Command(name, args...)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	defer wg.Done()
}

// Start begins executing the task
func (task *Task) Start(parentWg *sync.WaitGroup) {
	var wg sync.WaitGroup

	wg.Add(1)
	go execute(&wg, task.cmd[0], task.cmd[1:]...)
	wg.Wait()

	task.complete()

	defer parentWg.Done()
}

// complete sets the Task done
func (task *Task) complete() {
	task.Done = true
}
