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

func ansi(pos int, code string) string {
	return fmt.Sprintf("%s%d%s", esc, pos, code)
}

// Up moves the cursor up n rows
func Up(n int) string {
	return ansi(n, "A")
}

// Down moves the cursor down n rows
func Down(n int) string {
	return ansi(n, "B")
}

// Right moves the cursor right n cols
func Right(n int) string {
	return ansi(n, "C")
}

// Left moves the cursor left n cols
func Left(n int) string {
	return ansi(n, "D")
}

// Move moves cursor to position r, c
func Move(r, c int) string {
	return fmt.Sprintf("%s%d;%d%s", esc, r, c, "H")
}

// Clearleft clears the line
func Clearleft() string {
	return ansi(1, "K")
}

// Hide makes the cursor invisible
func Hide() string {
	return fmt.Sprintf("%s?25l", esc)
}

// Show makes the cursor visible
func Show() string {
	return fmt.Sprintf("%s?25h", esc)
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
