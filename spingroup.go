package spingroup

import (
	"fmt"
	"sync"
	"time"

	"./pkg/tasks"
	"./pkg/ui"
)

// Spingroup struct holds information on the spin group
type Spingroup struct {
	chars []string
	sleep time.Duration
	tasks []*tasks.Task
}

var wg sync.WaitGroup

// Create creates a Spingroup
func Create(sleep time.Duration) *Spingroup {
	return &Spingroup{
		sleep: sleep,
	}
}

// Add appends a task to the Spingroup
func (sg *Spingroup) Add(name string, cmd ...string) {
	wg.Add(1)
	task := tasks.Create(name, cmd...)
	sg.tasks = append(sg.tasks, &task)

	go task.Start(&wg)
}

// Wait waits for all tasks in the spingroup to complete
func (sg *Spingroup) Wait() {
	var gui = ui.Create(len(sg.tasks))

	for i, task := range sg.tasks {
		row := gui.StartRow() + i

		fmt.Print(ui.Move(row, 1))
		fmt.Printf("  %s %d", task.Name(), i)
	}

	fmt.Print(ui.Hide())
	iter := 0

	for loop := true; loop; {
		var allDone = true

		for i, task := range sg.tasks {
			row := gui.StartRow() + i

			fmt.Print(ui.Move(row, 1))

			if !task.Done {
				fmt.Print(gui.Spinner(iter))
				allDone = false
			} else {
				fmt.Print("g")
			}
		}

		if allDone {
			break
		}

		iter = (iter + 1) % gui.Length()
		time.Sleep(sg.sleep)
	}
	fmt.Printf("\n")
	fmt.Print(ui.Show())

	wg.Wait()
}
