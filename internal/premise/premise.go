package premise

import (
	"fmt"

	"github.com/merenbach/syllogism/internal/form"
	"github.com/merenbach/syllogism/internal/symbol"
	"github.com/merenbach/syllogism/internal/symboltable"
	"github.com/merenbach/syllogism/internal/term"
)

// Set of all premises.
type Set struct {
	Premises    []*Premise
	SymbolTable *symboltable.SymbolTable
	LinkOrder   []int
	LArray      []int
	AArray      []int
}

// List output for premises, optionally in distribution-analysis format.
// TODO: use tabwriter for distribution-analysis format?
func (ps *Set) List(analyze bool) {
	for i := ps.LArray[0]; i != 0; i = ps.LArray[i] {
		prem := ps.Premises[i]
		if !analyze {
			fmt.Printf("%d  %s\n", prem.Number, prem.Statement)
		} else {
			fmt.Printf("%d  ", prem.Number)

			if prem.Form < 6 && prem.Predicate.TermType == term.TypeDesignator {
				prem.Form += 2
			}

			if prem.Form < 4 {
				fmt.Printf("%s  ", prem.Form.Quantifier())
			}

			fmt.Printf("%s%s%s  %s%s\n", prem.Subject.Term, prem.Form.SymbolForTermA(), prem.Form.Copula(), prem.Predicate.Term, prem.Form.SymbolForTermB())
		}
	}
}

// Link output for premises, optionally in distribution-analysis format.
// TODO: use tabwriter for distribution-analysis format?
func (ps *Set) Link(max int, analyze bool) {
	for i := 1; i <= max; i++ {
		idx := ps.LinkOrder[i]
		prem := ps.Premises[idx]
		if !analyze {
			fmt.Printf("%d  %s\n", prem.Number, prem.Statement)
		} else {
			fmt.Printf("%d  ", prem.Number)
			if prem.Form < 6 && prem.Predicate.TermType == term.TypeDesignator {
				prem.Form += 2
			}
			if prem.Form < 4 {
				fmt.Printf("%s  ", prem.Form.Quantifier())
			}
			fmt.Printf("%s%s%s  %s%s\n", prem.Subject.Term, prem.Form.SymbolForTermA(), prem.Form.Copula(), prem.Predicate.Term, prem.Form.SymbolForTermB())
		}
	}
}

// // NegativePremiseCount returns the count of negative premises.
// func (ps *Set) NegativePremiseCount() int {
// 	var negativePremises int
// 	// TODO: is there a better way to iterate?
// 	for i := ps.LArray[0]; i != 0; i = ps.LArray[i] {
// 		prem := ps.Premises[i]
// 		if prem.Form.IsNegative() {
// 			negativePremises++
// 		}
// 	}
// 	return negativePremises
// }

// NewPremiseSet creates a new premise set with the given size.
func NewPremiseSet(size int) *Set {
	ps := &Set{
		Premises:    make([]*Premise, size),
		SymbolTable: symboltable.New(size + 2),
		LinkOrder:   make([]int, size),
		AArray:      make([]int, size),
		LArray:      make([]int, size),
	}

	for i := range ps.AArray {
		ps.AArray[i] = i
	}
	ps.AArray[0] = 1

	return ps
}

// Premise stores a line number and program statement.
// TODO: store numbers with PremiseSet if possible
type Premise struct {
	Number    int
	Statement string
	Form      form.Form
	Subject   *symbol.Symbol
	Predicate *symbol.Symbol
}

func (pr *Premise) String() string {
	return fmt.Sprintf("%d  %s", pr.Number, pr.Statement)
}

// // Decrement table entries.
// TODO: this will be perfect if we can note the negative premise count automatically, rather than keeping a variable for it
// func (pr *Premise) Decrement(st *symboltable.SymbolTable) {
// 	var (
// 		pDecrement bool
// 		qDecrement bool
// 	)

// 	if pr.Form.IsNegative() {
// 		st.NegativePremiseCount--
// 		qDecrement = true
// 	} else if pr.Predicate.TermType == term.TypeDesignator {
// 		qDecrement = true
// 	}

// 	if pr.Form >= 2 {
// 		pDecrement = true
// 	}

// 	pr.Subject.ReduceDistributionCount(pDecrement)
// 	pr.Predicate.ReduceDistributionCount(qDecrement)
// }

// // Empty determines whether a line is empty.
// func (pr *Premise) Empty() bool {
// 	return pr.Statement == ""
// }
