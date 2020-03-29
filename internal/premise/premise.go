package premise

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	"github.com/merenbach/syllogism/internal/article"
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

// LinkedPremise returns the linked premise with the given index.
func (ps *Set) LinkedPremise(i int) *Premise {
	return ps.Premises[ps.LinkOrder[i]]
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

// Decrement table entries.
func (pr *Premise) Decrement() {
	var (
		pDecrement bool
		qDecrement bool
	)

	if pr.Form.IsNegative() {
		qDecrement = true
	} else if pr.Predicate.TermType == term.TypeDesignator {
		qDecrement = true
	}

	if pr.Form >= 2 {
		pDecrement = true
	}

	pr.Subject.ReduceDistributionCount(pDecrement)
	pr.Predicate.ReduceDistributionCount(qDecrement)
}

// Empty determines whether a line is empty.
func (pr *Premise) Empty() bool {
	return pr.Statement == ""
}
