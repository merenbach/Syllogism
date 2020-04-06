package main

/* TODO:
* Ensure printing output matches original program--static checking doesn't make this any easier.
* Ensure input/output matches with various unpredictable inputs.
* Fix output alignment to either more closely match, or improve upon, original
* Refactor into proper Golang project layout
* Use Golang list type for Line collection?
* Improve dumping further
* Remove any spurious whitespace in entries
*
* Testing notes:
* - Can simulate closed loop with sample, removing line 20, and then setting line 10 to "all men are mortal" (duplicating line 70)
*
* Porting notes on variables:
*
* a(63)  => unknown, but for premises; currently a(N) => premiseset.AArray(N)
* a$(3)  => article type names
* a1     => address of recently-entered line in the list of lines (???)
* b(63)  => term article type (index in a$ of proper article), so anywhere we see b(N) => symbols(N).ArticleType
* b1     => first unused location in symbol table after a particular starting point
			(first slot with symbols(N).Occurrences == 0)
* c      => count of conclusion terms (those that appear only once), used in determination of whether premises form a (poly)syllogism
* c(63)  => term indices of non-middle terms (that is, major and minor terms); currently symbols [ c(N) ] => symbolConclusionTerms(N)
* d(63)  => term distribution count, so anywhere we see d(N) => symbols(N).DistributionCount
* d1     => form of most recently entered premise, either for entry into l$ or for evaluation with /
* e(2)   => article type (index in a$ of article type)
* g      => term type as integer
* g(63)  => term type (index in g$ of term type), so anywhere we see g(N) => symbols(N).TermType
* g$(2)  => term type names
* g1     => term type as integer (termType1)
* g2     => term type as integer (termType2)
* k(63)  => linking order??? (TODO: figure this out), currently premises.LinkOrder
* i1     => local iterator index that is passed through different functions
			appears independent in substitution routine, but spans gosubs 3400 and 3950.
* j(4)   => appears to have held subject and predicate indices in symbol table for decrementing
* j1     => appears to be error code/status for syllogism (???)
* l(63) => occupied line slots (???) in premise list; currently l(N) => premiseset.LArray(N)
* l$(63) => line statements
* l1$    => TBD
* l2$    => TBD
* l1     => highest symbol table location used, so symbolTable.HighestLocationUsed
* n(63)  => line numbers
* n1     => negative premise count on premise set, supplanted by negativePremiseCount
* o(63)  => term occurrence count, so anywhere we see o(N) => symbols(N).Occurrences
* p(63)  => index of subject in symbol table for premise at given index, currently premises(N).Symbol
* p1     => term type as integer
* q(63)  => index of predicate in symbol table for premise at given index, currently premises(N).Symbol
* r(63)  => forms for each premise, currently premises(N).Form
* s$     => TBD
* s$(6)  => parsed line tokens
* t(7)   => token type in parsing
* t1     => major (or minor??) premise in conclusion
* t2     => minor (or major??) premise in conclusion
* t$(65) => term array: words like man, socrates, etc., so anywhere we see t$(N) => symbols(N).Term
* v1     => flag for modern validity (index in symbol table of term involved)
* w$     => most recently entered premise, either for entry into l$ or for evaluation with /
* x$(7)  => quantifiers for each form
* y$(7)  => term A types for each form, followed by copulas for each form
* z$     => conclusion (during computation)
* z$(7)  => term B types for each form
*/
import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/merenbach/syllogism/internal/article"
	"github.com/merenbach/syllogism/internal/form"
	"github.com/merenbach/syllogism/internal/help"
	"github.com/merenbach/syllogism/internal/premise"
	"github.com/merenbach/syllogism/internal/stringutil"
	"github.com/merenbach/syllogism/internal/symbol"
	"github.com/merenbach/syllogism/internal/term"
	"github.com/merenbach/syllogism/internal/token"
)

const basicDimMax = 64

var (
	intarray_t [8]token.Type
	intarray_e [3]article.Type // TODO: about ready to redefine locally where used

	premiseSet  = make(premise.Set, 0)
	symbolTable = make(symbol.Table, 0)
	// SymbolConclusionTerms are major and minor (i.e., all the non-middle) terms
	symbolConclusionTerms = make([]*symbol.Symbol, basicDimMax)
	negativePremiseCount  = 0 // kludge because we aren't able to tally dynamically yet (chicken-and-egg)

	stringarray_s [7]string // appears to hold parsed line tokens
	recentWord1   string    // most recently-input first word (subject?)
	recentWord2   string    // most recently-input second word (predicate?)

	localint_i  int
	localint_j  int
	localint_j1 int
	localint_k  int
	localint_l  int
	localint_p1 term.Type

	localsymbol_v1 *symbol.Symbol

	msg bool
)

// Substitute terms
func substitute() error {
	if len(premiseSet) == 0 {
		return fmt.Errorf(help.NoPremises)
	}

	for {
		fmt.Println("Enter address of old term; or 0 for help, -1 to exit, -2 for dump")

		i1temp_string := lineInput("? ")
		localint_i1, err := strconv.Atoi(i1temp_string)
		if err != nil {
			continue
		}
		if localint_i1 == -1 {
			break
		}
		if localint_i1 != -2 {
			if localint_i1 > 0 {
				if localint_i1 <= len(symbolTable) {
					fmt.Printf("Enter new term to replace %s %q\n", symbolTable[localint_i1-1].TermType, symbolTable[localint_i1-1].Term)

					w := lineInput("? ")
					symbolTable[localint_i1-1].Term = w
					fmt.Printf("Replaced by %q\n", w)
				} else {
					fmt.Printf("Address %d too large.  Symbol table only of length %d.\n", localint_i1, len(symbolTable))
				}
			} else {
				fmt.Println(help.SyllogismHelpForSubstitute)
			}
			fmt.Println()
		} else {
			fmt.Println(Dump())
		}
	}

	fmt.Println("Exit from substitution routine")
	return nil
}

func basicGosub5880() {
	// 5880
	//---See if conclusion possible---

	symbol1 := symbolConclusionTerms[1]
	symbol2 := symbolConclusionTerms[2]

	for i, s := range symbolTable {
		if s.Occurrences < 2 {
			continue
		}

		if s.DistributionCount == 0 {
			if localint_j1 == 0 {
				fmt.Println("Undistributed middle terms:")
				localint_j1 = 5
			}

			fmt.Printf("%s%s\n", basicTabString(5), s.Term)
		}

		if s.DistributionCount != 1 && s.TermType != term.TypeDesignator {
			localsymbol_v1 = symbolTable[i]
		}
	}

	if negativePremiseCount > 1 {
		localint_j1 = 6
		fmt.Println("More than one negative premise:")
	}

	if localint_j1 > 0 {
		fmt.Println(help.NoPossibleConclusion)
		return
	}

	if negativePremiseCount == 0 {
		return
	}

	if symbol1.DistributionCount == 0 && symbol2.DistributionCount == 0 {
		fmt.Printf("Terms %q and %q, one of which is\n", symbol1.Term, symbol2.Term)
	} else if symbol1.DistributionCount == 0 && symbol2.TermType == term.TypeDesignator {
		fmt.Printf("Term %q\n", symbol1.Term)
	} else if symbol2.DistributionCount == 0 && symbol1.TermType == term.TypeDesignator {
		fmt.Printf("Term %q\n", symbol2.Term)
	} else {
		return
	}

	fmt.Println("required in predicate of negative conclusion")
	fmt.Println("not distributed in the premises.")
	localint_j1 = 7

	fmt.Println(help.NoPossibleConclusion)
}

func basicGosub5070() premise.Set {
	// 5070
	//---See if syllogism---
	var temp_symbol *symbol.Symbol

	localint_j1 = 0
	localsymbol_v1 = nil // flag for modern validity
	if len(premiseSet) == 0 {
		localint_j1 = 1
		return nil
	}

	var localint_c int

	for i, s := range symbolTable {
		if s.Occurrences == 0 || s.Occurrences == 2 {
			continue
		}

		if s.Occurrences == 1 {
			localint_c++
			symbolConclusionTerms[localint_c] = symbolTable[i]
			continue
		}

		if localint_j1 != 2 {
			fmt.Println("Not a syllogism:")
			localint_j1 = 2
		}

		fmt.Printf("   %s %q occurs %d times in premises.\n", s.TermType, s.Term, s.Occurrences)
	}

	if localint_c != 2 {
		fmt.Println("Not a syllogism:")
		localint_j1 = 3

		if localint_c > 0 {
			fmt.Printf("   %d  terms occur exactly once in premises.\n", localint_c)

			for i := 1; i <= localint_c; i++ {
				// TODO: use tabwriter here?
				sym := symbolConclusionTerms[i]
				fmt.Printf("%s%s -- %s\n", basicTabString(6), sym.Term, sym.TermType)
			}
		} else {
			fmt.Println("   no terms occur exactly once in premises.")
		}
	}

	if localint_j1 != 0 {
		return nil
	}

	linkedPremises := make(premise.Set, len(premiseSet))
	copy(linkedPremises, premiseSet)
	localint_l = len(premiseSet)

	if localint_l == 1 {
		goto Line5750
	}

	if symbolConclusionTerms[1].DistributionCount == 0 && symbolConclusionTerms[2].DistributionCount == 1 {
		temp_symbol = symbolConclusionTerms[2]
	} else {
		temp_symbol = symbolConclusionTerms[1]
	}
	localint_i = 0

Line5460: // 5460
	localint_k = localint_i

Line5470: // 5470
	if linkedPremises[localint_k].Subject == temp_symbol {
		temp_symbol = linkedPremises[localint_k].Predicate
	} else if linkedPremises[localint_k].Predicate == temp_symbol {
		temp_symbol = linkedPremises[localint_k].Subject
	} else {
		localint_k++
		if localint_k < localint_l {
			goto Line5470
		}

		temp_symbol = linkedPremises[localint_i].Predicate

		if localint_j1 == 0 {
			localint_j1 = 4
			fmt.Println("Not a syllogism: no way to order premises so that each premise")
			fmt.Println("shares exactly one term with its successor; there is a")
		}
		fmt.Println(help.ClosedLoopHelp)
		fmt.Println(linkedPremises[localint_i])
		goto Line5730
	}

	if localint_k != localint_i {
		// Swap premises at locations k and i
		linkedPremises[localint_k], linkedPremises[localint_i] = linkedPremises[localint_i], linkedPremises[localint_k]
	}

	if localint_j1 != 0 {
		fmt.Println(linkedPremises[localint_i])
	}

Line5730: // 5730
	localint_i++

	if localint_i < localint_l {
		goto Line5460
	}

Line5750: // 5750
	if localint_j1 > 0 {
		return nil
	}

	return linkedPremises
}

func basicGosub6200() {
	// 6200
	//---Compute conclusion---
	z := premiseSet.Compute(negativePremiseCount, symbolConclusionTerms[1], symbolConclusionTerms[2])

	// PRINT  conclusion
	fmt.Printf("  / %s\n", z)
	if localsymbol_v1 != nil {
		fmt.Print("  * Aristotle-valid only, i.e. on requirement that term ")
		fmt.Printf("%q denotes.\n", localsymbol_v1.Term)
	}
}

func basicGosub6630(p1 term.Type) {
	// 6630
	//---test offered conclusion---
	var (
		localstring_w  string
		localsymbol_t1 *symbol.Symbol
		localsymbol_t2 *symbol.Symbol
	)
	var termType1 term.Type = term.TypeGeneralTerm // formerly g1
	var termType2 term.Type = term.TypeGeneralTerm // formerly g2

	//--conc. poss, line in s$()
	d1, err := basicGosub2890()
	if err != nil {
		fmt.Println(err)
		if msg {
			fmt.Println("Enter SYNTAX for help with statements")
		}
	}

	if d1 == form.Undefined {
		return
	}

	if d1 == form.AIsT || d1 == form.AIsNotT || d1 == form.AEqualsT || d1 == form.ADoesNotEqualT {
		termType1 = term.TypeDesignator
		termType2 = p1
	}
	if termType2 == term.TypeDesignator {
		switch d1 {
		case form.AIsT:
			d1 = form.AEqualsT
		case form.AIsNotT:
			d1 = form.ADoesNotEqualT
		}
	}

	localstring_w = stringutil.Singularize(recentWord1)
	if localint_j1 != 0 {
		recentWord1 = localstring_w
	} else {
		localint_j = 1
		if !symbolConclusionTerms[localint_j].MatchesWordAndTermType(localstring_w, termType1) {
			localint_j = 2
			if !symbolConclusionTerms[localint_j].MatchesWordAndTermType(localstring_w, termType1) {
				fmt.Printf("** Conclusion may not contain %s %q.\n", termType1, localstring_w)
				localint_j = 0
			}
		}
	}

	localstring_w = stringutil.Singularize(recentWord2)
	if localint_j1 != 0 {
		if localstring_w != recentWord1 {
			fmt.Println(help.ConclusionFromNoPremises)
			return
		} else if d1 == form.AIsT && termType2 != term.TypeUndetermined {
			fmt.Printf("** Subject is a %s, predicate is a %s -- but\n", term.TypeDesignator, term.TypeGeneralTerm)
			fmt.Println(help.ConclusionFromNoPremises)
			return
		} else {
			goto Line7120
		}
	}

	if localint_j > 0 {
		localsymbol_t1 = symbolConclusionTerms[localint_j]
		localsymbol_t2 = symbolConclusionTerms[3-localint_j]
		if localstring_w != localsymbol_t2.Term {
			goto Line7060
		}
		if localsymbol_t2.TermType != term.TypeUndetermined {
			if termType2 != term.TypeUndetermined && termType2 != localsymbol_t2.TermType {
				goto Line7060
			}
		} else if termType2 != term.TypeUndetermined {
			fmt.Printf("Note: %q used in premises taken to be %s\n", localsymbol_t2.Term, termType2)
		}
		if negativePremiseCount > 0 && !d1.IsNegative() {
			fmt.Println("** Negative conclusion required.")
			return
		}
		goto Line7120
	}
	if localstring_w == symbolConclusionTerms[1].Term {
		localsymbol_t2 = symbolConclusionTerms[2]
	} else {
		localsymbol_t2 = symbolConclusionTerms[1]
	}
	goto Line7070

Line7060: // 7060
	fmt.Printf("** Conclusion may not contain %s %q;\n", termType2, localstring_w)

Line7070: // 7070
	fmt.Printf("** Conclusion must contain %s %q.\n", localsymbol_t2.TermType, localsymbol_t2.Term)
	return

Line7120: // 7120
	if negativePremiseCount == 0 && d1.IsNegative() {
		fmt.Println("** Affirmative conclusion required.")
		return
	}

	if localint_j1 != 1 {
		if localsymbol_t1.DistributionCount == 0 && (d1 == form.AllAIsB || d1 == form.NoAIsB) {
			help.ShowTermDistributionError(localsymbol_t1.Term)
			return
		} else if localsymbol_t2.DistributionCount == 0 && (d1.IsNegative() || d1 == form.AEqualsT) {
			help.ShowTermDistributionError(localsymbol_t2.Term)
			return
		}
	}

	fmt.Println("-->  VALID!")

	if localint_j1 != 0 {
		if d1 != form.SomeAIsB {
			return
		}
		localsymbol_v1 = &symbol.Symbol{
			Term: localstring_w,
		}
	} else if localsymbol_t1.DistributionCount > 0 && (d1 == form.SomeAIsB || d1 == form.SomeAIsNotB) {
		localsymbol_v1 = localsymbol_t1
	} else {
		if localsymbol_t2.DistributionCount > 0 && !d1.IsNegative() && d1 != form.AIsT && d1 != form.AEqualsT {
			localsymbol_v1 = localsymbol_t2
		}

		if localsymbol_v1 == nil {
			return
		}
	}

	fmt.Println("    but on Aristotelian interpretation only, i.e. on requirement")
	fmt.Printf("    that term %q denotes.\n", localsymbol_v1.Term)
}

func basicGosub1840() {
	// 1840
	//---New---

	if len(premiseSet) == 0 {
		return
	}

	symbolConclusionTerms = make([]*symbol.Symbol, basicDimMax)
	premiseSet = make(premise.Set, 0)
	symbolTable = make(symbol.Table, 0)
	negativePremiseCount = 0
}

func basicGosub3400(d1 form.Form, p1 term.Type, prem *premise.Premise) {
	// 3400
	//---Add W$(1), W$(2) to table T$()---
	// d1 is guaranteed not to be form.Undefined unless `sample` method isn't working (TODO: funnel sample through same logic as user input)
	var termType term.Type // formerly g
	if d1.IsNegative() {
		negativePremiseCount++
		if negativePremiseCount > 1 && msg {
			fmt.Printf("Warning: %d negative premises\n", negativePremiseCount)
		}
	}

	intarray_e[1] = article.TypeNone
	temp := func(localint_j int, raw_string string) {
		if d1 == form.SomeAIsB || d1 == form.SomeAIsNotB || d1 == form.AllAIsB || d1 == form.NoAIsB {
			termType = term.TypeGeneralTerm
		} else if localint_j == 1 {
			termType = term.TypeDesignator
		} else {
			termType = p1
		}

		w := stringutil.Singularize(raw_string)

		sym := func() *symbol.Symbol {
			var localint_i1 int

			symbolTable = Prune(symbolTable)
			for ; ; localint_i1++ { // 3500
				localint_i1 = symbolTable.Search(localint_i1, w)

				if localint_i1 == len(symbolTable) {
					symbolTable = append(symbolTable, &symbol.Symbol{
						Term:     w,
						TermType: termType,
					})
					break
				}

				sym := symbolTable[localint_i1]
				if termType == term.TypeUndetermined {
					if sym.TermType != term.TypeUndetermined || msg {
						fmt.Printf("Note: predicate term %q", w)
						fmt.Printf(" taken as the %s used earlier\n", sym.TermType)
					}
					break
				} else if sym.TermType == term.TypeUndetermined {
					if msg {
						fmt.Printf("Note: earlier use of %q taken as the %s used here\n", w, termType)
					}
					if termType == term.TypeDesignator {
						sym.DistributionCount = sym.Occurrences
					}
					sym.TermType = termType
					break
				} else if termType == sym.TermType {
					break
				} else if msg {
					fmt.Printf("Warning: %s %q has also occurred as a %s\n", termType, w, termType.Other())
				}
			}

			return symbolTable[localint_i1]
		}()

		if intarray_e[localint_j] != article.TypeNone {
			sym.ArticleType = intarray_e[localint_j]
		} else if sym.ArticleType == article.TypeNone && w != raw_string {
			if stringutil.HasPrefixVowel(w) {
				// AN
				intarray_e[localint_j] = article.TypeAn
			} else {
				// A
				intarray_e[localint_j] = article.TypeA
			}
			sym.ArticleType = intarray_e[localint_j]
		}

		sym.Occurrences++

		if sym.Occurrences >= 3 && msg {
			fmt.Printf("Warning: %s %q has occurred %d times\n", sym.TermType, w, sym.Occurrences)
		}

		if localint_j != 2 {
			prem.Subject = sym

			if d1 != form.SomeAIsB && d1 != form.SomeAIsNotB {
				sym.DistributionCount++
			}

		} else {
			prem.Predicate = sym

			if prem.Subject == prem.Predicate && msg {
				fmt.Printf("Warning: same term occurs twice in line %s\n", stringarray_s[1])
			}

			if sym.TermType == term.TypeDesignator {
				d1 += 2
			}

			if d1 == form.AEqualsT || d1.IsNegative() {
				sym.DistributionCount++
			}
		}

		if sym.Occurrences == 2 && sym.DistributionCount == 0 && msg {
			fmt.Printf("Warning: undistributed middle term %q\n", sym.Term)
		}
	}

	localint_j = 1
	temp(localint_j, recentWord1)
	localint_j = 2
	temp(localint_j, recentWord2)

	prem.Form = d1
}

// type formInfo struct {
// 	formType   int
// 	firstWord  string
// 	secondWord string
// }

func basicGosub2890() (form.Form, error) {
	// 2890
	//---Parse line in S$()---
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

					recentWord1 = stringarray_s[2]
					recentWord2 = stringarray_s[5]
					return form.AIsNotT, nil // a is not T
				} else {
					if intarray_t[4] != token.TypeTerm {
						return form.Undefined, errors.New(help.MissingPredicate)
					}

					recentWord1 = stringarray_s[2]
					recentWord2 = stringarray_s[4]
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

			recentWord1 = stringarray_s[3]
			recentWord2 = stringarray_s[5]

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
			recentWord1 = stringarray_s[3]
			recentWord2 = stringarray_s[6]
			return form.SomeAIsNotB, nil // some A is not B
		}
		if intarray_t[5] != token.TypeTerm {
			return form.Undefined, errors.New(help.MissingPredicate)
		}
		recentWord1 = stringarray_s[3]
		recentWord2 = stringarray_s[5]
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
	recentWord1 = stringarray_s[3]
	recentWord2 = stringarray_s[5]
	return form.AllAIsB, nil // all A is B
}

func tokenize(localstring_l1 string) ([7]string, [8]token.Type, [3]article.Type, error) {
	// 2020
	//---L1$ into array S$()---

	// T(): 1:line num., 2:"/", 3:quantifier, 4:no/not, 5:is/are, 6:term
	//                     10 SOME  FRIED COCONUTS   ARE  NOT  TASTY
	//                      1   3        6            5    4     6

	// TODO: try to return when we need to return
	var returnErr error

	var localstring_s string
	var shadowstringarray_s [7]string
	var shadowintarray_t [8]token.Type
	var shadowintarray_e [3]article.Type

	localint_p1 = term.TypeUndetermined
	shadowintarray_e[2] = article.TypeNone
	localint_j = 1

	localint_l = len(localstring_l1)

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
	localint_i = 0
Iterate:
	for _, word := range strings.Split(localstring_l1, " ") {
		localint_i++

		if word == "" {
			continue
		}

		// find beginning of next word, skipping any spaces
		localstring_s = word
		localint_k = len(localstring_s)

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

	return shadowstringarray_s, shadowintarray_t, shadowintarray_e, returnErr
}

func basicGosub8980() {
	// 8980
	//--sample--

	// TODO: factor into more local scope for conditionals
	var err error
	for _, localstring_l1 := range strings.Split(help.SampleData, "\n") {
		fmt.Println(localstring_l1)
		stringarray_s, intarray_t, intarray_e, err = tokenize(localstring_l1)
		if err != nil {
			fmt.Println(err)
		}
		d1, err := basicGosub2890()
		if err != nil {
			fmt.Println(err)
			if msg {
				fmt.Println("Enter SYNTAX for help with statements")
			}
		}

		n, err := strconv.Atoi(stringarray_s[1])
		if err != nil {
			log.Println(err)
		}
		localint_s := len(stringarray_s[1]) + 1
		prem := addPremise(n, localstring_l1[localint_s:])
		basicGosub3400(d1, localint_p1, prem)
	}

	if msg {
		fmt.Println("Suggestion: try the LINK or LINK* command.")
	}
}

// BasicTab prints tabs in the manner of BASIC's TAB(N)
func basicTabString(n int) string {
	return strings.Repeat(" ", n)
}

// BasicLeft emulates LEFT$
func basicLeft(s string, n int) string {
	return string(s[:n])
}

// BasicRight emulates RIGHT$
func basicRight(s string, n int) string {
	return string(s[len(s)-n:])
}

func lineInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	// fmt.Printf("%s ", prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		// TODO: return err instead
		log.Fatal(err)
	}

	// err = f(strings.TrimSpace(input))
	// if err != nil {
	//         fmt.Println(err)
	// }

	return strings.TrimSpace(input)
}

func runloop() bool {
	// TODO: factor into more local scope for conditionals
	var err error

	//---Input line---
	localstring_l1 := lineInput("> ")
	localint_l = len(localstring_l1)

	if localint_l == 0 {
		if msg {
			fmt.Println("Enter HELP for list of commands")
		}
		return true
	}

	for {
		localstring_l2 := basicRight(localstring_l1, 1)
		if localstring_l2 != " " {

			if localstring_l2 != "." && localstring_l2 != "?" && localstring_l2 != "!" {
				break
			}

			fmt.Printf("%s ^   Punctuation mark ignored\n", basicTabString(localint_l))
		}

		if localint_l == 1 {
			return true
		}

		localint_l--
		localstring_l1 = basicLeft(localstring_l1, localint_l)
	}

	for {
		if basicLeft(localstring_l1, 1) != " " {
			break
		}
		if localint_l == 1 {
			return true
		}
		localint_l--
		localstring_l1 = basicRight(localstring_l1, localint_l)
	}

	/*
	   rem / FOR I = 1 TO L
	    rem / V = ASC(MID$(L1$,I,1))
	    rem / IF V >= 65 AND V <= 90  THEN  MID$(L1$,I,1) = CHR$(V+32)
	    rem / NEXT
	   rem Metal doesn't support mid$ as command, but lcase$() is well supported...

	*/

	localstring_l1 = strings.ToLower(localstring_l1)

	switch localstring_l1 {
	case "stop":
		if msg {
			fmt.Println("(Some versions support typing CONT to continue)")
		}
		return false
	case "new":
		fmt.Println("Begin new syllogism")
		basicGosub1840()
		return true
	case "sample":
		basicGosub1840()
		basicGosub8980()
		return true
	case "help":
		help.ShowInputsHelp()
		return true
	case "syntax":
		help.ShowSyntaxHelp()
		return true
	case "info":
		help.ShowGeneralHelp()
		return true
	case "dump":
		fmt.Println(Dump())
		return true
	case "msg":
		msg = !msg

		fmt.Print("Messages turned ")
		if msg {
			fmt.Println("on")
		} else {
			fmt.Println("off")
		}
		return true

	case "substitute":
		if err := substitute(); err != nil {
			fmt.Println(err)
		}
		return true

	case "link":
		fallthrough
	case "link*":
		if len(premiseSet) == 0 {
			fmt.Println(help.NoPremises)
		} else {
			linkedPremises := basicGosub5070()
			if linkedPremises != nil {
				fmt.Println("Premises of syllogism in order of term links:")
				// TODO: did we need anything with localint_l (passed in as `max` before)?
				analyze := strings.HasSuffix(localstring_l1, "*")
				_ = linkedPremises.List(analyze)
			}
		}
		return true
	case "list":
		if err := premiseSet.List(false); err != nil {
			fmt.Println(err)
		}
		return true
	case "list*":
		if err := premiseSet.List(true); err != nil {
			fmt.Println(err)
		}
		return true
	}

	//--scan line L1$ into array S$()

	stringarray_s, intarray_t, intarray_e, err = tokenize(localstring_l1)
	if err != nil {
		fmt.Println(err)
	}
	if intarray_t[1] == token.TypeLineNumber {
		if intarray_t[2] != token.TypeReserved {
			d1, err := basicGosub2890() // parse the line in S$()
			if err != nil {
				fmt.Println(err)
				if msg {
					fmt.Println("Enter SYNTAX for help with statements")
				}
			}
			if d1 != form.Undefined {
				n, err := strconv.Atoi(stringarray_s[1])
				if err != nil {
					log.Println(err)
				}
				localint_s := len(stringarray_s[1]) + 1
				prem := addPremise(n, localstring_l1[localint_s:])
				basicGosub3400(d1, localint_p1, prem) // add terms to symbol table
			}
		} else {
			if len(premiseSet) == 0 {
				fmt.Println(help.NoPremises)
			} else {
				n, err := strconv.Atoi(stringarray_s[1])
				if err != nil {
					log.Println(err)
				}
				if err := delPremise(n); err != nil {
					fmt.Println(err)
				}
			}
		}
		return true
	} else if intarray_t[1] == token.TypeReserved {
		if msg {
			fmt.Println("Enter HELP for list of commands")
		}
		return true
	}

	// draw/test conclusion

	// NOTE: mods j1, checks and mods premiseSet.ConclusionTerms
	_ = basicGosub5070() // is it a syl?
	if localint_j1 > 1 {
		return true
	}
	if localint_j1 == 0 {
		// NOTE: checks and mods j1, checks premiseSet.ConclusionTerms
		basicGosub5880() // poss. conclusion?
	}

	if localint_j1 > 1 {
		return true
	}

	if intarray_t[2] != token.TypeReserved {
		// NOTE: checks j1, checks premiseSet.ConclusionTerms
		basicGosub6630(localint_p1)
	} else {
		basicGosub6200() // test/draw conclusion
	}

	return true
}

func addPremise(n int, s string) *premise.Premise {
	// Delete existing entry instead of replacing in-place
	_ = delPremise(n)

	prem := premise.New(n, s)

	// Append the new premise and sort by line number
	premiseSet = append(premiseSet, prem)
	sort.Slice(premiseSet, func(i, j int) bool {
		return premiseSet[i].Number < premiseSet[j].Number
	})

	return prem
}

func delPremise(n int) error {
	for i, p := range premiseSet {
		if p.Number == n {
			if p.Form.IsNegative() {
				negativePremiseCount--
			}
			p.Decrement()

			// Delete without leaving uncollected pointers
			// https://github.com/golang/go/wiki/SliceTricks
			if i < len(premiseSet)-1 {
				copy(premiseSet[i:], premiseSet[i+1:])
			}
			premiseSet[len(premiseSet)-1] = nil
			premiseSet = premiseSet[:len(premiseSet)-1]

			return nil
		}
	}
	return fmt.Errorf("Line %d not found", n)
}

// Dump values of variables in a SymbolTable.
func Dump() string {
	dump := new(bytes.Buffer)
	fmt.Fprintf(dump, "Highest symbol table loc. used: %d  Negative premises: %d\n", len(symbolTable), negativePremiseCount)
	if len(symbolTable) > 0 {
		w := tabwriter.NewWriter(dump, 0, 0, 2, ' ', 0)
		fmt.Fprint(w, "Adr.\tart.\tterm\ttype\toccurs\tdist. count")
		for i, s := range symbolTable {
			fmt.Fprintf(w, "\n%d\t%s", i+1, s.Dump())
		}
		w.Flush()
	}
	return dump.String()
}

// // Delete a term from the table.
// func (st *SymbolTable) Delete(sym *symbol.Symbol) {
// 	for i, s := range st.Symbols {
// 		if s.Term == sym.Term {
// 			// Delete without leaving uncollected pointers
// 			// https://github.com/golang/go/wiki/SliceTricks
// 			if i < len(st.Symbols)-1 {
// 				copy(st.Symbols[i:], st.Symbols[i+1:])
// 			}
// 			st.Symbols[len(st.Symbols)-1] = nil
// 			st.Symbols = st.Symbols[:len(st.Symbols)-1]
// 			break
// 		}
// 	}
// }

// Prune orphaned terms with no occurrences.
func Prune(st symbol.Table) symbol.Table {
	var ss symbol.Table
	for _, s := range st {
		if s.Occurrences > 0 {
			ss = append(ss, s)
		}
	}

	return ss
}

func main() {
	/*
	   Syllogism 1.0. November 8, 2002
	   I edited this program in 2002, for compatibility with freeware BASIC
	   interpreters for the Mac: Chipmunk BASIC 3.5.7 and Metal BASIC 1.7.3.
	   I hope compatibility with two indpendently developed BASICs should
	   assure some universality regardless of platform. The following
	   notes should help anybody with a similar project.
	   Summary of changes.....
	   * Metal doesn't support PRINT TAB(N). It supports the command HTAB, but a
	   bug makes it useless for formatting more than one column of text. The only
	   standard BASIC solution is the " " character, implemented with tb$ and Left$()
	   * Metal doesn't support MID$() as a command. Both support LCASE$().
	   * Metal crashes when it reads an empty ("") DATA item. Had to hack it.
	   * Metal requires ; between PRINT items. Also, it doesn't add a space
	   when the ; separates a number and a string.
	   Chipmunk buggy with IF...ELSE on one line--use colons (and GOTOs)
	   Chipmunk requires quotes around DATA items
	   Chipmunk buggy when using integer variables (J%), use floating point (J)
	   Neither support PRINT CHR$(27) as a way to clear the screen; used cls
	   Peace. Ben Sharvy. luvnpeas99@yahoo.com
	*/

	/*

	   // 100
	    rem Metal doesn't support PRINT TAB(N)...


	   // 102
	    tb$ = "                                                  "
	*/

	help.ShowCopyright()
	fmt.Println()

	msg = true

	if msg {
		fmt.Println("Enter HELP for list of commands")
	}

	for runloop() {
		// Keep running till runloop() returns false
	}
}
