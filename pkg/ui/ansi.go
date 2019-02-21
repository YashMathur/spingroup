package ui

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

const esc = "\x1b["

func ansi(format string, args ...interface{}) {
	fmt.Print(fmt.Sprintf("%s%s", esc, fmt.Sprintf(format, args...)))
}

// Up moves the cursor up n rows
func Up(n int) {
	ansi("%d%s", n, "A")
}

// Down moves the cursor down n rows
func Down(n int) {
	ansi("%d%s", n, "B")
}

// Right moves the cursor right n cols
func Right(n int) {
	ansi("%d%s", n, "C")
}

// Left moves the cursor left n cols
func Left(n int) {
	ansi("%d%s", n, "D")
}

// Move moves cursor to position r, c
func Move(r, c int) {
	ansi("%d;%d%s", r, c, "H")
}

// Clearleft clears the line
func Clearleft() {
	ansi("%d%s", 1, "K")
}

// Hide makes the cursor invisible
func Hide() {
	ansi("%s", "?25l")
}

// Show makes the cursor visible
func Show() {
	ansi("%s", "?25h")
}

// CursorPosition prints the cursor's position
func CursorPosition() int {
	rawMode := exec.Command("/bin/stty", "raw")
	rawMode.Stdin = os.Stdin
	_ = rawMode.Run()
	rawMode.Wait()

	cmd := exec.Command("printf", fmt.Sprintf("%c[6n", 27))
	randomBytes := &bytes.Buffer{}
	cmd.Stdout = randomBytes

	_ = cmd.Start()

	reader := bufio.NewReader(os.Stdin)
	cmd.Wait()

	fmt.Print(randomBytes)
	text, _ := reader.ReadSlice('R')

	fmt.Printf(string([]byte{0x1b, '[', '1', 'K'}))

	rawModeOff := exec.Command("/bin/stty", "-raw")
	rawModeOff.Stdin = os.Stdin
	_ = rawModeOff.Run()
	rawModeOff.Wait()

	var line []string

	if strings.Contains(string(text), ";") {
		re := regexp.MustCompile(`[0-9]+`)
		line = re.FindAllString(string(text), -1)
	}

	row, err := strconv.Atoi(line[0])

	if err != nil {
		return -1
	}

	return row
}
