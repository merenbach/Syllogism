package symboltable

import (
	"github.com/merenbach/syllogism/internal/symbol"
)

// A SymbolTable contains a list of symbols.
type SymbolTable struct {
	Symbols []*symbol.Symbol
	// ConclusionTerms are major and minor (i.e., all the non-middle) terms
	ConclusionTerms []*symbol.Symbol
}

// New symbol table.
func New(size int) *SymbolTable {
	return &SymbolTable{
		Symbols:         make([]*symbol.Symbol, 0),
		ConclusionTerms: make([]*symbol.Symbol, size),
	}
}

// Prune orphaned terms with no occurrences.
func (st *SymbolTable) Prune() {
	var ss []*symbol.Symbol
	for _, s := range st.Symbols {
		if s.Occurrences > 0 {
			ss = append(ss, s)
		}
	}

	st.Symbols = ss
}

// // Delete a term from the table.
// func (st *SymbolTable) Delete(sym *symbol.Symbol) {
// 	for i, s := range st.Symbols {
// 		if s.Term == sym.Term {
// 			// Delete without leaving uncollected pointers
// 			// https://github.com/golang/go/wiki/SliceTricks
// 			if i < len(st.Symbols)-1 {
// 				copy(st.Symbols[i:], st.Symbols[i+1:])
// 			}
// 			st.Symbols[len(st.Symbols)-1] = nil
// 			st.Symbols = st.Symbols[:len(st.Symbols)-1]
// 			break
// 		}
// 	}
// }

// Search a symbol table for a term matching a given string.
// Porting notes: All variable use is encapsulated, so if porting needs to be re-done in future, re-porting this function can be avoided by invoking the equivalent of `i1, b1 = search(start, w$)`.
func (st *SymbolTable) Search(start int, w string) int {
	// 3950
	//---Search T$() for W$ from I1 to L1---

	// If found, I1 = L1; else I1 = L1+1. B1 set to 1st empty loc.
	for i, s := range st.Symbols {
		if s.Term == w {
			break
		}

		start = i + 1
	}

	return start
}
