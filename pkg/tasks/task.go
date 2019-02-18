package tasks

import (
	"log"
	"os/exec"
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
func Execute(name string, args ...string) {
	cmd := exec.Command(name, args...)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

// Start begins executing the task
func (task *Task) Start() {
	go Execute(task.cmd[0], task.cmd[1:]...)
}
