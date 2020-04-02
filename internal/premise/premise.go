package premise

import (
	"fmt"

	"github.com/merenbach/syllogism/internal/form"
	"github.com/merenbach/syllogism/internal/symbol"
	"github.com/merenbach/syllogism/internal/term"
)

// Premise stores a line number and program statement.
// TODO: store numbers with PremiseSet if possible
type Premise struct {
	Number    int
	Statement string
	Form      form.Form
	Subject   *symbol.Symbol
	Predicate *symbol.Symbol
}

// New premise.
// TODO: accept string here and tokenize afterward
func New(n int, s string) *Premise {
	return &Premise{
		Number:    n,
		Statement: s,
	}
}

func (pr *Premise) String() string {
	return fmt.Sprintf("%d %s", pr.Number, pr.Statement)
}

// Decrement table entries.
func (pr *Premise) Decrement() {
	var (
		pDecrement bool
		qDecrement bool
	)

	if pr.Form.IsNegative() || pr.Predicate.TermType == term.TypeDesignator {
		qDecrement = true
	}

	if pr.Form >= 2 {
		pDecrement = true
	}

	pr.Subject.ReduceDistributionCount(pDecrement)
	pr.Predicate.ReduceDistributionCount(qDecrement)
}

// Empty determines whether a line is empty.
func (pr *Premise) Empty() bool {
	return pr.Statement == ""
}
