package premiseset

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	"github.com/merenbach/syllogism/internal/article"
	"github.com/merenbach/syllogism/internal/premise"
	"github.com/merenbach/syllogism/internal/symbol"
	"github.com/merenbach/syllogism/internal/symboltable"
	"github.com/merenbach/syllogism/internal/term"
)

// Set of all premises.
type Set struct {
	Premises    []*premise.Premise
	SymbolTable *symboltable.SymbolTable
	LinkOrder   []int
	LArray      []int
	AArray      []int
}

// Enter line into list.
func (ps *Set) Enter(n int, s string) *premise.Premise {
	// Silently delete any existing line matching this line number
	_ = ps.Delete(n)
	newPremise := premise.New(n, s)

	var (
		localint_i  int
		localint_j1 int
	)

	for localint_i = 0; ; localint_i = localint_j1 {
		localint_j1 = ps.LArray[localint_i]

		if localint_j1 == 0 {
			break
		}

		if n < ps.Premises[localint_j1].Number {
			break
		}
	}

	a1 := ps.AArray[ps.AArray[0]]
	ps.Premises[a1] = newPremise
	ps.LArray[localint_i] = a1
	ps.LArray[a1] = localint_j1
	ps.AArray[0]++

	return newPremise
}

// Delete a line.
func (ps *Set) Delete(n int) error {
	for i := 0; ; i = ps.LArray[i] {
		j1 := ps.LArray[i]

		if j1 == 0 {
			return fmt.Errorf("Line %d not found", n)
		} else if n == ps.Premises[j1].Number {
			ps.AArray[0]--
			ps.AArray[ps.AArray[0]] = j1
			ps.LArray[i] = ps.LArray[j1]
			ps.Premises[j1].Decrement()
			break
		}
	}

	return nil
}

// Compute a conclusion.
func (ps *Set) Compute(symbol1 *symbol.Symbol, symbol2 *symbol.Symbol) string {
	if ps.LArray[0] == 0 {
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
	ps.print(ps.rawPremises(), analyze)
}

// Link output for premises, optionally in distribution-analysis format.
func (ps *Set) Link(max int, analyze bool) {
	ps.print(ps.linkedPremises(max), analyze)
}

// LinkedPremise returns the linked premise with the given index.
func (ps *Set) LinkedPremise(i int) *premise.Premise {
	return ps.Premises[ps.LinkOrder[i]]
}

// rawPremises returns premises in entry order.
func (ps *Set) rawPremises() []*premise.Premise {
	pp := make([]*premise.Premise, 0)
	for i := ps.LArray[0]; i != 0; i = ps.LArray[i] {
		pp = append(pp, ps.Premises[i])
	}
	return pp
}

// LinkedPremises returns premises in link order.
func (ps *Set) linkedPremises(max int) []*premise.Premise {
	pp := make([]*premise.Premise, 0)
	for i := 0; i < max; i++ {
		idx := ps.LinkOrder[i+1]
		pp = append(pp, ps.Premises[idx])
	}
	return pp
}

// NegativePremiseCount returns the count of negative premises.
func (ps *Set) NegativePremiseCount() int {
	var negativePremises int
	// TODO: is there a better way to iterate?
	for i := ps.LArray[0]; i != 0; i = ps.LArray[i] {
		prem := ps.Premises[i]
		if prem.Form.IsNegative() {
			negativePremises++
		}
	}
	return negativePremises
}

// New premise set with the given size.
func New(size int) *Set {
	ps := &Set{
		Premises:    make([]*premise.Premise, size),
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
