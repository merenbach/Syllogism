package premise

import (
	"errors"
	"fmt"
	"log"
	"sort"

	"github.com/merenbach/syllogism/internal/article"
	"github.com/merenbach/syllogism/internal/form"
	"github.com/merenbach/syllogism/internal/help"
	"github.com/merenbach/syllogism/internal/symbol"
	"github.com/merenbach/syllogism/internal/term"
)

// Set of all premises.
type Set []*Premise

// Map of all premises.
type Map map[int]*Premise

// Keys from a premise map, in ascending order.
func (ps Map) Keys() []int {
	keys := make([]int, 0, len(ps))
	for k := range ps {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

// ToMap converts the premise set to a map.
func (ps Set) ToMap() Map {
	m := make(map[int]*Premise, len(ps))
	for _, p := range ps {
		m[p.Number] = p
	}
	return m
}

// Linked premises
func (ps Set) Linked(sym *symbol.Symbol) (Set, error) {
	var errCount int
	linkedPremises := ps.copy()
	if len(linkedPremises) > 1 {
		for i := range linkedPremises {
			k := linkedPremises.Find(sym, i)
			if k == (-1) {
				// Not found
				sym = linkedPremises[i].Predicate

				if errCount == 0 {
					errCount++
					fmt.Println("Not a syllogism: no way to order premises so that each premise")
					fmt.Println("shares exactly one term with its successor; there is a")
				}
				fmt.Println(help.ClosedLoopHelp)
			} else {
				prem := linkedPremises[k]
				// linkedPremises[k], linkedPremises[i] = linkedPremises[i], linkedPremises[k]
				// linkedPremises.Swap(k, i)

				// NOTE: THIS WHOLE SCOPE SIMPLY SWAPS WHAT'S AT POSITION LOCALINT_I WITH LOCALINT_K
				// NOTE: ...or does it???
				premises := make([]*Premise, 3)

				localint_n := 1
				premises[1] = linkedPremises[i]

				for m := i; m < k; m++ {
					localint_n = 3 - localint_n
					premises[localint_n] = linkedPremises[m+1]
					linkedPremises[m+1] = premises[3-localint_n]
				}

				linkedPremises[i] = premises[localint_n]
				ts := prem.ContrastingTerm(sym)
				if ts == nil {
					// We should never get here provided that prem contains sym
					log.Printf("Could not find symbol %+v in premise %+v\n", sym, prem)
				}
				sym = ts
			}

			if errCount != 0 {
				fmt.Println(linkedPremises[i])
			}
		}
	}

	if errCount > 0 {
		return nil, errors.New("TODO: improve this error message; couldn't order premises")
	}
	return linkedPremises, nil
}

func (ps Set) Len() int {
	return len(ps)
}

// Swap items at the given positions.
func (ps Set) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

// Less determines sorting order between items.
func (ps Set) Less(i, j int) bool {
	return ps[i].Number < ps[j].Number
}

// Copy shallowly into a new premise set.
func (ps Set) copy() Set {
	c := make(Set, len(ps))
	copy(c, ps)
	return c
}

// Find the index of the first premise that contains the given symbol, or return (-1).
func (ps Set) Find(s *symbol.Symbol, start int) int {
	for i := start; i < len(ps); i++ {
		p := ps[i]

		if p.Subject == s || p.Predicate == s {
			return i
		}
	}
	return (-1)
}

// // Find the index of a premise with a given line number.
// // Find will return (-1) if no matching premises are found.
// func (ps Set) Find(n int) int {
// 	for i, p := range ps {
// 		if p.Number == n {
// 			return i
// 		}
// 	}
// 	return (-1)
// }

// Compute a conclusion.
func (ps Map) Compute(symbol1 *symbol.Symbol, symbol2 *symbol.Symbol) string {
	if len(ps) == 0 {
		return "A is A"
	}

	dc1 := ps.Distribution(symbol1)
	dc2 := ps.Distribution(symbol2)
	// if ps.NegativePremiseCount() == 0 {
	if ps.Negative() == 0 {
		// affirmative conclusion
		// TODO: can we push more of these conditionals inside the symbol type?

		if dc1 > 0 {
			return symbol1.ConclusionForAllIs(symbol2)
		} else if dc2 > 0 {
			return symbol2.ConclusionForAllIs(symbol1)
		} else if symbol1.ArticleType != article.TypeNone || symbol2.ArticleType == article.TypeNone {
			return symbol1.ConclusionForSomeIs(symbol2)
		} else {
			return symbol2.ConclusionForSomeIs(symbol1)
		}

	} else {
		// negative conclusion
		if dc2 == 0 {
			return fmt.Sprintf("Some %s is not %s%s", symbol2.Term, symbol1.ArticleType, symbol1.Term)
		} else if dc1 == 0 {
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

// List output for premises, optionally in distribution-analysis format.
// This may be used for link-order output if the premise set is arranged accordingly.
// TODO: use tabwriter for distribution-analysis format?
func (ps Set) List(analyze bool) error {
	if len(ps) == 0 {
		return fmt.Errorf(help.NoPremises)
	}

	for _, prem := range ps {
		if !analyze {
			fmt.Printf("%d  %s\n", prem.Number, prem.Statement)
		} else {
			fmt.Printf("%d  ", prem.Number)

			if prem.Form != form.AEqualsT && prem.Form != form.ADoesNotEqualT && prem.Predicate.TermType == term.TypeDesignator {
				prem.Form += 2
			}

			if prem.Form.IsParticular() || prem.Form.IsUniversal() {
				fmt.Printf("%s  ", prem.Form.Quantifier())
			}

			fmt.Printf("%s%s%s  %s%s\n", prem.Subject.Term, prem.Form.Subject(), prem.Form.Copula(), prem.Predicate.Term, prem.Form.Predicate())
		}
	}
	return nil
}

// CheckNegative validates the negative premise count.
func (ps Map) CheckNegative() error {
	negativePremiseCount := ps.Negative()
	if negativePremiseCount > 1 {
		return fmt.Errorf("Warning: %d negative premises", negativePremiseCount)
	}

	return nil
}

// CheckOccurrences validates the occurrence count of a symbol.
func (ps Map) CheckOccurrences(s *symbol.Symbol) error {
	o := ps.Occurrences(s)
	if o > 2 {
		return fmt.Errorf("Warning: %s %q has occurred %d times", s.TermType, s.Term, o)
	} else if o == 2 && ps.Distribution(s) == 0 {
		return fmt.Errorf("Warning: undistributed middle term %q", s.Term)
	}

	return nil
}

// Occurrences of a particular symbol.
// TODO: this should track symbol.Occurrences perfectly.
func (ps Map) Occurrences(s *symbol.Symbol) int {
	var n int
	for _, p := range ps {
		if p.Subject == s {
			n++
		}

		if p.Predicate == s {
			n++
		}
	}
	return n
}

// Distribution count for a particular symbol.
func (ps Map) Distribution(s *symbol.Symbol) int {
	var n int

	// if s.TermType == term.TypeDesignator {
	// 	n = ps.Occurrences(s)
	// }

	for _, p := range ps {
		// TODO: check universality rather than negate particularity?
		if p.Subject == s {
			// The subject of a universal statement is distributed
			// Source: <http://logiccurriculum.com/2015/10/28/distributed-terms/>
			if !p.Form.IsParticular() {
				n++
			}
		}

		if p.Predicate == s {
			// TODO: Decrement() reduced if predicate had a designator type; is that what we want instead?
			// or does d1 += 2 in syllogism.go mean that unless we have a designator==designator or negative premise,
			// we're not going to increase distribution for predicate?
			// In short: TODO: should this be p.Predicate.TermType == term.TypeDesignator?

			if p.Form.IsNegative() {
				// The predicate of negative statements is distributed.
				// Source: <http://logiccurriculum.com/2015/10/28/distributed-terms/>
				n++
			} else if p.Form == form.AEqualsT {
				n++
			}
		}
	}

	return n
}

// Negative premise count.
// TODO: this could return the slice of negative premises, and then we can do len() on that
func (ps Map) Negative() int {
	var n int
	for _, p := range ps {
		if p.Form.IsNegative() {
			n++
		}
	}
	return n
}
