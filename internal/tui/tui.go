package tui

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/merenbach/syllogism/internal/form"
)

// A PremiseSet stores a list of all premises.
type PremiseSet struct {
	Premises []*ProgramLine
}

// NewPremiseSet creates a new premise set with the given size.
func NewPremiseSet(size int) *PremiseSet {
	return &PremiseSet{
		Premises: make([]*ProgramLine, size),
	}
}

// ProgramLine stores a line number and program statement.
// TODO: rename to Premise
// TODO: store numbers with PremiseSet if possible
type ProgramLine struct {
	Number    int
	Statement string
	Form      form.Form
}

func (pl *ProgramLine) String() string {
	return fmt.Sprintf("%d  %s", pl.Number, pl.Statement)
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
