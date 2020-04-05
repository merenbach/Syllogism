package help

import (
	"fmt"

	"github.com/merenbach/syllogism/internal/tui"
)

const (
	// SyllogismCopyright holds a copyright message.
	syllogismCopyright = `Syllogism Program Copyright (c) 1988 Richard Sharvy
Syllogism 1.0 (c) 2002 Richard Sharvy's estate
Ben Sharvy: luvnpeas99@yahoo.com or bsharvy@efn.org

Golang port by Andrew Merenbach <andrew@merenbach.com>.`

	// SampleData contains a sample program to run.
	SampleData = `10 all mortals are fools
20 all athenians are men
30 all philosophers are geniuses
40 all people with good taste are philosophers
50 richter is a diamond broker
60 richter is the most hedonistic person in florida
70 all men are mortal
80 no genius is a fool
90 all diamond brokers are people with good taste
100 the most hedonistic person in florida is a decision-theorist`

	// SyllogismHelpForInputs holds the help message for inputs.
	syllogismHelpForInputs = `Valid commands are:
   <n>  [ <statement> ]   Insert, delete, or replace premise number  <n>
                            Examples:   10  All men are mortal
                                        10
  DUMP               Prints symbol table, distribution count, etc.
  HELP               Prints this list
  INFO               Gives information about syllogisms
  LIST               Lists premises
  LIST*              Same, but displays distribution analysis:
                         distributed positions marked with '*',
                         designators marked with '+'
  LINK               Lists premises in order of term-links (if possible)
  LINK*              Same, but in distribution-analysis format
  MSG                Turns on/off Printing of certain messages and warnings
  NEW                Erases current syllogism
  SAMPLE             Erases current syllogism and enters sample syllogism
  STOP               Stops entire program
  SUBSTITUTE         Allows uniform substitution of new terms in old premises
  SYNTAX             Explains statement syntax, with examples
  /                  Asks program to draw conclusion
  /  <statement>     Tests  <statement>  as conclusion
                         Note: this can be done even if there are no premises`

	// SyllogismHelpForInfo holds the help for the info command.
	syllogismHelpForInfo = `   To use this program, enter a syllogism, one line at a time,
and  THEN  test conclusions or ask the program to draw a conclusion.

   A syllogism as (mis)defined here is a (possibly empty) set of
numbered premises, each of a form specified in the SYNTAX list.
No term may occur more than twice.  Exactly two terms must occur
exactly once: these are the two 'end' terms, which will appear in
the conclusion.  Furthermore, each premise must have exactly one
term in common with its successor, for some ordering of the premises.
Example:
   10 Socrates is a Greek
   20 All men are mortal
   30 All Greeks are men
   40 No gods are mortal

Note: using a '/' command to draw or test a conclusion does not
require you to stop.  You can continue, adding or deleting premises
and drawing and testing more conclusions.

Reference:  H. Gensler, 'A Simplified Decision Procedure for Categor-
   ical Syllogisms,' Notre Dame J. of Formal Logic 14 (1973) 457-466.`

	// SyllogismHelpForSubstitute holds help for the substitute command.
	SyllogismHelpForSubstitute = `   This subroutine allows a term in a syllogism to be uniformly
replaced by another term.  This is useful e.g. for finding an
interpretation which actually makes the premises true, to produce as
an obvious example of invalidity an argument having exactly the same
logical form.  The substitution does not take place in the premises
as originally entered; it takes place in the terms as stored within
the program.  Thus, the LINK and LIST commands will display the
original premises; to see the changed ones, use the LIST* and LINK*
commands.
   To find the 'addresses' of the terms, enter -2 to run the DUMP.
   Warning: if you replace a term with another one already occurring
in the syllogism, the result will not make much sense.  However,
this routine does not convert entered term to lower-case or singular.`

	// SyllogismHelpForSyntax holds help for the syntax command
	syllogismHelpForSyntax = `Valid statement forms:
  All    <general term #1>   is/are       <general term #2>
  Some   <general term #1>   is/are       <general term #2>
  Some   <general term #1>   is/are not   <general term #2>
  No     <general term #1>   is/are       <general term #2>

   <designator>      is/are       <general term>
   <designator>      is/are not   <general term>
   <designator A>    is/are       <designator B>
   <designator A>    is/are not   <designator B>

Examples:
  All tall men are Greek gods             The teacher of Plato is wise
  Some cheese is tasty                    Socrates is not handsome
  Some cheese is not soft                 The teacher of Plato is Socrates
  No libertarians are cringing wimps      Socrates is not the teacher of Thales

Since e.g. 'Socrates is grunch' is ambiguous ('grunch' could be
either a designator or a general term), the program will try to
resolve the ambiguity from other uses of the term in the syllogism.
The indefinite article 'sm' may be used with mass terms in predicates
(e.g. 'This puddle is sm ink') to ensure that the mass term is taken
as a general term rather than as a designator.`
)

// ShowCopyright shows the copyright for the program.
func ShowCopyright() {
	tui.Clear()
	fmt.Println(syllogismCopyright)
}

// ShowGeneralHelp shows help for the program.
func ShowGeneralHelp() {
	// 8290
	//---Info---
	tui.Clear()
	fmt.Println(syllogismHelpForInfo)
}

// ShowInputsHelp prints help for inputs.
func ShowInputsHelp() {
	// 7660
	//---List valid inputs---
	tui.Clear()
	fmt.Println(syllogismHelpForInputs)
}

// ShowSyntaxHelp prints help for syntax.
func ShowSyntaxHelp() {
	// 7960
	//--"syntax"--
	tui.Clear()
	fmt.Println(syllogismHelpForSyntax)
}

// ShowTermDistributionError prints a term distribution error.
func ShowTermDistributionError(t string) {
	fmt.Printf("** Term %q not distributed in premises\n", t)
	fmt.Println("   may not be distributed in conclusion.")
}

// ClosedLoopHelp is a help message to print when erroring about a closed loop.
const ClosedLoopHelp = "closed loop in the term chain within the premise set--"

// MissingCopula indicates a missing copula.
const MissingCopula = "** Missing copula is/are"

// MissingPredicate indicates a missing predicate term.
const MissingPredicate = "** Predicate term bad or missing"

// MissingSubject indicates a missing subject term.
const MissingSubject = "** Subject term bad or missing"

// NoPremises indicates an empty premise set.
const NoPremises = "No premises"

// NoPossibleConclusion indicates that no conclusion is possible.
const NoPossibleConclusion = "No possible conclusion."

// ConclusionFromNoPremises indicates a conclusion from no premises.
const ConclusionFromNoPremises = "** Conclusion from no premises must have same subject and predicate."
