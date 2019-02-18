package spingroup

import (
	"sync"

	"github.com/YashMathur/spingroup/pkg/tasks"
)

// Spingroup struct holds information on the spin group
type Spingroup struct {
	chars []string
	tasks []tasks.Task
}

var wg sync.WaitGroup

// Add appends a task to the Spingroup
func (sg *Spingroup) Add(name string, cmd ...string) {
	wg.Add(1)
	task := tasks.Create(name, cmd...)
	sg.tasks = append(sg.tasks, task)

	go task.Start(&wg)
}
