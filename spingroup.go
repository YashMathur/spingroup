package spingroup

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

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
	gui := ui.Create(len(sg.tasks))
	lenSpinner := utf8.RuneCountInString(gui.GetSpinner())

	for _, task := range sg.tasks {
		fmt.Printf("%s %s\n", strings.Repeat(" ", lenSpinner), task.Name)
	}

	ui.Hide()

	for loop := true; loop; {
		var allDone = true

		for i, task := range sg.tasks {
			row := gui.StartRow + i

			ui.Move(row, 1)

			if !task.Done {
				fmt.Print(gui.GetSpinner())
				allDone = false
			} else {
				fmt.Print("✓")
			}
		}

		if allDone {
			break
		}

		gui.Step()
		time.Sleep(sg.sleep)
	}

	fmt.Printf("\n")
	ui.Show()

	wg.Wait()
}
