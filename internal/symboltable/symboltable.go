package symboltable

import (
	"github.com/merenbach/syllogism/internal/symbol"
)

// A SymbolTable contains a list of symbols.
type SymbolTable struct {
	Symbols             []*symbol.Symbol
	highestLocationUsed int
	CArray              []int
}

// HighestLocationUsed returns the highest location used.
func (st *SymbolTable) HighestLocationUsed() int {
	return st.highestLocationUsed
}

// IncreaseLocationMax increases the highest location used.
func (st *SymbolTable) IncreaseLocationMax() {
	st.highestLocationUsed++
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
