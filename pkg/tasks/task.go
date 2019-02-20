package tasks

import (
	"fmt"
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

// IsDone checks whether a Task is done
func (task *Task) IsDone() bool {
	return task.done
}

// Name returns the task's name
func (task *Task) Name() string {
	return task.name
}

// complete sets the Task done
func (task *Task) complete() {
	fmt.Printf("Completing %s\n", task.name)
	task.done = true
}
