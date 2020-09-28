package premise

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/merenbach/syllogism/internal/article"
	"github.com/merenbach/syllogism/internal/form"
	"github.com/merenbach/syllogism/internal/help"
	"github.com/merenbach/syllogism/internal/symbol"
	"github.com/merenbach/syllogism/internal/term"
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

func testForm(s string, t3 token.Type, t4 token.Type, t5 token.Type, t6 token.Type) (bool, error) {
	switch {
	case t3 != token.TypeTerm:
		return false, errors.New(help.MissingSubject)
	case t4 != token.TypeCopula:
		return false, errors.New(help.MissingCopula)
	case s == form.WordNot:
		// Used for particular forms ending in "NOT B"--we need to make sure that B is a term now
		if t6 != token.TypeUndetermined {
			if t6 != token.TypeTerm {
				return false, errors.New(help.MissingPredicate)
			}
			return true, nil
		}
		fallthrough
	case t5 != token.TypeTerm:
		return false, errors.New(help.MissingPredicate)
	default:
		return false, nil
	}
}

// PremiseForm determines the form of a premise. Pass a function to set new subject and predicate.
func PremiseForm(stringarray_s []string, intarray_t []token.Type, f func(string, string)) (form.Form, error) {
	if stringarray_s[2] == form.WordAll {
		if _, err := testForm(stringarray_s[5], intarray_t[3], intarray_t[4], intarray_t[5], token.TypeUndetermined); err != nil {
			return form.Undefined, err
		}
		f(stringarray_s[3], stringarray_s[5])
		return form.AllAIsB, nil // all A is B
	}

	if stringarray_s[2] == form.WordSome {
		if negative, err := testForm(stringarray_s[5], intarray_t[3], intarray_t[4], intarray_t[5], intarray_t[6]); err != nil {
			return form.Undefined, err
		} else if negative {
			f(stringarray_s[3], stringarray_s[6])
			return form.SomeAIsNotB, nil // some A is not B
		}
		f(stringarray_s[3], stringarray_s[5])
		return form.SomeAIsB, nil // Some A is B
	}

	if stringarray_s[2] == form.WordNo {
		if _, err := testForm(stringarray_s[5], intarray_t[3], intarray_t[4], intarray_t[5], token.TypeUndetermined); err != nil {
			return form.Undefined, err
		}
		f(stringarray_s[3], stringarray_s[5])
		return form.NoAIsB, nil // no A is B
	}

	if negative, err := testForm(stringarray_s[4], intarray_t[2], intarray_t[3], intarray_t[4], intarray_t[5]); err != nil {
		return form.Undefined, err
	} else if negative {
		f(stringarray_s[2], stringarray_s[5])
		return form.AIsNotT, nil // a is not T
	}
	f(stringarray_s[2], stringarray_s[4])
	return form.AIsT, nil // a is T
}

// BasicTab prints tabs in the manner of BASIC's TAB(N)
// This is duplicated from syllogism.go. TODO: REMOVE
func basicTabString(n int) string {
	return strings.Repeat(" ", n)
}

// Tokenize a string.
func Tokenize(localstring_l1 string) ([7]string, [8]token.Type, [3]article.Type, term.Type, error) {
	// 2020
	//---L1$ into array S$()---

	// T(): 1:line num., 2:"/", 3:quantifier, 4:no/not, 5:is/are, 6:term
	//                     10 SOME  FRIED COCONUTS   ARE  NOT  TASTY
	//                      1   3        6            5    4     6

	// TODO: try to return when we need to return
	//       or consider no-ops: https://blog.golang.org/error-handling-and-go
	var returnErr error

	var localstring_s string
	var shadowstringarray_s [7]string
	var shadowintarray_t [8]token.Type
	var shadowintarray_e [3]article.Type

	localint_p1 := term.TypeUndetermined
	shadowintarray_e[2] = article.TypeNone
	localint_j := 1

	localint_l := len(localstring_l1)

	nextToken := func() {
		localint_j++
	}
	setToken := func(s string, t token.Type) {
		shadowstringarray_s[localint_j] = s
		shadowintarray_t[localint_j] = t
	}
	addTermToken := func(s string) {
		setToken(s, token.TypeTerm)
	}
	closeToken := func(s string, t token.Type) {
		setToken(s, t)
		nextToken()
	}

	line2670 := func(tabCount int) { // 2670
		returnErr = fmt.Errorf("%s^\nReserved word %q may not occur within a term", basicTabString(tabCount), localstring_s)
		shadowintarray_t[1] = token.TypeReserved
	}

	// TODO: use strings.Fields here; but need to properly increase letter count
	var localint_i int
Iterate:
	for _, word := range strings.Split(localstring_l1, " ") {
		localint_i++

		if word == "" {
			continue
		}

		// find beginning of next word, skipping any spaces
		localstring_s = word
		localint_k := len(localstring_s)

		if localint_j > 1 {
			goto Line2520
		}

		if localstring_s == "/" {
			closeToken(localstring_s, token.TypeSlash)
		} else {
			if _, err := strconv.Atoi(localstring_s); err != nil {
				returnErr = fmt.Errorf("%s^   Invalid numeral or command", basicTabString(localint_i+len(localstring_s)))
				break
			}
			closeToken(localstring_s, token.TypeLineNumber)
		}
		goto Line2860

	Line2520: // 2520
		// Scan

		/* General overview:
		1. If word is somebody/something/nobody/nothing/someone/everyone/everybody/everything, raise reserved word.
		2. If word is all/some, append as type3 quantifier token if current token isn't a term, then increase counter; otherwise, raise reserved word.
		3. If word is no/not, append as type4 no/not token if current token isn't a term, then increase counter; otherwise, raise reserved word.
		4. If word is is/are, raise error if current token is NOT a term or if either of previous two tokens include is/are; otherwise, increase counter to close current term, THEN append as type5 is/are token as next term, then increase counter.                                                                                                                                                                                                                            5. If current token is term, append space and then word.  Do NOT increment counter.
		6. If prev two tokens do NOT include is/are, then append this word as a type6 term token.  Do NOT increment counter.
		7. If word is "the," set flag P1=2 (designator/definite article) and append as type6 term token.  Do NOT increment counter.
		8. If word is NOT any of a/an/sm, append as type6 term token.  Do NOT increment counter.
		9. If this is the LAST token, append as type6 term token.  Do NOT increment counter.
		10. Set E[2] = index of type of article (a/an/sm) and set flag P1=1 (general term/indefinite article), then and append as type6 term token.  Do NOT increment counter.
		11. If number of tokens > 6, end processing.
		*/

		switch localstring_s {
		case "somebody", "something", "nobody", "nothing":
			line2670(localint_i + localint_k - 1)
			break Iterate

		case "someone", "everyone", "everybody", "everything":
			line2670(localint_i + localint_k - 1)
			break Iterate

		case form.WordAll, form.WordSome:
			if shadowintarray_t[localint_j] == token.TypeTerm {
				line2670(localint_i + localint_k - 1)
				break Iterate
			}
			closeToken(localstring_s, token.TypeQuantifier)

		case form.WordNo, form.WordNot:
			if shadowintarray_t[localint_j] == token.TypeTerm {
				line2670(localint_i + localint_k - 1)
				break Iterate
			}
			closeToken(localstring_s, token.TypeNegation)

		case "is", "are":
			if shadowintarray_t[localint_j] != token.TypeTerm {
				line2670(localint_i + localint_k - 1)
				break Iterate
			} else if shadowintarray_t[localint_j-1] == token.TypeCopula || shadowintarray_t[localint_j-2] == token.TypeCopula {
				line2670(localint_i + localint_k - 1)
				break Iterate
			}
			// NOTE: This is needed here, and not above, because of positioning of to-be verbs in lines.
			// All/some and no/not will occur either at the beginning of the line or right after a to-be verb.
			// To-be verbs are the only words to (legitimately) show up right after the end of a term token (type 6).
			nextToken()
			closeToken(localstring_s, token.TypeCopula)

		default:
			if shadowintarray_t[localint_j] == token.TypeTerm {
				shadowstringarray_s[localint_j] += " " + localstring_s

			} else if shadowintarray_t[localint_j-1] != token.TypeCopula && shadowintarray_t[localint_j-2] != token.TypeCopula {
				addTermToken(localstring_s)

			} else if localstring_s != article.WordA && localstring_s != article.WordAn && localstring_s != article.WordSm {
				if localstring_s == article.WordThe {
					// DESIGNATOR (definite article)
					localint_p1 = term.TypeDesignator
				}
				addTermToken(localstring_s)

			} else if localint_i == localint_l {
				addTermToken(localstring_s)

			} else {
				shadowintarray_e[2] = article.TypeFromString(localstring_s)
				// GENERAL TERM (indefinite article)
				localint_p1 = term.TypeGeneralTerm
			}
		}
		goto Line2860

	Line2860: // 2860
		localint_i += len(word)
		if localint_j > 6 {
			break
		}
	}

	return shadowstringarray_s, shadowintarray_t, shadowintarray_e, localint_p1, returnErr
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
