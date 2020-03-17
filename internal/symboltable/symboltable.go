package symboltable

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	"github.com/merenbach/syllogism/internal/symbol"
)

// A SymbolTable contains a list of symbols.
type SymbolTable struct {
	Symbols              []*symbol.Symbol
	HighestLocationUsed  int
	NegativePremiseCount int
}

// New symbol table.
func New(size int) *SymbolTable {
	t := SymbolTable{
		Symbols: make([]*symbol.Symbol, size),
	}
	for i := range t.Symbols {
		t.Symbols[i] = &symbol.Symbol{}
	}
	return &t
}

// Dump values of variables in a SymbolTable.
func (t *SymbolTable) Dump() string {
	dump := new(bytes.Buffer)
	fmt.Fprintf(dump, "Highest symbol table loc. used: %d  Negative premises: %d\n", t.HighestLocationUsed, t.NegativePremiseCount)
	if t.HighestLocationUsed != 0 {
		w := tabwriter.NewWriter(dump, 0, 0, 2, ' ', 0)
		fmt.Fprint(w, "Adr.\tart.\tterm\ttype\toccurs\tdist. count")
		// for address, symbol := range t.Symbols
		for address := 1; address <= t.HighestLocationUsed; address++ {
			symbolDump := t.Symbols[address].Dump()
			fmt.Fprintf(w, "\n%d\t%s", address, symbolDump)
		}
		w.Flush()
	}
	return dump.String()
}

// Iterate over a symbol table with a given function.
// This function should return `false` to continue.
// This function should return `true` when stopping condition is reached.
func (t *SymbolTable) Iterate(start int, f func(int, *symbol.Symbol) bool) {
	for i := start; i <= t.HighestLocationUsed; i++ {
		if f(i, t.Symbols[i]) {
			break
		}
	}
}

// Search a symbol table for a term matching a given string.
// Porting notes: All variable use is encapsulated, so if porting needs to be re-done in future, re-porting this function can be avoided by invoking the equivalent of `i1, b1 = search(start, w$)`.
func (t *SymbolTable) Search(start int, w string) (int, int) {
	// 3950
	//---Search T$() for W$ from I1 to L1---

	// If found, I1 = L1; else I1 = L1+1. B1 set to 1st empty loc.
	firstEmptyLocation := 0

	t.Iterate(start, func(i int, s *symbol.Symbol) bool {
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
