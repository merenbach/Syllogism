package premise

import (
	"fmt"

	"github.com/merenbach/syllogism/internal/form"
)

// A PremiseSet stores a list of all premises.
type PremiseSet struct {
	Premises    []*Premise
	SubjIndices []int
	PredIndices []int
	LinkOrder   []int
}

// NewPremiseSet creates a new premise set with the given size.
func NewPremiseSet(size int) *PremiseSet {
	return &PremiseSet{
		Premises:    make([]*Premise, size),
		SubjIndices: make([]int, size),
		PredIndices: make([]int, size),
		LinkOrder:   make([]int, size),
	}
}

// Premise stores a line number and program statement.
// TODO: store numbers with PremiseSet if possible
type Premise struct {
	Number                   int
	Statement                string
	Form                     form.Form
	ExperimentalLinkingOrder int
}

func (pr *Premise) String() string {
	return fmt.Sprintf("%d  %s", pr.Number, pr.Statement)
}

// Empty determines whether a line is empty.
func (pr *Premise) Empty() bool {
	return pr.Statement == ""
}
