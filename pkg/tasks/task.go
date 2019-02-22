package tasks

import (
	"log"
	"os/exec"
	"sync"
)

// Task struct stores the properties of a task
type Task struct {
	Name    string
	Done    bool
	Success bool
	cmd     []string
	message Message
}

// Message stores messages to be displayed in an event
type Message struct {
	success string
	failure string
}

// Create creates and returns a Task
func Create(name string, cmd ...string) Task {
	return Task{
		Name:    name,
		Done:    false,
		cmd:     cmd,
		Success: false,
		message: Message{
			success: "✓",
			failure: "✗",
		},
	}
}

// execute executes a process
func (task *Task) execute(wg *sync.WaitGroup) {
	cmd := exec.Command(task.cmd[0], task.cmd[1:]...)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	} else if err == nil {
		task.Success = true
	}

	defer wg.Done()
}

// Start begins executing the task
func (task *Task) Start(parentWg *sync.WaitGroup) {
	var wg sync.WaitGroup

	wg.Add(1)
	go task.execute(&wg)
	wg.Wait()

	task.complete()

	defer parentWg.Done()
}

// FinalMessage returns the final message after a task is complete
func (task *Task) FinalMessage() string {
	if task.Success {
		return task.message.success
	}

	return task.message.failure
}

// complete sets the Task done
func (task *Task) complete() {
	task.Done = true
}
