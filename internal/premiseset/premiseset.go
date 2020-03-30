package premiseset

import (
	"bytes"
	"fmt"
	"sort"
	"text/tabwriter"

	"github.com/merenbach/syllogism/internal/article"
	"github.com/merenbach/syllogism/internal/premise"
	"github.com/merenbach/syllogism/internal/symbol"
	"github.com/merenbach/syllogism/internal/symboltable"
	"github.com/merenbach/syllogism/internal/term"
)

// Set of all premises.
type Set struct {
	NewPremises    []*premise.Premise
	LinkedPremises []*premise.Premise
	SymbolTable    *symboltable.SymbolTable
}

// Enter line into list.
func (ps *Set) Enter(n int, s string) *premise.Premise {
	// Silently delete any existing line matching this line number
	_ = ps.Delete(n)
	newPremise := premise.New(n, s)

	// NOTE: for new experimental refactor
	ps.NewPremises = append(ps.NewPremises, newPremise)
	ps.Sort()

	return newPremise
}

// Delete a line.
func (ps *Set) Delete(n int) error {
	for i, p := range ps.NewPremises {
		if p.Number == n {
			p.Decrement()
			ps.NewPremises = append(ps.NewPremises[:i], ps.NewPremises[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Line %d not found", n)
}

// Sort premises by line number.
func (ps *Set) Sort() {
	sort.Slice(ps.NewPremises, func(i, j int) bool { return ps.NewPremises[i].Number < ps.NewPremises[j].Number })
}

// // Len returns the length of the premise set.
// func (ps *Set) Len() int {
// 	return len(ps.NewPremises)
// }

// Compute a conclusion.
func (ps *Set) Compute(symbol1 *symbol.Symbol, symbol2 *symbol.Symbol) string {
	if ps.Empty() {
		return "A is A"
	}

	if ps.NegativePremiseCount() == 0 {
		// affirmative conclusion
		// TODO: can we push more of these conditionals inside the symbol type?
		if symbol1.DistributionCount > 0 {
			return symbol1.ConclusionForAllIs(symbol2)
		} else if symbol2.DistributionCount > 0 {
			return symbol2.ConclusionForAllIs(symbol1)
		} else if symbol1.ArticleType != article.TypeNone || symbol2.ArticleType == article.TypeNone {
			return symbol1.ConclusionForSomeIs(symbol2)
		} else {
			return symbol2.ConclusionForSomeIs(symbol1)
		}

	} else {
		// negative conclusion
		if symbol2.DistributionCount == 0 {
			return fmt.Sprintf("Some %s is not %s%s", symbol2.Term, symbol1.ArticleType, symbol1.Term)
		} else if symbol1.DistributionCount == 0 {
			return fmt.Sprintf("Some %s is not %s%s", symbol1.Term, symbol2.ArticleType, symbol2.Term)
		} else if symbol1.TermType == term.TypeDesignator {
			return fmt.Sprintf("%s is not %s%s", symbol1.Term, symbol2.ArticleType, symbol2.Term)
		} else if symbol2.TermType == term.TypeDesignator {
			return fmt.Sprintf("%s is not %s%s", symbol2.Term, symbol1.ArticleType, symbol1.Term)
		} else if symbol1.ArticleType == article.TypeNone && symbol2.ArticleType != article.TypeNone {
			return fmt.Sprintf("No %s is %s%s", symbol2.Term, symbol1.ArticleType, symbol1.Term)
		} else {
			return fmt.Sprintf("No %s is %s%s", symbol1.Term, symbol2.ArticleType, symbol2.Term)
		}
	}
}

// Dump values of variables in a SymbolTable.
func (ps *Set) Dump() string {
	dump := new(bytes.Buffer)
	fmt.Fprintf(dump, "Highest symbol table loc. used: %d  Negative premises: %d\n", ps.SymbolTable.HighestLocationUsed, ps.NegativePremiseCount())
	if ps.SymbolTable.HighestLocationUsed != 0 {
		w := tabwriter.NewWriter(dump, 0, 0, 2, ' ', 0)
		fmt.Fprint(w, "Adr.\tart.\tterm\ttype\toccurs\tdist. count")
		// for address, symbol := range t.Symbols
		for address := 1; address <= ps.SymbolTable.HighestLocationUsed; address++ {
			symbolDump := ps.SymbolTable.Symbols[address].Dump()
			fmt.Fprintf(w, "\n%d\t%s", address, symbolDump)
		}
		w.Flush()
	}
	return dump.String()
}

// Print ordered output of premises.
// TODO: use tabwriter for distribution-analysis format?
func (ps *Set) print(premises []*premise.Premise, analyze bool) {
	for _, prem := range premises {
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

// List output for premises, optionally in distribution-analysis format.
func (ps *Set) List(analyze bool) {
	ps.print(ps.NewPremises, analyze)
}

// Link output for premises, optionally in distribution-analysis format.
func (ps *Set) Link(max int, analyze bool) {
	ps.print(ps.LinkedPremises, analyze)
}

// NegativePremiseCount returns the count of negative premises.
func (ps *Set) NegativePremiseCount() int {
	var negativePremises int
	for _, p := range ps.NewPremises {
		if p.Form.IsNegative() {
			negativePremises++
		}
	}
	return negativePremises
}

// Empty determines whether the premise set is empty.
func (ps *Set) Empty() bool {
	return len(ps.NewPremises) == 0
}

// New premise set with the given size.
func New(size int) *Set {
	ps := &Set{
		NewPremises:    make([]*premise.Premise, 0),
		LinkedPremises: make([]*premise.Premise, size),
		SymbolTable:    symboltable.New(size + 2),
	}

	return ps
}
