package premise

import (
	"fmt"

	"github.com/merenbach/syllogism/internal/form"
	"github.com/merenbach/syllogism/internal/symbol"
)

// A PremiseSet stores a list of all premises.
type PremiseSet struct {
	Premises  []*Premise
	LinkOrder []int
}

// NewPremiseSet creates a new premise set with the given size.
func NewPremiseSet(size int) *PremiseSet {
	return &PremiseSet{
		Premises:  make([]*Premise, size),
		LinkOrder: make([]int, size),
	}
}

// Premise stores a line number and program statement.
// TODO: store numbers with PremiseSet if possible
type Premise struct {
	Number                   int
	Statement                string
	Form                     form.Form
	ExperimentalLinkingOrder int
	Subject                  *symbol.Symbol
	Predicate                *symbol.Symbol
}

func (pr *Premise) String() string {
	return fmt.Sprintf("%d  %s", pr.Number, pr.Statement)
}

// Empty determines whether a line is empty.
func (pr *Premise) Empty() bool {
	return pr.Statement == ""
}
