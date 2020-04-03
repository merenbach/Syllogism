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

// HighestLocationUsed returns the highest location used.
func (st *SymbolTable) HighestLocationUsed() int {
	return len(st.Symbols) - 1
}

// IncreaseLocationMax increases the highest location used.
func (st *SymbolTable) IncreaseLocationMax() {
	st.Symbols = append(st.Symbols, &symbol.Symbol{})
}

// New symbol table.
func New(size int) *SymbolTable {
	t := SymbolTable{
		Symbols:         make([]*symbol.Symbol, 1),
		ConclusionTerms: make([]*symbol.Symbol, size),
	}
	for i := range t.Symbols {
		t.Symbols[i] = &symbol.Symbol{}
	}
	return &t
}

// Iterate over a symbol table with a given function.
// This function should return `false` to continue.
// This function should return `true` when stopping condition is reached.
func (st *SymbolTable) Iterate(start int, f func(int, *symbol.Symbol) bool) {
	for i := start; i <= st.HighestLocationUsed(); i++ {
		if f(i, st.Symbols[i]) {
			break
		}
	}
}

// // Prune orphaned terms with no occurrences.
// func (st *SymbolTable) Prune() {
// 	var ss []*symbol.Symbol
// 	for i := 1; i <= st.HighestLocationUsed(); i++ {
// 		s := st.Symbols[i]
// 		if s.Occurrences > 0 {
// 			ss = append(ss, s)
// 		}
// 	}

// 	st.Symbols = ss
// }

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
func (st *SymbolTable) Search(start int, w string) (int, int) {
	// 3950
	//---Search T$() for W$ from I1 to L1---

	// If found, I1 = L1; else I1 = L1+1. B1 set to 1st empty loc.
	var firstEmptyLocation int

	st.Iterate(start, func(i int, s *symbol.Symbol) bool {
		if s.Term == w {
			return true
		}

		if s.Occurrences == 0 && firstEmptyLocation == 0 {
			firstEmptyLocation = i
		}

		start = i + 1
		return false
	})

	return start, firstEmptyLocation
}
