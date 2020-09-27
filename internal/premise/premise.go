package premise

import (
	"errors"
	"fmt"
	"strings"

	"github.com/merenbach/syllogism/internal/form"
	"github.com/merenbach/syllogism/internal/help"
	"github.com/merenbach/syllogism/internal/symbol"
	"github.com/merenbach/syllogism/internal/token"
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

// PremiseForm determines the form of a premise. This is being transitioned to not needing input.
func PremiseForm(stringarray_s []string, intarray_t []token.Type, recentWord1 *string, recentWord2 *string) (form.Form, error) {
	if stringarray_s[2] != form.WordAll {
		if stringarray_s[2] != form.WordSome {
			if stringarray_s[2] != form.WordNo {
				if intarray_t[2] != token.TypeTerm {
					return form.Undefined, errors.New(help.MissingSubject)
				}

				if intarray_t[3] != token.TypeCopula {
					return form.Undefined, errors.New(help.MissingCopula)
				}

				if stringarray_s[4] == form.WordNot {
					if intarray_t[5] != token.TypeTerm {
						return form.Undefined, errors.New(help.MissingPredicate)
					}

					*recentWord1 = stringarray_s[2]
					*recentWord2 = stringarray_s[5]
					return form.AIsNotT, nil // a is not T
				} else {
					if intarray_t[4] != token.TypeTerm {
						return form.Undefined, errors.New(help.MissingPredicate)
					}

					*recentWord1 = stringarray_s[2]
					*recentWord2 = stringarray_s[4]
					return form.AIsT, nil // a is T
				}
			}

			if intarray_t[3] != token.TypeTerm {
				return form.Undefined, errors.New(help.MissingSubject)
			}

			if intarray_t[4] != token.TypeCopula {
				return form.Undefined, errors.New(help.MissingCopula)
			}

			if intarray_t[5] != token.TypeTerm {
				return form.Undefined, errors.New(help.MissingPredicate)
			}

			*recentWord1 = stringarray_s[3]
			*recentWord2 = stringarray_s[5]

			return form.NoAIsB, nil // no A is B
		}
		if intarray_t[3] != token.TypeTerm {
			return form.Undefined, errors.New(help.MissingSubject)
		}
		if intarray_t[4] != token.TypeCopula {
			return form.Undefined, errors.New(help.MissingCopula)
		}
		if stringarray_s[5] == form.WordNot {
			if intarray_t[6] != token.TypeTerm {
				return form.Undefined, errors.New(help.MissingPredicate)
			}
			*recentWord1 = stringarray_s[3]
			*recentWord2 = stringarray_s[6]
			return form.SomeAIsNotB, nil // some A is not B
		}
		if intarray_t[5] != token.TypeTerm {
			return form.Undefined, errors.New(help.MissingPredicate)
		}
		*recentWord1 = stringarray_s[3]
		*recentWord2 = stringarray_s[5]
		return form.SomeAIsB, nil // Some A is B
	}
	if intarray_t[3] != token.TypeTerm {
		return form.Undefined, errors.New(help.MissingSubject)
	}
	if intarray_t[4] != token.TypeCopula {
		return form.Undefined, errors.New(help.MissingCopula)
	}
	if intarray_t[5] != token.TypeTerm {
		return form.Undefined, errors.New(help.MissingPredicate)
	}
	*recentWord1 = stringarray_s[3]
	*recentWord2 = stringarray_s[5]
	return form.AllAIsB, nil // all A is B
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
