package premiseset

import (
	"fmt"
	"sort"

	"github.com/merenbach/syllogism/internal/article"
	"github.com/merenbach/syllogism/internal/premise"
	"github.com/merenbach/syllogism/internal/symbol"
	"github.com/merenbach/syllogism/internal/term"
)

// Set of all premises.
type Set []*premise.Premise

// Enter line into list.
func (ps *Set) Enter(n int, s string) *premise.Premise {
	// Silently delete any existing line matching this line number
	_ = ps.Delete(n)
	newPremise := premise.New(n, s)

	// NOTE: for new experimental refactor
	*ps = append(*ps, newPremise)
	ps.Sort()

	return newPremise
}

// Delete a line.
func (ps *Set) Delete(n int) error {
	for i, p := range *ps {
		if p.Number == n {
			p.Decrement()
			*ps = append((*ps)[:i], (*ps)[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Line %d not found", n)
}

// Sort premises by line number.
func (ps *Set) Sort() {
	sort.Slice(*ps, func(i, j int) bool { return (*ps)[i].Number < (*ps)[j].Number })
}

// Compute a conclusion.
func (ps *Set) Compute(symbol1 *symbol.Symbol, symbol2 *symbol.Symbol) string {
	if ps.Empty() {
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

// List output for premises, optionally in distribution-analysis format.
// TODO: use tabwriter for distribution-analysis format?
func (ps *Set) List(analyze bool) {
	for _, prem := range *ps {
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

// NegativePremiseCount returns the count of negative premises.
func (ps *Set) NegativePremiseCount() int {
	var negativePremises int
	for _, p := range *ps {
		if p.Form.IsNegative() {
			negativePremises++
		}
	}
	return negativePremises
}

// Empty determines whether the premise set is empty.
func (ps *Set) Empty() bool {
	return len(*ps) == 0
}
