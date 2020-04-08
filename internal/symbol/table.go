package symbol

import (
	"fmt"

	"github.com/merenbach/syllogism/internal/term"
)

// A Table of symbols.
type Table []*Symbol

// Search symbol table for a term matching a given string, or return nil.
func (st Table) Search(w string, termType term.Type, msg bool) *Symbol {
	for _, sym := range st {
		if sym.Term != w {
			continue
		}

		if termType == term.TypeUndetermined {
			// TODO: while this matches the original BASIC, should this be && msg?
			if sym.TermType != term.TypeUndetermined || msg {
				fmt.Printf("Note: predicate term %q taken as the %s used earlier\n", w, sym.TermType)
			}
			return sym
		} else if sym.TermType == term.TypeUndetermined {
			if msg {
				fmt.Printf("Note: earlier use of %q taken as the %s used here\n", w, termType)
			}
			if termType == term.TypeDesignator {
				sym.DistributionCount = sym.Occurrences
			}
			sym.TermType = termType
			return sym
		} else if termType == sym.TermType {
			return sym
		} else if msg {
			fmt.Printf("Warning: %s %q has also occurred as a %s\n", termType, w, termType.Other())
		}
	}

	return nil
}
