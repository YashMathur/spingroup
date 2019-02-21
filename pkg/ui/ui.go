package ui

// UI contains info on the UI of the spingroup
type UI struct {
	tasks    int
	StartRow int
	Spinner  []string
	Reverse  bool
	iter     int
	step     int
}

// Create creates a new UI struct
func Create(tasks int) *UI {
	x := CursorPosition()

	return &UI{
		tasks:    tasks,
		StartRow: x,
		Spinner:  []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		Reverse:  false,
		iter:     0,
		step:     1,
	}
}

// GetLength returns the length of the spinner animation
func (ui *UI) GetLength() int {
	return len(ui.Spinner)
}

// GetSpinner returns spinner at a point in time
func (ui *UI) GetSpinner() string {
	return ui.Spinner[ui.iter]
}

// Step steps the spinner index to the next index depending on ui.Reverse
func (ui *UI) Step() {
	if ui.Reverse {
		if ui.step > 0 && (ui.iter+ui.step)%ui.GetLength() == 0 {
			ui.step = ui.step * -1
		}

		if ui.step < 0 && (ui.iter+ui.step)%ui.GetLength() == -1 {
			ui.step = ui.step * -1
		}
	}

	ui.iter = (ui.iter + ui.step) % ui.GetLength()
}
