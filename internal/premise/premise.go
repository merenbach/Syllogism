package premise

import (
	"fmt"
	"strings"

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

// NumberLen counts the digits in a base 10 integer.
func numberLen(n int) int {
	var i int
	for n != 0 {
		n /= 10
		i++
	}
	return i
}

// New premise.
func New(s string) (*Premise, error) {
	var lineNumber int
	if _, err := fmt.Sscanf(s, "%d", &lineNumber); err != nil {
		return nil, err
	}

	stmt := s[numberLen(lineNumber):]
	return &Premise{
		Number:    lineNumber,
		Statement: strings.TrimSpace(stmt),
	}, nil
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

func (pr *Premise) String() string {
	return fmt.Sprintf("%d %s", pr.Number, pr.Statement)
}
