package ui

// UI contains info on the UI of the spingroup
type UI struct {
	tasks    int
	startRow int
	spinner  []string
}

// Create creates a new UI struct
func Create(tasks int) *UI {
	x := CursorPosition()

	return &UI{
		tasks:    tasks,
		startRow: x,
		spinner:  []string{"⬒", "⬔", "⬓", "⬕"},
	}
}

// StartRow returns the UI's start row
func (ui *UI) StartRow() int {
	return ui.startRow
}

// Length returns the length of the spinner animation
func (ui *UI) Length() int {
	return len(ui.spinner)
}

// Spinner returns spinner at a point in time
func (ui *UI) Spinner(i int) string {
	return ui.spinner[i]
}
