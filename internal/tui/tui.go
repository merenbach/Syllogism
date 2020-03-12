package tui

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// ProgramLine stores a line number and program statement.
type ProgramLine struct {
	Number    int
	Statement string
}

func (pl *ProgramLine) String() string {
	return pl.plain()
}

func (pl *ProgramLine) plain() string {
	return fmt.Sprintf("%d %s", pl.Number, pl.Statement)
}

// Empty determines whether a line is empty.
func (pl *ProgramLine) Empty() bool {
	return pl.Statement == ""
}

// Clear the screen
func Clear() {
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		cmd := exec.Command("/usr/bin/clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		panic("Unsupported OS: " + runtime.GOOS)
	}
}
