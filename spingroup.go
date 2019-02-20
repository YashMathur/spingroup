package spingroup

import (
	"sync"

	"./pkg/tasks"
)

// Spingroup struct holds information on the spin group
type Spingroup struct {
	chars []string
	tasks []*tasks.Task
}

var wg sync.WaitGroup

// Add appends a task to the Spingroup
func (sg *Spingroup) Add(name string, cmd ...string) {
	wg.Add(1)
	task := tasks.Create(name, cmd...)
	sg.tasks = append(sg.tasks, &task)

	go task.Start(&wg)
}

// Wait waits for all tasks in the spingroup to complete
func (sg *Spingroup) Wait() {
	for loop := true; loop; {
		var allDone = true

		for _, task := range sg.tasks {
			if !task.IsDone() {
				allDone = false
			} else {
				allDone = true
			}
		}

		if allDone {
			break
		}
	}

	wg.Wait()
}
