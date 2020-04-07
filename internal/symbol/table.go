package symbol

// A Table of symbols.
type Table []*Symbol

// Search a symbol table for a term matching a given string.
// Porting notes: All variable use is encapsulated, so if porting needs to be re-done in future, re-porting this function can be avoided by invoking the equivalent of `i1, b1 = search(start, w$)`.
func (st Table) Search(start int, w string) int {
	// 3950
	//---Search T$() for W$ from I1 to L1---

	// If found, I1 = L1; else I1 = L1+1. B1 set to 1st empty loc.
	for i, s := range st {
		if i < start {
			continue
		}

		if s.Term == w {
			break
		}

		start = i + 1
	}

	return start
}
