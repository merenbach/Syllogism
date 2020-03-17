package main

/* TODO:
* Ensure printing output matches original program--static checking doesn't make this any easier.
* Ensure input/output matches with various unpredictable inputs.
* Fix output alignment to either more closely match, or improve upon, original
* Refactor into proper Golang project layout
* Use Golang list type for Line collection?
* Improve dumping further
* Ensure that ssQuantifiers and ssCopulas match x$() and y$(), respectively, from original BASIC
*
* Porting notes on variables:
*
* l$(63) => line statements
* n(63)  => line numbers
* t$(65) => term array: words like man, socrates, etc., so anywhere we see t$(N) => symbols(N).Term
* b(63)  => term article type (index in a$ of proper article), so anywhere we see b(N) => symbols(N).ArticleType
* g(63)  => term type (index in g$ of term type), so anywhere we see g(N) => symbols(N).TermType
* o(63)  => term occurrence count, so anywhere we see o(N) => symbols(N).Occurrences
* d(63)  => term distribution count, so anywhere we see d(N) => symbols(N).DistributionCount
* b1     => first unused location in symbol table after a particular starting point
            (first slot with symbols(N).Occurrences == 0)
* i1     => local iterator index that is passed through different functions
            appears independent in substitution routine, but spans gosubs 3400 and 3950.
* l1     => highest symbol table location used, so symboltable.HighestLocationUsed
* n1     => negative premise count on symbol table, so symboltable.NegativePremiseCount
* s$     => parsed line tokens
* w$     => most recently entered premise, either for entry into l$ or for evaluation with /
* d1     => form of most recently entered premise, either for entry into l$ or for evaluation with /
* a1     => address of recently-entered line in the list of lines (???)
*/
import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/merenbach/syllogism/internal/help"
	"github.com/merenbach/syllogism/internal/stringutil"
	"github.com/merenbach/syllogism/internal/symbol"
	"github.com/merenbach/syllogism/internal/symboltable"
	"github.com/merenbach/syllogism/internal/term"
	"github.com/merenbach/syllogism/internal/tui"
)

const basicDimMax = 64

const (
	articleTypeNone = 0
	articleTypeA    = 1
	articleTypeAn   = 2
	articleTypeSm   = 3

	symbolUndeterminedType = 0
	symbolGeneralTerm      = 1
	symbolDesignator       = 2
)

const (
	formUndefined   = (-1)
	formSomeAIsB    = 0
	formSomeAIsNotB = 1
	formAllAIsB     = 2
	formNoAIsB      = 3
	formAIsT        = 4
	formAIsNotT     = 5
)

var (
	intarray_a [basicDimMax]int
	intarray_c [basicDimMax]int
	intarray_l [basicDimMax]int
	intarray_p [basicDimMax]int
	intarray_q [basicDimMax]int

	intarray_r [basicDimMax]int
	intarray_k [basicDimMax]int
	intarray_t [8]int
	intarray_e [3]int

	symbolTable = symboltable.NewSymbolTable(basicDimMax + 2)

	stringarray_g = []string{
		symbol.UndeterminedTypeString,
		symbol.GeneralTermString,
		symbol.DesignatorString,
	}
	stringarray_s [7]string // appears to hold parsed line tokens
	stringarray_w [3]string // appears to hold the most recently-input first and second terms for parsing or testing
	ssQuantifiers [8]string
	ssCopulas     [8]string
	stringarray_z [8]string

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
	localint_t     int
	localint_v1    int
	localstring_l  string
	localstring_l1 string
	localstring_l2 string
	localstring_s  string
	localstring_t  string
	localstring_w  string
	localstring_z  string

	programLines = make([]*tui.ProgramLine, basicDimMax)

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

		if s.DistributionCount > 0 {
			goto Line5980
		}

		if localint_j1 > 0 {
			goto Line5970
		}

		fmt.Println("Undistributed middle terms:")
		localint_j1 = 5

	Line5970: // 5970
		fmt.Printf("%s%s\n", basicTabString(5), s.Term)

	Line5980: // 5980
		if s.DistributionCount != 1 && s.TermType != symbolDesignator {
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

	if symbol1.DistributionCount > 0 || symbol2.DistributionCount > 0 {
		if symbol1.DistributionCount > 0 || symbol2.TermType < 2 {
			if symbol2.DistributionCount > 0 || symbol1.TermType < 2 {
				return
			}

			fmt.Printf("Term %q\n", symbol2.Term)
		} else {
			fmt.Printf("Term %q\n", symbol1.Term)
		}
	} else {
		fmt.Printf("Terms %q and %q, one of which is\n", symbol1.Term, symbol2.Term)
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
			fmt.Printf("   %d terms occur exactly once in premises.\n", localint_c)

			for i := 1; i <= localint_c; i++ {
				fmt.Printf("%s%s -- %s\n", basicTabString(6), symbolTable.Symbols[intarray_c[i]].Term, symbolTable.Symbols[intarray_c[i]].TermType)
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
		intarray_k[localint_l] = localint_i
		localint_i = intarray_l[localint_i]

		if localint_i == 0 {
			break
		}
	}

	if localint_l == 1 {
		goto Line5750
	}

	if symbolTable.Symbols[intarray_c[1]].DistributionCount == 0 && symbolTable.Symbols[intarray_c[2]].DistributionCount == 1 {
		localint_t = intarray_c[2]
	} else {
		localint_t = intarray_c[1]
	}
	localint_i = 1

Line5460: // 5460
	localint_k = localint_i

Line5470: // 5470
	if intarray_p[intarray_k[localint_k]] != localint_t {
		goto Line5500
	}
	localint_t = intarray_q[intarray_k[localint_k]]
	goto Line5520

Line5500: // 5500
	if intarray_q[intarray_k[localint_k]] != localint_t {
		goto Line5620
	}

	localint_t = intarray_p[intarray_k[localint_k]]

Line5520: // 5520
	if localint_k == localint_i {
		goto Line5610
	}

	localint_n = 1
	intarray_h[1] = intarray_k[localint_i]

	for m := localint_i; m <= localint_k-1; m++ {
		localint_n = 3 - localint_n
		intarray_h[localint_n] = intarray_k[m+1]
		intarray_k[m+1] = intarray_h[3-localint_n]
	}

	intarray_k[localint_i] = intarray_h[localint_n]

Line5610: // 5610
	if localint_j1 != 0 {
		goto Line5710
	} else {
		goto Line5730
	}

Line5620: // 5620
	localint_k++
	if localint_k <= localint_l {
		goto Line5470
	}

	localint_t = intarray_q[intarray_k[localint_i]]

	if localint_j1 > 0 {
		goto Line5700
	}
	localint_j1 = 4
	fmt.Println("Not a syllogism: no way to order premises so that each premise")
	fmt.Println("shares exactly one term with its successor; there is a")

Line5700: // 5700
	fmt.Println("closed loop in the term chain within the premise set--")

Line5710: // 5710
	fmt.Println(programLines[intarray_k[localint_i]])

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

	for localint_i = 1; localint_i <= localint_l; localint_i++ {

		if localstring_l1 == "link" {
			fmt.Println(programLines[intarray_k[localint_i]])
		} else {
			fmt.Printf("%d  ", programLines[intarray_k[localint_i]].Number)
			if intarray_r[intarray_k[localint_i]] < 6 && symbolTable.Symbols[intarray_q[intarray_k[localint_i]]].TermType == symbolDesignator {
				intarray_r[intarray_k[localint_i]] += 2
			}
			if intarray_r[intarray_k[localint_i]] < 4 {
				fmt.Printf("%s  ", ssQuantifiers[intarray_r[intarray_k[localint_i]]])
			}
			fmt.Printf("%s%s  %s%s\n", symbolTable.Symbols[intarray_p[intarray_k[localint_i]]].Term, ssCopulas[intarray_r[intarray_k[localint_i]]], symbolTable.Symbols[intarray_q[intarray_k[localint_i]]].Term, stringarray_z[intarray_r[intarray_k[localint_i]]])
		}
	}
}

func basicGosub4890() {
	// 4890
	//---Decrement table entries---
	var intarray_j [5]int

	intarray_j[1] = intarray_p[localint_j1]
	intarray_j[2] = intarray_q[localint_j1]
	if intarray_r[localint_j1]%2 != 0 {
		symbolTable.NegativePremiseCount--
		intarray_j[4] = 1
	} else if symbolTable.Symbols[intarray_q[localint_j1]].TermType == symbolDesignator {
		intarray_j[4] = 1
	} else {
		intarray_j[4] = 0
	}

	if intarray_r[localint_j1] >= 2 {
		intarray_j[3] = 1
	} else {
		intarray_j[3] = 0
	}

	reduceDistributionCount := func(k int) {
		idx := intarray_j[k]

		localsymbol := symbolTable.Symbols[idx]
		localsymbol.Occurrences--
		if localsymbol.Empty() {
			localsymbol.Term = ""
			localsymbol.ArticleType = articleTypeNone
			localsymbol.TermType = 0
		}

		localsymbol.DistributionCount -= intarray_j[k+2]
	}
	reduceDistributionCount(1)
	reduceDistributionCount(2)
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
		} else if localint_n == programLines[localint_j1].Number {
			intarray_a[0]--
			intarray_a[intarray_a[0]] = localint_j1
			intarray_l[localint_i] = intarray_l[localint_j1]
			basicGosub4890()
			break
		}
		localint_i = intarray_l[localint_i]
	}
}

func basicGosub6200() {
	// 6200
	//---Compute conclusion---

	if intarray_l[0] == 0 {
		localstring_z = "A is A"
		goto Line6580
	}

	if symbolTable.NegativePremiseCount == 0 {
		// affirmative conclusion
		if symbolTable.Symbols[localint_c1].DistributionCount == 0 {
			if symbolTable.Symbols[localint_c2].DistributionCount == 0 {
				if symbolTable.Symbols[localint_c1].ArticleType != articleTypeNone || symbolTable.Symbols[localint_c2].ArticleType == articleTypeNone {
					localstring_z = fmt.Sprintf("Some %s is %s%s", symbolTable.Symbols[localint_c1].Term, symbolTable.Symbols[localint_c2].ArticleTypeString(), symbolTable.Symbols[localint_c2].Term)
					goto Line6570
				}

				localstring_z = fmt.Sprintf("Some %s is %s%s", symbolTable.Symbols[localint_c2].Term, symbolTable.Symbols[localint_c1].ArticleTypeString(), symbolTable.Symbols[localint_c1].Term)
				goto Line6570
			}

			if symbolTable.Symbols[localint_c2].TermType == symbolDesignator {
				localstring_z = fmt.Sprintf("%s is %s%s", symbolTable.Symbols[localint_c2].Term, symbolTable.Symbols[localint_c1].ArticleTypeString(), symbolTable.Symbols[localint_c1].Term)
				goto Line6570
			}

			localstring_z = fmt.Sprintf("All %s is %s", symbolTable.Symbols[localint_c2].Term, symbolTable.Symbols[localint_c1].Term)
			goto Line6570

		}
		if symbolTable.Symbols[localint_c1].TermType == symbolDesignator {
			localstring_z = fmt.Sprintf("%s is %s%s", symbolTable.Symbols[localint_c1].Term, symbolTable.Symbols[localint_c2].ArticleTypeString(), symbolTable.Symbols[localint_c2].Term)
			goto Line6570
		}

		localstring_z = fmt.Sprintf("All %s is %s", symbolTable.Symbols[localint_c1].Term, symbolTable.Symbols[localint_c2].Term)
		goto Line6570

	} else {
		// negative conclusion
		if symbolTable.Symbols[localint_c2].DistributionCount > 0 {
			if symbolTable.Symbols[localint_c1].DistributionCount > 0 {
				if symbolTable.Symbols[localint_c1].TermType < 2 {
					if symbolTable.Symbols[localint_c2].TermType < 2 {
						if symbolTable.Symbols[localint_c1].ArticleType != articleTypeNone || symbolTable.Symbols[localint_c2].ArticleType == articleTypeNone {
							localstring_z = fmt.Sprintf("No %s is %s%s", symbolTable.Symbols[localint_c1].Term, symbolTable.Symbols[localint_c2].ArticleTypeString(), symbolTable.Symbols[localint_c2].Term)
						} else {

							localstring_z = fmt.Sprintf("No %s is %s%s", symbolTable.Symbols[localint_c2].Term, symbolTable.Symbols[localint_c1].ArticleTypeString(), symbolTable.Symbols[localint_c1].Term)
						}
					} else {

						localstring_z = fmt.Sprintf("%s is not %s%s", symbolTable.Symbols[localint_c2].Term, symbolTable.Symbols[localint_c1].ArticleTypeString(), symbolTable.Symbols[localint_c1].Term)
					}
				} else {

					localstring_z = fmt.Sprintf("%s is not %s%s", symbolTable.Symbols[localint_c1].Term, symbolTable.Symbols[localint_c2].ArticleTypeString(), symbolTable.Symbols[localint_c2].Term)
				}
			} else {
				localstring_z = fmt.Sprintf("Some %s is not %s%s", symbolTable.Symbols[localint_c1].Term, symbolTable.Symbols[localint_c2].ArticleTypeString(), symbolTable.Symbols[localint_c2].Term)
			}
		} else {
			localstring_z = fmt.Sprintf("Some %s is not %s%s", symbolTable.Symbols[localint_c2].Term, symbolTable.Symbols[localint_c1].ArticleTypeString(), symbolTable.Symbols[localint_c1].Term)
		}
	}

	goto Line6570

Line6570: // 6570
	// PRINT  conclusion

Line6580: // 6580
	fmt.Printf("  / %s\n", localstring_z)
	if localint_v1 != 0 {
		fmt.Print("  * Aristotle-valid only, i.e. on requirement that term ")
		fmt.Printf("%q denotes.\n", symbolTable.Symbols[localint_v1].Term)
	}
}

func basicGosub6630() {
	// 6630
	//---test offered conclusion---
	var localint_g1 term.Type = 1
	var localint_g2 term.Type = 1

	//--conc. poss, line in s$()
	d1, err := basicGosub2890()
	if err != nil {
		fmt.Println(err)
		if msg {
			fmt.Println("Enter SYNTAX for help with statements")
		}
	}

	if d1 < 0 {
		return
	} else if d1 >= 4 {
		localint_g1 = 2
		localint_g2 = localint_p1
	}
	if localint_g2 == 2 && d1 < 6 && d1 > 3 {
		d1 += 2
	}

	localstring_w = stringutil.Singularize(stringarray_w[1])
	if localint_j1 != 0 {
		stringarray_w[1] = localstring_w
	} else {
		for localint_j = 1; localint_j <= 2; localint_j++ {
			if localstring_w == symbolTable.Symbols[intarray_c[localint_j]].Term {
				if symbolTable.Symbols[intarray_c[localint_j]].TermType > 0 {
					if localint_g1 == symbolTable.Symbols[intarray_c[localint_j]].TermType {
						goto Line6840
					}
				} else {
					fmt.Printf("Note: %q used in premises taken to be %s\n", symbolTable.Symbols[intarray_c[localint_j]].Term, stringarray_g[localint_g1])
					goto Line6840
				}
			}
		}

		fmt.Printf("** Conclusion may not contain %s %q.\n", stringarray_g[localint_g1], localstring_w)
		localint_j = 0
	}

Line6840: // 6840
	localstring_w = stringutil.Singularize(stringarray_w[2])
	if localint_j1 != 0 {
		if localstring_w == stringarray_w[1] {
			if d1 != 4 || localint_g2 == 0 {
				goto Line7120
			}
			fmt.Printf("** Subject is a %s, predicate is a %s -- but\n", stringarray_g[2], stringarray_g[1])
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
		if symbolTable.Symbols[localint_t2].TermType > 0 {
			if localint_g2 != 0 && localint_g2 != symbolTable.Symbols[localint_t2].TermType {
				goto Line7060
			}
		} else if localint_g2 != 0 {
			fmt.Printf("Note: %q used in premises taken to be %s\n", symbolTable.Symbols[localint_t2].Term, stringarray_g[localint_g2])
		}
		if symbolTable.NegativePremiseCount != 0 && d1%2 == 0 {
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
	fmt.Printf("** Conclusion may not contain %s %q;\n", stringarray_g[localint_g2], localstring_w)

Line7070: // 7070
	fmt.Printf("** Conclusion must contain %s %q.\n", symbolTable.Symbols[localint_t2].TermType, symbolTable.Symbols[localint_t2].Term)
	return

Line7120: // 7120
	if symbolTable.NegativePremiseCount > 0 || d1%2 == 0 {
		if localint_j1 != 1 {
			if symbolTable.Symbols[localint_t1].DistributionCount > 0 || d1 <= 1 || d1 >= 4 {
				if symbolTable.Symbols[localint_t2].DistributionCount > 0 {
					goto Line7250
				}
				if d1%2 == 1 || d1 == 6 {
					fmt.Printf("** Term %q not distributed in premises\n", symbolTable.Symbols[localint_t2].Term)
					fmt.Println("   may not be distributed in conclusion.")
					return
				}
			} else {
				fmt.Printf("** Term %q not distributed in premises\n", symbolTable.Symbols[localint_t1].Term)
				fmt.Println("   may not be distributed in conclusion.")
				return
			}
		}
		goto Line7250
	}

	fmt.Println("** Affirmative conclusion required.")
	return

Line7250: // 7250
	fmt.Println("-->  VALID!")

	if localint_j1 == 0 {
		if symbolTable.Symbols[localint_t1].DistributionCount == 0 || d1 >= 2 {
			if symbolTable.Symbols[localint_t2].DistributionCount > 0 && d1%2 == 0 && d1 != 4 && d1 != 6 {
				localint_v1 = localint_t2
			}

			if localint_v1 == 0 {
				return
			}
		} else {
			localint_v1 = localint_t1
		}
	} else {
		if d1 > 0 {
			return
		}
		symbolTable.Symbols[0].Term = localstring_w
	}

	fmt.Println("    but on Aristotelian interpretation only, i.e. on requirement")
	fmt.Printf("    that term %q denotes.\n", symbolTable.Symbols[localint_v1].Term)
}

// TODO: make this a method on a line collection type
func basicGosub7460(analyze bool) {
	// 7460
	//---list---
	if !analyze {
		for localint_i = intarray_l[0]; localint_i != 0; localint_i = intarray_l[localint_i] {
			fmt.Println(programLines[localint_i])
		}
	} else {
		for localint_i = intarray_l[0]; localint_i != 0; localint_i = intarray_l[localint_i] {
			line := programLines[localint_i]
			if !line.Empty() {
				fmt.Printf("%d ", line.Number)

				if intarray_r[localint_i] < 6 && symbolTable.Symbols[intarray_q[localint_i]].TermType == symbolDesignator {
					intarray_r[localint_i] += 2
				}

				plocalinti := intarray_p[localint_i]
				qlocalinti := intarray_q[localint_i]
				rlocalinti := intarray_r[localint_i]

				if rlocalinti < 4 {
					fmt.Printf("%s  ", ssQuantifiers[rlocalinti])
				}

				fmt.Printf("%s%s  %s%s\n", symbolTable.Symbols[plocalinti].Term, ssCopulas[rlocalinti], symbolTable.Symbols[qlocalinti].Term, stringarray_z[rlocalinti])
			}
		}
	}
}

func basicGosub1840() {
	// 1840
	//---New---

	if intarray_l[0] == 0 {
		return
	}

	symbolTable = symboltable.NewSymbolTable(basicDimMax + 2)

	for localint_j = intarray_l[0]; localint_j > 0; localint_j = intarray_l[localint_j] {
		intarray_a[0]--
		intarray_a[intarray_a[0]] = localint_j
	}
	intarray_l[0] = 0
}

func basicGosub3400(d1 int, a1 int) {
	// 3400
	//---Add W$(1), W$(2) to table T$()---
	var localint_b1 int
	var localint_g term.Type
	if d1%2 == 1 {
		symbolTable.NegativePremiseCount++

		if symbolTable.NegativePremiseCount > 1 && msg {
			fmt.Printf("Warning: %d negative premises\n", symbolTable.NegativePremiseCount)
		}
	}

	intarray_e[1] = articleTypeNone
	for localint_j = 1; localint_j <= 2; localint_j++ {
		localstring_w = stringarray_w[localint_j]
		if d1 < 4 {
			localint_g = 1
		} else if localint_j == 1 {
			localint_g = 2
		} else {
			localint_g = localint_p1
		}

		localstring_w = stringutil.Singularize(localstring_w)
		localint_i1 = 1

	Line3500: // 3500
		localint_i1, localint_b1 = symbolTable.Search(localint_i1, localstring_w)

		if localint_i1 > symbolTable.HighestLocationUsed {
			if localint_b1 > 0 {
				localint_i1 = localint_b1
			} else {
				symbolTable.HighestLocationUsed++
			}

			symbolTable.Symbols[localint_i1].Term = localstring_w
			goto Line3720
		}

		if localint_g == 0 {
			if symbolTable.Symbols[localint_i1].TermType != symbolUndeterminedType || msg {
				fmt.Printf("Note: predicate term %q", localstring_w)
				fmt.Printf(" taken as the %s used earlier\n", symbolTable.Symbols[localint_i1].TermType)
			}
			goto Line3730
		}
		if symbolTable.Symbols[localint_i1].TermType == symbolUndeterminedType {
			if msg {
				fmt.Printf("Note: earlier use of %q taken as the %s used here\n", localstring_w, stringarray_g[localint_g])
			}
			goto Line3710
		}
		if localint_g == symbolTable.Symbols[localint_i1].TermType {
			goto Line3730
		}

		if msg {
			fmt.Printf("Warning: %s %q has also occurred as a %s\n", stringarray_g[localint_g], localstring_w, stringarray_g[3-localint_g])
		}

		localint_i1++
		goto Line3500

	Line3710: // 3710
		if localint_g == 2 {
			symbolTable.Symbols[localint_i1].DistributionCount = symbolTable.Symbols[localint_i1].Occurrences
		}

	Line3720: // 3720
		symbolTable.Symbols[localint_i1].TermType = localint_g

	Line3730: // 3730
		if intarray_e[localint_j] == articleTypeNone {

			// 3740
			if symbolTable.Symbols[localint_i1].ArticleType != articleTypeNone || localstring_w == stringarray_w[localint_j] {
				goto Line3780
			}

			if stringutil.HasPrefixVowel(localstring_w) {
				// AN
				intarray_e[localint_j] = articleTypeAn
			} else {
				// A
				intarray_e[localint_j] = articleTypeA
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

			intarray_p[a1] = localint_i1

			if d1 >= 2 {
				symbolTable.Symbols[localint_i1].DistributionCount++
			}

		} else {

			intarray_q[a1] = localint_i1

			if intarray_p[a1] == intarray_q[a1] {
				if msg {
					fmt.Printf("Warning: same term occurs twice in line %s\n", stringarray_s[1])
				}
			}

			if symbolTable.Symbols[localint_i1].TermType == symbolDesignator {
				d1 += 2
			}

			if d1 == 6 || d1%2 != 0 {
				symbolTable.Symbols[localint_i1].DistributionCount++
			}
		}

		if symbolTable.Symbols[localint_i1].Occurrences == 2 && symbolTable.Symbols[localint_i1].DistributionCount == 0 {
			if msg {
				fmt.Printf("Warning: undistributed middle term %q\n", symbolTable.Symbols[localint_i1].Term)
			}
		}
	}

	intarray_r[a1] = d1
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

		if localint_n == programLines[localint_j1].Number {
			basicGosub4890()
			programLines[localint_j1] = &tui.ProgramLine{
				Number:    localint_n,
				Statement: localstring_l,
			}
			return localint_j1
		}

		if localint_n < programLines[localint_j1].Number {
			break
		}
	}

	a1 := intarray_a[intarray_a[0]]
	programLines[a1] = &tui.ProgramLine{
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

func basicGosub2890() (int, error) {
	// 2890
	//---Parse line in S$()---
	if stringarray_s[2] != "all" {
		if stringarray_s[2] != "some" {
			if stringarray_s[2] != "no" {
				if intarray_t[2] != tokenizeType6Term {
					goto Line3350
				}

				if intarray_t[3] != tokenizeType5IsAre {
					goto Line3330
				}

				if stringarray_s[4] == "not" {
					if intarray_t[5] != tokenizeType6Term {
						goto Line3370
					}

					stringarray_w[1] = stringarray_s[2]
					stringarray_w[2] = stringarray_s[5]
					return formAIsNotT, nil // a is not T
				} else {
					if intarray_t[4] != tokenizeType6Term {
						goto Line3370
					}

					stringarray_w[1] = stringarray_s[2]
					stringarray_w[2] = stringarray_s[4]
					return formAIsT, nil // a is T
				}
			}

			if intarray_t[3] != tokenizeType6Term {
				goto Line3350
			}

			if intarray_t[4] != tokenizeType5IsAre {
				goto Line3330
			}

			if intarray_t[5] != tokenizeType6Term {
				goto Line3370
			}

			stringarray_w[1] = stringarray_s[3]
			stringarray_w[2] = stringarray_s[5]

			return formNoAIsB, nil // no A is B
		}
		if intarray_t[3] != tokenizeType6Term {
			goto Line3350
		}
		if intarray_t[4] != tokenizeType5IsAre {
			goto Line3330
		}
		if stringarray_s[5] == "not" {
			if intarray_t[6] != tokenizeType6Term {
				goto Line3370
			}
			stringarray_w[1] = stringarray_s[3]
			stringarray_w[2] = stringarray_s[6]
			return formSomeAIsNotB, nil // some A is not B
		}
		if intarray_t[5] != tokenizeType6Term {
			goto Line3370
		}
		stringarray_w[1] = stringarray_s[3]
		stringarray_w[2] = stringarray_s[5]
		return formSomeAIsB, nil // Some A is B
	}
	if intarray_t[3] != tokenizeType6Term {
		goto Line3350
	}
	if intarray_t[4] != tokenizeType5IsAre {
		goto Line3330
	}
	if intarray_t[5] != tokenizeType6Term {
		goto Line3370
	}
	stringarray_w[1] = stringarray_s[3]
	stringarray_w[2] = stringarray_s[5]
	return formAllAIsB, nil // all A is B

Line3330: // 3330
	return formUndefined, errors.New("** Missing copula is/are")

Line3350: // 3350
	return formUndefined, errors.New("** Subject term bad or missing")

Line3370: // 3370
	return formUndefined, errors.New("** Predicate term bad or missing")
}

const (
	tokenizeType0Reserved   = 0
	tokenizeType1LineNum    = 1
	tokenizeType2Slash      = 2
	tokenizeType3Quantifier = 3
	tokenizeType4NoNot      = 4
	tokenizeType5IsAre      = 5
	tokenizeType6Term       = 6
)

func tokenize() ([7]string, [8]int, [3]int, error) {
	// 2020
	//---L1$ into array S$()---

	// T(): 1:line num., 2:"/", 3:quantifier, 4:no/not, 5:is/are, 6:term
	//                     10 SOME  FRIED COCONUTS   ARE  NOT  TASTY
	//                      1   3        6            5    4     6

	// TODO: try to return when we need to return
	var returnErr error

	var shadowstringarray_s [7]string
	var shadowintarray_t [8]int
	var shadowintarray_e [3]int

	localint_p1 = 0
	shadowintarray_e[2] = articleTypeNone
	localint_j = 1

	localint_l = len(localstring_l1)

	nextToken := func() {
		localint_j++
	}
	setToken := func(s string, t int) {
		shadowstringarray_s[localint_j] = s
		shadowintarray_t[localint_j] = t
	}
	addTermToken := func(s string) {
		setToken(s, tokenizeType6Term)
	}
	closeToken := func(s string, t int) {
		setToken(s, t)
		nextToken()
	}

	line2670 := func(tabCount int) { // 2670
		returnErr = fmt.Errorf("%s^\nReserved word %q may not occur within a term", basicTabString(tabCount), localstring_s)
		shadowintarray_t[1] = tokenizeType0Reserved
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
			closeToken(localstring_s, tokenizeType2Slash)
		} else {
			if _, err := strconv.Atoi(localstring_s); err != nil {
				returnErr = fmt.Errorf("%s^   Invalid numeral or command", basicTabString(localint_i+len(localstring_s)))
				break
			}
			closeToken(localstring_s, tokenizeType1LineNum)
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

		case "all", "some":
			if shadowintarray_t[localint_j] == tokenizeType6Term {
				line2670(localint_i + localint_k - 1)
				break Iterate
			}
			closeToken(localstring_s, tokenizeType3Quantifier)

		case "no", "not":
			if shadowintarray_t[localint_j] == tokenizeType6Term {
				line2670(localint_i + localint_k - 1)
				break Iterate
			}
			closeToken(localstring_s, tokenizeType4NoNot)

		case "is", "are":
			if shadowintarray_t[localint_j] != tokenizeType6Term {
				line2670(localint_i + localint_k - 1)
				break Iterate
			} else if shadowintarray_t[localint_j-1] == tokenizeType5IsAre || shadowintarray_t[localint_j-2] == tokenizeType5IsAre {
				line2670(localint_i + localint_k - 1)
				break Iterate
			}
			// NOTE: This is needed here, and not above, because of positioning of to-be verbs in lines.
			// All/some and no/not will occur either at the beginning of the line or right after a to-be verb.
			// To-be verbs are the only words to (legitimately) show up right after the end of a term token (type 6).
			nextToken()
			closeToken(localstring_s, tokenizeType5IsAre)

		default:
			if shadowintarray_t[localint_j] == tokenizeType6Term {
				shadowstringarray_s[localint_j] += " " + localstring_s

			} else if shadowintarray_t[localint_j-1] != tokenizeType5IsAre && shadowintarray_t[localint_j-2] != tokenizeType5IsAre {
				addTermToken(localstring_s)

			} else if localstring_s != "a" && localstring_s != "an" && localstring_s != "sm" {
				if localstring_s == "the" {
					// DESIGNATOR (definite article)
					localint_p1 = 2
				}
				addTermToken(localstring_s)

			} else if localint_i == localint_l {
				addTermToken(localstring_s)

			} else {
				switch localstring_s {
				case "a":
					shadowintarray_e[2] = articleTypeA
				case "an":
					shadowintarray_e[2] = articleTypeAn
				case "sm":
					shadowintarray_e[2] = articleTypeSm
				}
				// GENERAL TERM (indefinite article)
				localint_p1 = 1
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

func syllogize() {
	// TODO: factor into more local scope for conditionals
	var err error
	fmt.Println("Running...")
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

	ssQuantifiers[0] = "some"
	ssQuantifiers[1] = "some"
	ssQuantifiers[2] = "all"
	ssQuantifiers[3] = "no"
	ssQuantifiers[4] = ""
	ssQuantifiers[5] = ""
	ssQuantifiers[6] = ""
	ssQuantifiers[7] = ""

	ssCopulas[0] = "  is"
	ssCopulas[1] = "  is not"
	ssCopulas[2] = "*  is"
	ssCopulas[3] = "*  is"
	ssCopulas[4] = "+  is"
	ssCopulas[5] = "+  is not"
	ssCopulas[6] = "+  = "
	ssCopulas[7] = "+   = / = "

	stringarray_z[0] = ""
	stringarray_z[1] = "*"
	stringarray_z[2] = ""
	stringarray_z[3] = "*"
	stringarray_z[4] = ""
	stringarray_z[5] = "*"
	stringarray_z[6] = "+"
	stringarray_z[7] = "*"

	/*

	   // 602
	    rem /error check/ for err = 0 to 7 : print x$(err),y$(err),z$(err) : next err
	*/

	msg = true
	for i := range intarray_a {
		intarray_a[i] = i
	}
	intarray_a[0] = 1

Line1070: // 1070
	if msg {
		fmt.Println("Enter HELP for list of commands")
	}

Line1080: // 1080
	//---Input line---

	localstring_l1 = lineInput("> ")
	localint_l = len(localstring_l1)

	if localint_l == 0 {
		goto Line1070
	}

Line1120: // 1120
	localstring_l2 = basicRight(localstring_l1, 1)
	if localstring_l2 == " " {
		goto Line1160
	}

	if localstring_l2 != "." && localstring_l2 != "?" && localstring_l2 != "!" {
		goto Line1181
	}

	fmt.Printf("%s ^   Punctuation mark ignored\n", basicTabString(localint_l))

Line1160: // 1160
	if localint_l == 1 {
		goto Line1080
	}

	localint_l--
	localstring_l1 = basicLeft(localstring_l1, localint_l)
	goto Line1120

Line1181: // 1181
	if basicLeft(localstring_l1, 1) != " " {
		goto Line1190
	}
	if localint_l == 1 {
		goto Line1080
	}
	localint_l--
	localstring_l1 = basicRight(localstring_l1, localint_l)
	goto Line1181

Line1190: // 1190
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
		os.Exit(0)
		goto Line1080
	case "new":
		fmt.Println("Begin new syllogism")
		basicGosub1840()
		goto Line1080
	case "sample":
		basicGosub1840()
		basicGosub8980()
		goto Line1080
	case "help":
		help.ShowInputsHelp()
		goto Line1080
	case "syntax":
		help.ShowSyntaxHelp()
		goto Line1080
	case "info":
		help.ShowGeneralHelp()
		goto Line1080
	case "dump":
		fmt.Println(symbolTable.Dump())
		goto Line1080
	case "msg":
		msg = !msg

		fmt.Print("Messages turned ")
		if msg {
			fmt.Println("on")
		} else {
			fmt.Println("off")
		}
		goto Line1080

	case "substitute":
		if intarray_l[0] == 0 {
			goto Line1612
		}
		basicGosub9060()
		goto Line1080
	case "link":
		fallthrough
	case "link*":
		if intarray_l[0] == 0 {
			goto Line1612
		}
		basicGosub5070()
		goto Line1080
	case "list":
		fallthrough
	case "list*":
		if intarray_l[0] == 0 {
			goto Line1612
		}

		basicGosub7460(strings.HasSuffix(localstring_l1, "*"))
		goto Line1080
	}

	//--scan line L1$ into array S$()

	stringarray_s, intarray_t, intarray_e, err = tokenize()
	if err != nil {
		fmt.Println(err)
	}
	if intarray_t[1] != tokenizeType1LineNum {
		goto Line1745
	}
	if intarray_t[2] != tokenizeType0Reserved {
		goto Line1640
	}

	if intarray_l[0] != 0 {
		goto Line1620
	}

Line1612: // 1612
	fmt.Println("No premises")
	goto Line1080

Line1620: // 1620
	basicGosub4760() // delete line
	goto Line1080

Line1640: // 1640
	func() {
		d1, err := basicGosub2890() // parse the line in S$()
		if err != nil {
			fmt.Println(err)
			if msg {
				fmt.Println("Enter SYNTAX for help with statements")
			}
		}
		if d1 != formUndefined {
			a1 := basicGosub4530(localstring_l1) // enter line into list
			basicGosub3400(d1, a1)               // add terms to symbol table
		}
	}()
	goto Line1080

Line1745: // 1745
	if intarray_t[1] == tokenizeType0Reserved {
		goto Line1070
	}

	// draw/test conclusion

	basicGosub5070() // is it a syl?
	if localint_j1 > 1 {
		goto Line1080
	}
	if localint_j1 == 0 {
		basicGosub5880() // poss. conclusion?
	}

	if localint_j1 > 1 {
		goto Line1080
	}

	if intarray_t[2] != tokenizeType0Reserved {
		basicGosub6630()
	} else {
		basicGosub6200() // test/draw conclusion
	}

	goto Line1080

}

func main() {
	syllogize()
}
