package premise

import (
	"fmt"

	"github.com/merenbach/syllogism/internal/form"
	"github.com/merenbach/syllogism/internal/symbol"
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

// ContrastingTerm returns the opposite term in a premise.
func (pr *Premise) ContrastingTerm(s *symbol.Symbol) *symbol.Symbol {
	if pr.Subject == s {
		return pr.Predicate
	} else if pr.Predicate == s {
		return pr.Subject
	}

	return nil
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

// Empty determines whether a line is empty.
func (pr *Premise) Empty() bool {
	return pr.Statement == ""
}
