package symboltable

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	"github.com/merenbach/syllogism/internal/article"
	"github.com/merenbach/syllogism/internal/symbol"
	"github.com/merenbach/syllogism/internal/term"
)

// A SymbolTable contains a list of symbols.
type SymbolTable struct {
	Symbols              []*symbol.Symbol
	HighestLocationUsed  int
	NegativePremiseCount int
	CArray               []int
}

// New symbol table.
func New(size int) *SymbolTable {
	t := SymbolTable{
		Symbols: make([]*symbol.Symbol, size),
		CArray:  make([]int, size),
	}
	for i := range t.Symbols {
		t.Symbols[i] = &symbol.Symbol{}
	}
	return &t
}

// Compute a conclusion.
func (st *SymbolTable) Compute(symbol1 *symbol.Symbol, symbol2 *symbol.Symbol) string {
	if st.NegativePremiseCount == 0 {
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
func (st *SymbolTable) Dump() string {
	dump := new(bytes.Buffer)
	fmt.Fprintf(dump, "Highest symbol table loc. used: %d  Negative premises: %d\n", st.HighestLocationUsed, st.NegativePremiseCount)
	if st.HighestLocationUsed != 0 {
		w := tabwriter.NewWriter(dump, 0, 0, 2, ' ', 0)
		fmt.Fprint(w, "Adr.\tart.\tterm\ttype\toccurs\tdist. count")
		// for address, symbol := range t.Symbols
		for address := 1; address <= st.HighestLocationUsed; address++ {
			symbolDump := st.Symbols[address].Dump()
			fmt.Fprintf(w, "\n%d\t%s", address, symbolDump)
		}
		w.Flush()
	}
	return dump.String()
}

// Iterate over a symbol table with a given function.
// This function should return `false` to continue.
// This function should return `true` when stopping condition is reached.
func (st *SymbolTable) Iterate(start int, f func(int, *symbol.Symbol) bool) {
	for i := start; i <= st.HighestLocationUsed; i++ {
		if f(i, st.Symbols[i]) {
			break
		}
	}
}

// Search a symbol table for a term matching a given string.
// Porting notes: All variable use is encapsulated, so if porting needs to be re-done in future, re-porting this function can be avoided by invoking the equivalent of `i1, b1 = search(start, w$)`.
func (st *SymbolTable) Search(start int, w string) (int, int) {
	// 3950
	//---Search T$() for W$ from I1 to L1---

	// If found, I1 = L1; else I1 = L1+1. B1 set to 1st empty loc.
	firstEmptyLocation := 0

	st.Iterate(start, func(i int, s *symbol.Symbol) bool {
		if s.Term == w {
			return true
		}

		if s.Empty() && firstEmptyLocation == 0 {
			firstEmptyLocation = i
		}

		start = i + 1
		return false
	})

	return start, firstEmptyLocation
}
