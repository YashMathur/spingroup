package spingroup

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
	"unicode/utf8"

	"github.com/YashMathur/spingroup/pkg/tasks"
	"github.com/YashMathur/spingroup/pkg/ui"
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

func cleanup() {
	ui.Show()
}

// Add appends a task to the Spingroup
func (sg *Spingroup) Add(name string, cmd ...string) {
	wg.Add(1)
	task := tasks.Create(name, cmd...)
	sg.tasks = append(sg.tasks, &task)

	go task.Start(&wg)
}

// AddFunc adds a task that executes a function
func (sg *Spingroup) AddFunc(name string, fn func(*sync.WaitGroup, *bool, *bool)) {
	wg.Add(1)

	task := tasks.Create(name)
	sg.tasks = append(sg.tasks, &task)

	go fn(&wg, &task.Done, &task.Success)
}

// Wait waits for all tasks in the spingroup to complete
func (sg *Spingroup) Wait() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	gui := ui.Create(len(sg.tasks))
	lenSpinner := utf8.RuneCountInString(gui.GetSpinner())

	for _, task := range sg.tasks {
		fmt.Printf("%s %s\n", strings.Repeat(" ", lenSpinner), task.Name)
	}

	ui.Hide()
	ui.Up(len(sg.tasks) + 1)

	for loop := true; loop; {
		var allDone = true

		for _, task := range sg.tasks {
			ui.Down(1)
			ui.MoveHorizontal(1)

			if !task.Done {
				fmt.Print(gui.GetSpinner())
				allDone = false
			} else {
				fmt.Print(task.FinalMessage())
			}
		}

		if allDone {
			break
		} else {
			ui.Up(len(sg.tasks))
		}

		gui.Step()
		time.Sleep(sg.sleep)
	}

	fmt.Printf("\n")
	ui.Show()

	wg.Wait()
}
