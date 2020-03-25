package main

/* TODO:
* Ensure printing output matches original program--static checking doesn't make this any easier.
* Ensure input/output matches with various unpredictable inputs.
* Fix output alignment to either more closely match, or improve upon, original
* Refactor into proper Golang project layout
* Use Golang list type for Line collection?
* Improve dumping further
*
* Porting notes on variables:
*
* a$(3)  => article type names
* a1     => address of recently-entered line in the list of lines (???)
* b(63)  => term article type (index in a$ of proper article), so anywhere we see b(N) => symbols(N).ArticleType
* b1     => first unused location in symbol table after a particular starting point
            (first slot with symbols(N).Occurrences == 0)
* d(63)  => term distribution count, so anywhere we see d(N) => symbols(N).DistributionCount
* d1     => form of most recently entered premise, either for entry into l$ or for evaluation with /
* e(2)   => article type (index in a$ of article type)
* g      => term type as integer
* g(63)  => term type (index in g$ of term type), so anywhere we see g(N) => symbols(N).TermType
* g$(2)  => term type names
* g1     => term type as integer
* g2     => term type as integer
* k(63)  => linking order??? (TODO: figure this out), currently premises.LinkOrder
* i1     => local iterator index that is passed through different functions
            appears independent in substitution routine, but spans gosubs 3400 and 3950.
* l$(63) => line statements
* l1     => highest symbol table location used, so symboltable.HighestLocationUsed
* n(63)  => line numbers
* n1     => negative premise count on symbol table, so symboltable.NegativePremiseCount
* o(63)  => term occurrence count, so anywhere we see o(N) => symbols(N).Occurrences
* p(63)  => index of subject in symbol table for premise at given index, currently premises(N).Symbol
* p1     => term type as integer
* q(63)  => index of predicate in symbol table for premise at given index, currently premises(N).Symbol
* r(63)  => forms for each premise, currently premises(N).Form
* s$     => parsed line tokens
* t(7)   => token type in parsing
* t$(65) => term array: words like man, socrates, etc., so anywhere we see t$(N) => symbols(N).Term
* v1     => flag for modern validity
* w$     => most recently entered premise, either for entry into l$ or for evaluation with /
* x$(7)  => quantifiers for each form
* y$(7)  => term A types for each form, followed by copulas for each form
* z$     => conclusion (during computation)
* z$(7)  => term B types for each form
*/
import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/merenbach/syllogism/internal/article"
	"github.com/merenbach/syllogism/internal/form"
	"github.com/merenbach/syllogism/internal/help"
	"github.com/merenbach/syllogism/internal/premise"
	"github.com/merenbach/syllogism/internal/stringutil"
	"github.com/merenbach/syllogism/internal/symbol"
	"github.com/merenbach/syllogism/internal/symboltable"
	"github.com/merenbach/syllogism/internal/term"
	"github.com/merenbach/syllogism/internal/token"
)

const basicDimMax = 64

var (
	intarray_a [basicDimMax]int
	intarray_c [basicDimMax]int
	intarray_l [basicDimMax]int

	intarray_t [8]token.Type
	intarray_e [3]article.Type // TODO: about ready to redefine locally where used

	symbolTable = symboltable.New(basicDimMax + 2)

	stringarray_s [7]string // appears to hold parsed line tokens
	stringarray_w [3]string // appears to hold the most recently-input first and second terms for parsing or testing

	localint_t1    int
	localint_t2    int
	localint_c     int
	localint_c1    int
	localint_c2    int
	localint_i     int
	localint_i1    int
	localint_j     int
	localint_j1    int
	localint_k     int
	localint_l     int
	localint_n     int
	localint_p1    term.Type
	localint_s     int
	localint_v1    int
	localstring_l  string
	localstring_l1 string
	localstring_l2 string
	localstring_s  string
	localstring_t  string
	localstring_w  string

	premiseSet = premise.NewPremiseSet(basicDimMax)

	msg bool
)

func basicGosub9060() {
	// 9060
	//---Substitute terms---
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
				if localint_i1 <= symbolTable.HighestLocationUsed {
					fmt.Printf("Enter new term to replace %s %q\n", symbolTable.Symbols[localint_i1].TermType, symbolTable.Symbols[localint_i1].Term)

					localstring_w = lineInput("? ")
					symbolTable.Symbols[localint_i1].Term = localstring_w
					fmt.Printf("Replaced by %q\n", localstring_w)
				} else {
					fmt.Printf("Address %d too large.  Symbol table only of length %d.\n", localint_i1, symbolTable.HighestLocationUsed)
				}
			} else {
				fmt.Println(help.SyllogismHelpForSubstitute)
			}
			fmt.Println()
		} else {
			fmt.Println(symbolTable.Dump())
		}
	}

	fmt.Println("Exit from substitution routine")
}

func basicGosub5880() {
	// 5880
	//---See if conclusion possible---

	localint_c1 = intarray_c[1]
	localint_c2 = intarray_c[2]

	symbol1 := symbolTable.Symbols[localint_c1]
	symbol2 := symbolTable.Symbols[localint_c2]

	symbolTable.Iterate(1, func(i int, s *symbol.Symbol) bool {
		if s.Occurrences < 2 {
			return false
		}

		if s.DistributionCount == 0 {
			if localint_j1 == 0 {
				fmt.Println("Undistributed middle terms:")
				localint_j1 = 5
			}

			fmt.Printf("%s%s\n", basicTabString(5), s.Term)
		}

		if s.DistributionCount != 1 && s.TermType != term.TypeDesignator {
			localint_v1 = localint_i
		}
		return false
	})

	if symbolTable.NegativePremiseCount > 1 {
		localint_j1 = 6
		fmt.Println("More than one negative premise:")
	}

	if localint_j1 > 0 {
		goto Line6180
	}

	if symbolTable.NegativePremiseCount == 0 {
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

Line6180: // 6180
	fmt.Println("No possible conclusion.")
}

func basicGosub5070() {
	// 5070
	//---See if syllogism---
	var intarray_h [3]int
	var temp_symbol *symbol.Symbol

	localint_j1 = 0
	localint_v1 = 0 // flag for modern validity
	if intarray_l[0] == 0 {
		localint_j1 = 1
		return
	}

	localint_c = 0

	symbolTable.Iterate(1, func(i int, s *symbol.Symbol) bool {
		if s.Occurrences != 0 && s.Occurrences != 2 {
			if s.Occurrences != 1 {
				if localint_j1 != 2 {
					fmt.Println("Not a syllogism:")
					localint_j1 = 2
				}

				fmt.Printf("   %s %q occurs %d times in premises.\n", s.TermType, s.Term, s.Occurrences)
			}
			localint_c++
			intarray_c[localint_c] = i
		}
		return false
	})

	if localint_c != 2 {
		fmt.Println("Not a syllogism:")
		localint_j1 = 3

		if localint_c > 0 {
			fmt.Printf("   %d  terms occur exactly once in premises.\n", localint_c)

			for i := 1; i <= localint_c; i++ {
				// TODO: use tabwriter here?
				sym := symbolTable.Symbols[intarray_c[i]]
				fmt.Printf("%s%s -- %s\n", basicTabString(6), sym.Term, sym.TermType)
			}
		} else {
			fmt.Println("   no terms occur exactly once in premises.")
		}
	}

	if localint_j1 != 0 {
		return
	}

	localint_i = intarray_l[0]
	localint_l = 0

	for {
		localint_l++
		premiseSet.LinkOrder[localint_l] = localint_i
		localint_i = intarray_l[localint_i]

		if localint_i == 0 {
			break
		}
	}

	if localint_l == 1 {
		goto Line5750
	}

	if symbolTable.Symbols[intarray_c[1]].DistributionCount == 0 && symbolTable.Symbols[intarray_c[2]].DistributionCount == 1 {
		temp_symbol = symbolTable.Symbols[intarray_c[2]]
	} else {
		temp_symbol = symbolTable.Symbols[intarray_c[1]]
	}
	localint_i = 1

Line5460: // 5460
	localint_k = localint_i

Line5470: // 5470
	if premiseSet.Premises[premiseSet.LinkOrder[localint_k]].Subject == temp_symbol {
		temp_symbol = premiseSet.Premises[premiseSet.LinkOrder[localint_k]].Predicate
	} else if premiseSet.Premises[premiseSet.LinkOrder[localint_k]].Predicate == temp_symbol {
		temp_symbol = premiseSet.Premises[premiseSet.LinkOrder[localint_k]].Subject
	} else {
		localint_k++
		if localint_k <= localint_l {
			goto Line5470
		}

		temp_symbol = premiseSet.Premises[premiseSet.LinkOrder[localint_i]].Predicate

		if localint_j1 > 0 {
			fmt.Println(help.ClosedLoopHelp)
			goto Line5710
		}
		localint_j1 = 4
		fmt.Println("Not a syllogism: no way to order premises so that each premise")
		fmt.Println("shares exactly one term with its successor; there is a")
		fmt.Println(help.ClosedLoopHelp)
		goto Line5710
	}

	if localint_k != localint_i {
		localint_n = 1
		intarray_h[1] = premiseSet.LinkOrder[localint_i]

		for m := localint_i; m <= localint_k-1; m++ {
			localint_n = 3 - localint_n
			intarray_h[localint_n] = premiseSet.LinkOrder[m+1]
			premiseSet.LinkOrder[m+1] = intarray_h[3-localint_n]
		}

		premiseSet.LinkOrder[localint_i] = intarray_h[localint_n]
	}

	if localint_j1 != 0 {
		goto Line5710
	} else {
		goto Line5730
	}

Line5710: // 5710
	fmt.Println(premiseSet.Premises[premiseSet.LinkOrder[localint_i]])

Line5730: // 5730
	localint_i++

	if localint_i <= localint_l {
		goto Line5460
	}

Line5750: // 5750
	if localint_j1 > 0 {
		return
	}
	if localstring_l1 != "link" && localstring_l1 != "link*" {
		return
	}
	fmt.Println("Premises of syllogism in order of term links:")
	premiseSet.Link(localint_l, strings.HasSuffix(localstring_l1, "*"))
}

func basicGosub4890(j1 int) {
	// 4890
	//---Decrement table entries---
	var (
		pDecrement bool
		qDecrement bool
	)

	prem := premiseSet.Premises[j1]
	if prem.Form.IsNegative() {
		symbolTable.NegativePremiseCount--
		qDecrement = true
	} else if premiseSet.Premises[j1].Predicate.TermType == term.TypeDesignator {
		qDecrement = true
	}

	if prem.Form >= 2 {
		pDecrement = true
	}

	reduceDistributionCount := func(sym *symbol.Symbol, decrement bool) {
		sym.Occurrences--
		if sym.Empty() {
			sym.Term = ""
			sym.ArticleType = article.TypeNone
			sym.TermType = term.TypeUndetermined
		}

		if decrement {
			sym.DistributionCount--
		}
	}
	reduceDistributionCount(premiseSet.Premises[j1].Subject, pDecrement)
	reduceDistributionCount(premiseSet.Premises[j1].Predicate, qDecrement)
}

func basicGosub4760() {
	// 4760
	//---Delete a line---

	localint_n, _ = strconv.Atoi(stringarray_s[1])
	localint_i = 0

	for {
		localint_j1 = intarray_l[localint_i]

		if localint_j1 == 0 {
			fmt.Printf("Line %d not found\n", localint_n)
			break
		} else if localint_n == premiseSet.Premises[localint_j1].Number {
			intarray_a[0]--
			intarray_a[intarray_a[0]] = localint_j1
			intarray_l[localint_i] = intarray_l[localint_j1]
			basicGosub4890(localint_j1)
			break
		}
		localint_i = intarray_l[localint_i]
	}
}

func basicGosub6200() {
	// 6200
	//---Compute conclusion---

	compute := func() string {
		if intarray_l[0] == 0 {
			return "A is A"
		}

		symbol1 := symbolTable.Symbols[localint_c1]
		symbol2 := symbolTable.Symbols[localint_c2]
		if symbolTable.NegativePremiseCount == 0 {
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

	z := compute()

	// PRINT  conclusion
	fmt.Printf("  / %s\n", z)
	if localint_v1 != 0 {
		fmt.Print("  * Aristotle-valid only, i.e. on requirement that term ")
		fmt.Printf("%q denotes.\n", symbolTable.Symbols[localint_v1].Term)
	}
}

func basicGosub6630() {
	// 6630
	//---test offered conclusion---
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

	if d1 >= 4 {
		termType1 = term.TypeDesignator
		termType2 = localint_p1
	}
	if termType2 == term.TypeDesignator && d1 < 6 && d1 > 3 {
		d1 += 2
	}

	localstring_w = stringutil.Singularize(stringarray_w[1])
	if localint_j1 != 0 {
		stringarray_w[1] = localstring_w
	} else {
		symbolIsUndeterminedTerm := func(j int) bool {
			sym := symbolTable.Symbols[intarray_c[j]]
			if localstring_w == sym.Term {
				switch sym.TermType {
				case term.TypeUndetermined:
					fmt.Printf("Note: %q used in premises taken to be %s\n", sym.Term, termType1)
					return true
				case termType1:
					return true
				}
			}
			return false
		}

		for localint_j = 1; localint_j <= 2; localint_j++ {
			if symbolIsUndeterminedTerm(localint_j) {
				goto Line6840
			}
		}

		fmt.Printf("** Conclusion may not contain %s %q.\n", termType1, localstring_w)
		localint_j = 0
	}

Line6840: // 6840
	localstring_w = stringutil.Singularize(stringarray_w[2])
	if localint_j1 != 0 {
		if localstring_w == stringarray_w[1] {
			if d1 != 4 || termType2 == term.TypeUndetermined {
				goto Line7120
			}
			fmt.Printf("** Subject is a %s, predicate is a %s -- but\n", term.TypeDesignator, term.TypeGeneralTerm)
		} else {
			fmt.Println("** Conclusion from no premises must have same subject and predicate.")
			return
		}
	}

	if localint_j > 0 {
		localint_t1 = intarray_c[localint_j]
		localint_t2 = intarray_c[3-localint_j]
		if localstring_w != symbolTable.Symbols[localint_t2].Term {
			goto Line7060
		}
		if symbolTable.Symbols[localint_t2].TermType != term.TypeUndetermined {
			if termType2 != term.TypeUndetermined && termType2 != symbolTable.Symbols[localint_t2].TermType {
				goto Line7060
			}
		} else if termType2 != term.TypeUndetermined {
			fmt.Printf("Note: %q used in premises taken to be %s\n", symbolTable.Symbols[localint_t2].Term, termType2)
		}
		if symbolTable.NegativePremiseCount > 0 && !d1.IsNegative() {
			fmt.Println("** Negative conclusion required.")
			return
		}
		goto Line7120
	}
	if localstring_w == symbolTable.Symbols[intarray_c[1]].Term {
		localint_t2 = intarray_c[2]
	} else {
		localint_t2 = intarray_c[1]
	}
	goto Line7070

Line7060: // 7060
	fmt.Printf("** Conclusion may not contain %s %q;\n", termType2, localstring_w)

Line7070: // 7070
	fmt.Printf("** Conclusion must contain %s %q.\n", symbolTable.Symbols[localint_t2].TermType, symbolTable.Symbols[localint_t2].Term)
	return

Line7120: // 7120
	if symbolTable.NegativePremiseCount == 0 && d1.IsNegative() {
		fmt.Println("** Affirmative conclusion required.")
		return
	}

	if localint_j1 != 1 {
		if symbolTable.Symbols[localint_t1].DistributionCount == 0 && d1 > 1 && d1 < 4 {
			help.ShowTermDistributionError(symbolTable.Symbols[localint_t1].Term)
			return
		} else if symbolTable.Symbols[localint_t2].DistributionCount == 0 && (d1.IsNegative() || d1 == 6) {
			help.ShowTermDistributionError(symbolTable.Symbols[localint_t2].Term)
			return
		}
	}

	fmt.Println("-->  VALID!")

	if localint_j1 != 0 {
		if d1 > 0 {
			return
		}
		symbolTable.Symbols[0].Term = localstring_w
	} else if symbolTable.Symbols[localint_t1].DistributionCount > 0 && d1 < 2 {
		localint_v1 = localint_t1
	} else {
		if symbolTable.Symbols[localint_t2].DistributionCount > 0 && !d1.IsNegative() && d1 != form.AIsT && d1 != 6 {
			localint_v1 = localint_t2
		}

		if localint_v1 == 0 {
			return
		}
	}

	fmt.Println("    but on Aristotelian interpretation only, i.e. on requirement")
	fmt.Printf("    that term %q denotes.\n", symbolTable.Symbols[localint_v1].Term)
}

func basicGosub1840() {
	// 1840
	//---New---

	if intarray_l[0] == 0 {
		return
	}

	symbolTable = symboltable.New(basicDimMax + 2)

	for localint_j = intarray_l[0]; localint_j > 0; localint_j = intarray_l[localint_j] {
		intarray_a[0]--
		intarray_a[intarray_a[0]] = localint_j
	}
	intarray_l[0] = 0
}

func basicGosub3400(d1 form.Form, a1 int) {
	// 3400
	//---Add W$(1), W$(2) to table T$()---
	var localint_b1 int
	var termType term.Type // formerly g
	if d1.IsNegative() {
		symbolTable.NegativePremiseCount++

		if symbolTable.NegativePremiseCount > 1 && msg {
			fmt.Printf("Warning: %d negative premises\n", symbolTable.NegativePremiseCount)
		}
	}

	intarray_e[1] = article.TypeNone
	for localint_j = 1; localint_j <= 2; localint_j++ {
		localstring_w = stringarray_w[localint_j]
		if d1 < 4 {
			termType = term.TypeGeneralTerm
		} else if localint_j == 1 {
			termType = term.TypeDesignator
		} else {
			termType = localint_p1
		}

		localstring_w = stringutil.Singularize(localstring_w)
		localint_i1 = 1

		for { // 3500
			localint_i1, localint_b1 = symbolTable.Search(localint_i1, localstring_w)

			if localint_i1 > symbolTable.HighestLocationUsed {
				if localint_b1 > 0 {
					localint_i1 = localint_b1
				} else {
					symbolTable.HighestLocationUsed++
				}

				symbolTable.Symbols[localint_i1].Term = localstring_w
				symbolTable.Symbols[localint_i1].TermType = termType
				break
			}

			if termType == term.TypeUndetermined {
				if symbolTable.Symbols[localint_i1].TermType != term.TypeUndetermined || msg {
					fmt.Printf("Note: predicate term %q", localstring_w)
					fmt.Printf(" taken as the %s used earlier\n", symbolTable.Symbols[localint_i1].TermType)
				}
				break
			}
			if symbolTable.Symbols[localint_i1].TermType == term.TypeUndetermined {
				if msg {
					fmt.Printf("Note: earlier use of %q taken as the %s used here\n", localstring_w, termType)
				}
				if termType == term.TypeDesignator {
					symbolTable.Symbols[localint_i1].DistributionCount = symbolTable.Symbols[localint_i1].Occurrences
				}
				symbolTable.Symbols[localint_i1].TermType = termType
				break
			}
			if termType == symbolTable.Symbols[localint_i1].TermType {
				break
			}

			if msg {
				// TODO: remove subtraction here--we really just want to use the other term type
				fmt.Printf("Warning: %s %q has also occurred as a %s\n", termType, localstring_w, 3-termType)
			}

			localint_i1++
		}

		if intarray_e[localint_j] == article.TypeNone {

			// 3740
			if symbolTable.Symbols[localint_i1].ArticleType != article.TypeNone || localstring_w == stringarray_w[localint_j] {
				goto Line3780
			}

			if stringutil.HasPrefixVowel(localstring_w) {
				// AN
				intarray_e[localint_j] = article.TypeAn
			} else {
				// A
				intarray_e[localint_j] = article.TypeA
			}
		}

		symbolTable.Symbols[localint_i1].ArticleType = intarray_e[localint_j]

	Line3780: // 3780
		symbolTable.Symbols[localint_i1].Occurrences++

		if symbolTable.Symbols[localint_i1].Occurrences < 3 {
			goto Line3810
		}

		if !msg {
			goto Line3810
		}
		fmt.Printf("Warning: %s %q has occurred %d times\n", symbolTable.Symbols[localint_i1].TermType, localstring_w, symbolTable.Symbols[localint_i1].Occurrences)

	Line3810: // 3810
		if localint_j != 2 {
			subj := symbolTable.Symbols[localint_i1]
			premiseSet.Premises[a1].Subject = subj

			if d1 >= 2 {
				subj.DistributionCount++
			}

		} else {
			pred := symbolTable.Symbols[localint_i1]
			premiseSet.Premises[a1].Predicate = pred

			prem := premiseSet.Premises[a1]
			if prem.Subject == prem.Predicate {
				if msg {
					fmt.Printf("Warning: same term occurs twice in line %s\n", stringarray_s[1])
				}
			}

			if pred.TermType == term.TypeDesignator {
				d1 += 2
			}

			if d1 == 6 || d1.IsNegative() {
				pred.DistributionCount++
			}
		}

		if symbolTable.Symbols[localint_i1].Occurrences == 2 && symbolTable.Symbols[localint_i1].DistributionCount == 0 {
			if msg {
				fmt.Printf("Warning: undistributed middle term %q\n", symbolTable.Symbols[localint_i1].Term)
			}
		}
	}

	premiseSet.Premises[a1].Form = d1
}

// basicGosub4530 enters the provided line (string with line number + statement) into the list.
func basicGosub4530(s string) int {
	// 4530
	//---Enter line into list---
	localint_n, _ = strconv.Atoi(stringarray_s[1])
	localint_s = len(stringarray_s[1]) + 1
	localint_l = len(s)
	localstring_l = basicMid(s, localint_s+1, localint_l-localint_s)

	for localint_i = 0; ; localint_i = localint_j1 {
		localint_j1 = intarray_l[localint_i]

		if localint_j1 == 0 {
			break
		}

		if localint_n == premiseSet.Premises[localint_j1].Number {
			basicGosub4890(localint_j1)
			premiseSet.Premises[localint_j1] = &premise.Premise{
				Number:    localint_n,
				Statement: localstring_l,
			}
			return localint_j1
		}

		if localint_n < premiseSet.Premises[localint_j1].Number {
			break
		}
	}

	a1 := intarray_a[intarray_a[0]]
	premiseSet.Premises[a1] = &premise.Premise{
		Number:    localint_n,
		Statement: localstring_l,
	}
	intarray_l[localint_i] = a1
	intarray_l[a1] = localint_j1
	intarray_a[0]++

	return a1
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

					stringarray_w[1] = stringarray_s[2]
					stringarray_w[2] = stringarray_s[5]
					return form.AIsNotT, nil // a is not T
				} else {
					if intarray_t[4] != token.TypeTerm {
						return form.Undefined, errors.New(help.MissingPredicate)
					}

					stringarray_w[1] = stringarray_s[2]
					stringarray_w[2] = stringarray_s[4]
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

			stringarray_w[1] = stringarray_s[3]
			stringarray_w[2] = stringarray_s[5]

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
			stringarray_w[1] = stringarray_s[3]
			stringarray_w[2] = stringarray_s[6]
			return form.SomeAIsNotB, nil // some A is not B
		}
		if intarray_t[5] != token.TypeTerm {
			return form.Undefined, errors.New(help.MissingPredicate)
		}
		stringarray_w[1] = stringarray_s[3]
		stringarray_w[2] = stringarray_s[5]
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
	stringarray_w[1] = stringarray_s[3]
	stringarray_w[2] = stringarray_s[5]
	return form.AllAIsB, nil // all A is B
}

func tokenize() ([7]string, [8]token.Type, [3]article.Type, error) {
	// 2020
	//---L1$ into array S$()---

	// T(): 1:line num., 2:"/", 3:quantifier, 4:no/not, 5:is/are, 6:term
	//                     10 SOME  FRIED COCONUTS   ARE  NOT  TASTY
	//                      1   3        6            5    4     6

	// TODO: try to return when we need to return
	var returnErr error

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
	for _, localstring_l1 = range strings.Split(help.SampleData, "\n") {
		fmt.Println(localstring_l1)
		stringarray_s, intarray_t, intarray_e, err = tokenize()
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
		a1 := basicGosub4530(localstring_l1)
		basicGosub3400(d1, a1)
	}

	if msg {
		fmt.Println("Suggestion: try the LINK or LINK* command.")
	}
}

// BasicTab prints tabs in the manner of BASIC's TAB(N)
func basicTabString(n int) string {
	return strings.Repeat(" ", n)
}

// BasicMid emulates MID$
func basicMid(s string, a, b int) string {
	return string(s[a-1 : a+b-1])
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

func syllogize() bool {
	// TODO: factor into more local scope for conditionals
	var err error

	//---Input line---

	localstring_l1 = lineInput("> ")
	localint_l = len(localstring_l1)

	if localint_l == 0 {
		if msg {
			fmt.Println("Enter HELP for list of commands")
		}
		return true
	}

	for {
		localstring_l2 = basicRight(localstring_l1, 1)
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
		fmt.Println(symbolTable.Dump())
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
		if intarray_l[0] == 0 {
			fmt.Println(help.NoPremises)
		} else {
			basicGosub9060()
		}
		return true
	case "link":
		fallthrough
	case "link*":
		if intarray_l[0] == 0 {
			fmt.Println(help.NoPremises)
		} else {
			basicGosub5070()
		}
		return true
	case "list":
		fallthrough
	case "list*":
		if intarray_l[0] == 0 {
			fmt.Println(help.NoPremises)
		} else {
			premiseSet.List(intarray_l[:], strings.HasSuffix(localstring_l1, "*"))
		}
		return true
	}

	//--scan line L1$ into array S$()

	stringarray_s, intarray_t, intarray_e, err = tokenize()
	if err != nil {
		fmt.Println(err)
	}
	if intarray_t[1] != token.TypeLineNumber {
		goto Line1745
	}
	if intarray_t[2] != token.TypeReserved {
		goto Line1640
	}

	if intarray_l[0] == 0 {
		fmt.Println(help.NoPremises)
	} else {
		basicGosub4760() // delete line
	}
	return true

Line1640: // 1640
	func() {
		d1, err := basicGosub2890() // parse the line in S$()
		if err != nil {
			fmt.Println(err)
			if msg {
				fmt.Println("Enter SYNTAX for help with statements")
			}
		}
		if d1 != form.Undefined {
			a1 := basicGosub4530(localstring_l1) // enter line into list
			basicGosub3400(d1, a1)               // add terms to symbol table
		}
	}()
	return true

Line1745: // 1745
	if intarray_t[1] == token.TypeReserved {
		if msg {
			fmt.Println("Enter HELP for list of commands")
		}
		return true
	}

	// draw/test conclusion

	basicGosub5070() // is it a syl?
	if localint_j1 > 1 {
		return true
	}
	if localint_j1 == 0 {
		basicGosub5880() // poss. conclusion?
	}

	if localint_j1 > 1 {
		return true
	}

	if intarray_t[2] != token.TypeReserved {
		basicGosub6630()
	} else {
		basicGosub6200() // test/draw conclusion
	}

	return true
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
	for i := range intarray_a {
		intarray_a[i] = i
	}
	intarray_a[0] = 1

	if msg {
		fmt.Println("Enter HELP for list of commands")
	}

	for syllogize() {
		// Keep running till syllogize() returns false
	}
}
