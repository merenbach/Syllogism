package symbol

import (
	"bytes"
	"fmt"
	"text/tabwriter"

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

// Dump values of variables in a symbol table.
// TODO: improve alignment?
func (st Table) Dump() string {
	dump := new(bytes.Buffer)
	if len(st) > 0 {
		w := tabwriter.NewWriter(dump, 0, 0, 2, ' ', 0)
		fmt.Fprint(w, "Adr.\tart.\tterm\ttype\toccurs\tdist. count")
		for i, s := range st {
			fmt.Fprintf(w, "\n%d\t%s", i+1, s.Dump())
		}
		w.Flush()
	}
	return dump.String()
}
