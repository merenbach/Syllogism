package main

import "fmt"

func ExampleClosedLoopTest() {
	// We can simulate a closed loop with sample, removing line 20, and then setting line 10 to "all men are mortal" (duplicating line 70).

	msg = true
	runloop("sample")
	fmt.Println()
	runloop("20")
	runloop("10 all men are mortal")
	fmt.Println()
	runloop("link")

	// Output:
	// 10 all mortals are fools
	// 20 all athenians are men
	// 30 all philosophers are geniuses
	// 40 all people with good taste are philosophers
	// 50 richter is a diamond broker
	// 60 richter is the most hedonistic person in florida
	// 70 all men are mortal
	// 80 no genius is a fool
	// 90 all diamond brokers are people with good taste
	// 100 the most hedonistic person in florida is a decision-theorist
	// Suggestion: try the LINK or LINK* command.
	//
	// Warning: undistributed middle term "mortal"
	//
	// Not a syllogism: no way to order premises so that each premise
	// shares exactly one term with its successor; there is a
	// closed loop in the term chain within the premise set--
	// 10 all men are mortal
	// 70 all men are mortal
}

func ExampleSampleSlash() {
	msg = true
	runloop("sample")
	fmt.Println()
	runloop("/")

	// Output:
	// 10 all mortals are fools
	// 20 all athenians are men
	// 30 all philosophers are geniuses
	// 40 all people with good taste are philosophers
	// 50 richter is a diamond broker
	// 60 richter is the most hedonistic person in florida
	// 70 all men are mortal
	// 80 no genius is a fool
	// 90 all diamond brokers are people with good taste
	// 100 the most hedonistic person in florida is a decision-theorist
	// Suggestion: try the LINK or LINK* command.
	//
	//   / Some decision-theorist is not an athenian
}

func ExampleSampleList() {
	msg = true
	runloop("sample")
	fmt.Println()
	runloop("list")
	fmt.Println()
	runloop("list*")

	// Output:
	// 10 all mortals are fools
	// 20 all athenians are men
	// 30 all philosophers are geniuses
	// 40 all people with good taste are philosophers
	// 50 richter is a diamond broker
	// 60 richter is the most hedonistic person in florida
	// 70 all men are mortal
	// 80 no genius is a fool
	// 90 all diamond brokers are people with good taste
	// 100 the most hedonistic person in florida is a decision-theorist
	// Suggestion: try the LINK or LINK* command.
	//
	// 10  all mortals are fools
	// 20  all athenians are men
	// 30  all philosophers are geniuses
	// 40  all people with good taste are philosophers
	// 50  richter is a diamond broker
	// 60  richter is the most hedonistic person in florida
	// 70  all men are mortal
	// 80  no genius is a fool
	// 90  all diamond brokers are people with good taste
	// 100  the most hedonistic person in florida is a decision-theorist
	//
	// 10  all  mortal*  is  fool
	// 20  all  athenian*  is  man
	// 30  all  philosopher*  is  genius
	// 40  all  person with good taste*  is  philosopher
	// 50  richter+  is  diamond broker
	// 60  richter+  =   the most hedonistic person in florida+
	// 70  all  man*  is  mortal
	// 80  no  genius*  is  fool*
	// 90  all  diamond broker*  is  person with good taste
	// 100  the most hedonistic person in florida+  is  decision-theorist
}

func ExampleSampleLink() {
	msg = true
	runloop("sample")
	fmt.Println()
	runloop("link")
	fmt.Println()
	runloop("link*")

	// Output:
	// 10 all mortals are fools
	// 20 all athenians are men
	// 30 all philosophers are geniuses
	// 40 all people with good taste are philosophers
	// 50 richter is a diamond broker
	// 60 richter is the most hedonistic person in florida
	// 70 all men are mortal
	// 80 no genius is a fool
	// 90 all diamond brokers are people with good taste
	// 100 the most hedonistic person in florida is a decision-theorist
	// Suggestion: try the LINK or LINK* command.
	//
	// Premises of syllogism in order of term links:
	// 20  all athenians are men
	// 70  all men are mortal
	// 10  all mortals are fools
	// 80  no genius is a fool
	// 30  all philosophers are geniuses
	// 40  all people with good taste are philosophers
	// 90  all diamond brokers are people with good taste
	// 50  richter is a diamond broker
	// 60  richter is the most hedonistic person in florida
	// 100  the most hedonistic person in florida is a decision-theorist
	//
	// Premises of syllogism in order of term links:
	// 20  all  athenian*  is  man
	// 70  all  man*  is  mortal
	// 10  all  mortal*  is  fool
	// 80  no  genius*  is  fool*
	// 30  all  philosopher*  is  genius
	// 40  all  person with good taste*  is  philosopher
	// 90  all  diamond broker*  is  person with good taste
	// 50  richter+  is  diamond broker
	// 60  richter+  =   the most hedonistic person in florida+
	// 100  the most hedonistic person in florida+  is  decision-theorist
}

func ExampleSampleDump() {
	msg = true
	runloop("sample")
	fmt.Println()
	runloop("dump")

	// Output:
	// 10 all mortals are fools
	// 20 all athenians are men
	// 30 all philosophers are geniuses
	// 40 all people with good taste are philosophers
	// 50 richter is a diamond broker
	// 60 richter is the most hedonistic person in florida
	// 70 all men are mortal
	// 80 no genius is a fool
	// 90 all diamond brokers are people with good taste
	// 100 the most hedonistic person in florida is a decision-theorist
	// Suggestion: try the LINK or LINK* command.
	//
	// Highest symbol table loc. used: 11  Negative premises: 1
	// Adr.  art.  term                                   type  occurs  dist. count
	// 1     a     mortal                                 1     2       1
	// 2     a     fool                                   1     2       1
	// 3     an    athenian                               1     1       1
	// 4     a     man                                    1     2       1
	// 5     a     philosopher                            1     2       1
	// 6     a     genius                                 1     2       1
	// 7     a     person with good taste                 1     2       1
	// 8           richter                                2     2       2
	// 9     a     diamond broker                         1     2       1
	// 10          the most hedonistic person in florida  2     2       2
	// 11    a     decision-theorist                      1     1       0
}
