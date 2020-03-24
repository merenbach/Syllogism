package premise

import (
	"fmt"
	"strings"

	"github.com/merenbach/syllogism/internal/form"
	"github.com/merenbach/syllogism/internal/symbol"
	"github.com/merenbach/syllogism/internal/term"
)

// Set of all premises.
type Set struct {
	Premises  []*Premise
	LinkOrder []int
}

// List output for premises.
// TODO: use tabwriter for distribution-analysis format?
func (ps *Set) List(lSlice []int, analyze bool) string {
	var b strings.Builder
	for i := lSlice[0]; i != 0; i = lSlice[i] {
		prem := ps.Premises[i]
		if !analyze {
			b.WriteString(fmt.Sprintf("%d  %s\n", prem.Number, prem.Statement))
		} else {
			b.WriteString(fmt.Sprintf("%d  ", prem.Number))

			if prem.Form < 6 && prem.Predicate.TermType == term.TypeDesignator {
				prem.Form += 2
			}

			if prem.Form < 4 {
				b.WriteString(fmt.Sprintf("%s  ", prem.Form.Quantifier()))
			}

			b.WriteString(fmt.Sprintf("%s%s%s  %s%s\n", prem.Subject.Term, prem.Form.SymbolForTermA(), prem.Form.Copula(), prem.Predicate.Term, prem.Form.SymbolForTermB()))
		}
	}
	return b.String()
}

// Link output for premises.
// TODO: use tabwriter for distribution-analysis format?
func (ps *Set) Link(max int, analyze bool) string {
	var b strings.Builder
	for i := 1; i <= max; i++ {
		idx := ps.LinkOrder[i]
		prem := ps.Premises[idx]
		if !analyze {
			b.WriteString(fmt.Sprintf("%d  %s\n", prem.Number, prem.Statement))
		} else {
			b.WriteString(fmt.Sprintf("%d  ", prem.Number))
			if prem.Form < 6 && prem.Predicate.TermType == term.TypeDesignator {
				prem.Form += 2
			}
			if prem.Form < 4 {
				b.WriteString(fmt.Sprintf("%s  ", prem.Form.Quantifier()))
			}
			b.WriteString(fmt.Sprintf("%s%s%s  %s%s\n", prem.Subject.Term, prem.Form.SymbolForTermA(), prem.Form.Copula(), prem.Predicate.Term, prem.Form.SymbolForTermB()))
		}
	}
	return b.String()
}

// NewPremiseSet creates a new premise set with the given size.
func NewPremiseSet(size int) *Set {
	return &Set{
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

// // Empty determines whether a line is empty.
// func (pr *Premise) Empty() bool {
// 	return pr.Statement == ""
// }
