rem Syllogism 1.0. November 8, 2002
rem I edited this program in 2002, for compatibility with freeware BASIC
rem interpreters for the Mac: Chipmunk BASIC 3.5.7 and Metal BASIC 1.7.3. 
rem I hope compatibility with two indpendently developed BASICs should
rem assure some universality regardless of platform. The following
rem notes should help anybody with a similar project.
rem Summary of changes.....
rem * Metal doesn't support MID$() as a command. Both support LCASE$().
rem * Metal crashes when it reads an empty ("") DATA item. Had to hack it.
rem * Metal requires ; between PRINT items. Also, it doesn't add a space
rem when the ; separates a number and a string.
rem Chipmunk buggy with IF...ELSE on one line--use colons (and GOTOs)
rem Chipmunk requires quotes around DATA items
rem Chipmunk buggy when using integer variables (J%), use floating point (J)
rem Neither support PRINT CHR$(27) as a way to clear the screen; used cls
rem Peace. Ben Sharvy. luvnpeas99@yahoo.com
cls : print "Syllogism Program Copyright (c) 1988 Richard Sharvy"
print "Syllogism 1.0 (c) 2002 Richard Sharvy's estate"
print "Ben Sharvy: luvnpeas99@yahoo.com or bsharvy@efn.org" : print
dim a(63),c(63),term_dist_count(63),term_type(63),l(63),line_numbers(63),term_occurrences(63),p(63),q(63)
dim r(63),term_article(63),k(63),j(4),t(7),recent_article_types(2),h(2)
dim article_strings$(3),line_strings$(63),term_strings$(65)
dim g$(2),s$(6),w$(2),x$(7),y$(7),z$(7)
read article_strings$(1),article_strings$(2),article_strings$(3),g$(0),g$(1),g$(2)
data "a ","an ","sm ","undetermined type","general term","designator"
for i = 0 to 7
	read x$(i),y$(i),z$(i)
	rem Metal crashes when "" is in DATA statement...
	if x$(i) = "<null>" then x$(i) = ""
	if y$(i) = "<null>" then y$(i) = ""
	if z$(i) = "<null>" then z$(i) = ""
next i
data "some","  is","<null>","some","  is not","*"
data "all","*  is","<null>","no","*  is","*"
data "<null>","+  is","<null>","<null>","+  is not","*"
data "<null>","+  = ","+","<null>","+   = / = ","*"
rem /error check/ for err = 0 to 7 : print x$(err),y$(err),z$(err) : next err
dim u$(75),v$(75)
i = 0
do
	i = i + 1
	read u$(i),v$(i)
loop until u$(i) = "ZZZZZ"
u1 = i - 1
data "socrates","socrates","parmenides","parmenides","epimenides","epimenides"
data "mice","mouse","lice","louse","geese","goose"
data "children","child","oxen","ox","people","person","teeth","tooth"
data "wolves","wolf","wives","wife","selves","self","lives","life","leaves","leaf"
data "shelves","shelf","elves","elf","dwarves","dwarf","knives","knife","thieves","thief"
data "neckties","necktie","hippies","hippie","yippies","yippie","yuppies","yuppie"
data "moonies","moonie","druggies","druggie","cookies","cookie","commies","commie"
data "groupies","groupie","tomatoes","tomato"
data "alcibiades","alcibiades","thales","thales","aries","aries","athens","athens"
data "species","species","feces","feces","geniuses","genius","sorites","sorites"
data "crises","crisis","emphases","emphasis","memoranda","memorandum","theses","thesis"
data "automata","automaton","formulae","formula","stigmata","stigma","lemmata","lemma"
data "vertices","vertex","vortices","vortex","indices","index","codices","codex"
data "matrices","matrix"
data "gasses","gas","gases","gas","buses","bus","aches","ache","headaches","headache"
data "grits","grits","molasses","molasses","gas","gas","christmas","christmas"
data "mathematics","mathematics","semantics","semantics","physics","physics"
data "metaphysics","metaphysics","ethics","ethics","linguistics","linguistics"
data "kiwis","kiwi","israelis","israeli"
data "goyim","goy","seraphim","seraph","cherubim","cherub"
data "semen","semen","amen","amen"
data "ZZZZZ","ZZZZZ"
msg = -1 : a(0) = 1 : for i = 1 to 63 : a(i) = i : next i
rem---Input line--- : rem 1080
	gosub 1070
	do
		print ">";
		line input l1$
		l = len(l1$)
		if l = 0 then
			gosub 1070
		else
			do
				l1$ = rtrim$(l1$)
				l2$ = right$(l1$,1)
				if l2$ <> "." and l2$ <> "?" and l2$ <> "!" then exit do
				l = len(l1$)
				print space$(l);" ^   Punctuation mark ignored"
				l1$ = left$(l1$, l - 1)
			loop
			l1$ = ltrim$(l1$)
			if len(l1$) > 0 then
				rem / FOR I = 1 TO L
					rem / V = ASC(MID$(L1$,I,1))
					rem / IF V >= 65 AND V <= 90  THEN  MID$(L1$,I,1) = CHR$(V+32)
					rem / NEXT
				rem Metal doesn't support mid$ as command, but lcase$() is well supported...
				l1$ = lcase$(l1$)
				if l1$ = "stop" then
					if msg then print "(Some versions support typing CONT to continue)"
					exit do
				elseif l1$ = "new" then
					print "Begin new syllogism"
					procNEW
				elseif l1$ = "sample" then
					procNEW
					gosub 8980
				elseif l1$ = "help" then
					procLIST_VALID_INPUTS
				elseif l1$ = "syntax" then
					procSYNTAX
				elseif l1$ = "info" then
					procINFO
				elseif l1$ = "dump" then
					gosub 8890
				elseif l1$ = "msg" then
					msg = not (msg)
					print "Messages turned ";
					if msg then print "on" : else print "off"
				elseif l1$ = "substitute" then
					if l(0) = 0 then
						gosub 1612
					else
						gosub 9060
					endif
				elseif l1$ = "link" or l1$ = "link*" then
					if l(0) = 0 then
						gosub 1612
					else
						gosub 5070
					endif
				elseif l1$ = "list" or l1$ = "list*" then
					if l(0) = 0 then
						gosub 1612
					else
						procLIST
					endif
				else
					gosub 1570
				endif
			endif
		endif
	loop
	stop
1570 rem--scan line L1$ into array S$()
	gosub 2020
	if t(1) = 1 then
		if t(2) then
			procPARSE_LINE : rem parse the line in S$()
			if d1 >= 0 then
				gosub 4530 : rem enter line into list
				procADD_TABLE_STRINGS : rem add terms to symbol table
			endif
		else
			if l(0) then
				procDELETE_LINE : rem delete line
			else
				gosub 1612
			endif
		endif
	else
		if t(1) = 0 then
			gosub 1070
		else
			rem draw/test conclusion
			gosub 5070 : rem is it a syl?
			if not (j1 > 1) then
				if j1 = 0 then gosub 5880 : rem poss. conclusion?
				if not (j1 > 1) then
					if t(2) then gosub 6630 : else gosub 6200 : rem test/draw conclusion
				endif
			endif
		endif
	endif
	return
1070 rem [am] precedes input
	if msg then print "Enter HELP for list of commands"
	return
1612 rem [am] subroutine for no premises
	print "No premises"
	return
rem---New--- : rem [am] 1840
def procNEW
	if l(0) <> 0 then
		for i = 1 to length_of_symbol_table
			term_dist_count(i) = 0 : term_strings$(i) = "" : term_article(i) = 0 : term_occurrences(i) = 0 : term_type(i) = 0
		next i
		length_of_symbol_table = 0 : negative_premise_count = 0
		j = l(0)
		do
			a(0) = a(0)-1
			a(a(0)) = j
			j = l(j)
		loop until not (j > 0)
		l(0) = 0
	endif
	endproc
2020 rem---L1$ into array S$()---
	rem T(): 1:line num., 2:"/", 3:quantifier, 4:no/not, 5:is/are, 6:term
	rem                     10 SOME  FRIED COCONUTS   ARE  NOT  TASTY
	rem                      1   3        6            5    4     6
	for j = 1 to 6
		s$(j) = ""
		t(j) = 0
	next j
	p1 = 0
	recent_article_types(2) = 0
	j = 1
	i = 1
	l = len(l1$)
	do
		do
			if i > l then 2885
			s$ = mid$(l1$,i,1)
			if s$ <> " " then exit do
			i = i+1
		loop
		for k = 1 to (l - i)
			s$ = mid$(l1$,i+k,1)
			if s$ = " " then exit for
		next k
		s$ = mid$(l1$,i,k) : rem S$ is set to next word
		if j <= 1 then
			if s$ = "/" then
				t(1) = 2
			else
				n = len(s$)
				if n > 4 then
					procERROR_INVALID_CMD(i + n)
					goto 2885
				else
					for n = 1 to len(s$)
						t$ = mid$(s$,n,1)
						if asc(t$) > 57 or asc(t$) < 48 then
							procERROR_INVALID_CMD(i + n)
							goto 2885
						endif
					next n
					t(1) = 1
				endif
			endif
			goto 2840
		endif
rem Scan : rem [am] 2520
		if s$ = "somebody" or s$ = "something" or s$ = "nobody" or s$ = "nothing" then
			procERR_RESERVED_WORD(s$, i + k - 1)
			t(1) = 0
			goto 2885
		elseif s$ = "someone" or s$ = "everyone" or s$ = "everybody" or s$ = "everything" then
			procERR_RESERVED_WORD(s$, i + k - 1)
			t(1) = 0
			goto 2885
		elseif s$ = "all" or s$ = "some" then
			if t(j) = 6 then
				procERR_RESERVED_WORD(s$, i + k - 1)
				t(1) = 0
				goto 2885
			else
				t(j) = 3
				goto 2840
			endif
		elseif s$ = "no" or s$ = "not" then
			if t(j) = 6 then
				procERR_RESERVED_WORD(s$, i + k - 1)
				t(1) = 0
				goto 2885
			else
				t(j) = 4
				goto 2840
			endif
		elseif s$ = "is" or s$ = "are" then
			if t(j) = 6 then
				if not (t(j-1) = 5 or t(j-2) = 5) then
					j = j+1
					t(j) = 5
					goto 2840
				endif
			endif
			procERR_RESERVED_WORD(s$, i + k - 1)
			t(1) = 0
			goto 2885
		elseif t(j) <> 6 then
			if t(j-1) = 5 or t(j-2) = 5 then
				if s$ = "a" or s$ = "an" or s$ = "sm" then
					if i <> l then
						if s$ = "a" then
							recent_article_types(2) = 1
						elseif s$ = "an" then
							recent_article_types(2) = 2
						else
							recent_article_types(2) = 3
						endif
						p1 = 1
					else
						gosub 2790
					endif
				else
					if s$ = "the" then p1 = 2
					gosub 2790
				endif
			else
				gosub 2790
			endif
		else
			s$(j) = s$(j)+" "+s$
		endif
		goto 2860
2840	s$(j) = s$
		j = j+1
2860	i = k+i
	loop until j > 6
2885 return
rem [am] 2460 : rem subroutine from 2020
def procERROR_INVALID_CMD(spaces)
	print space$(spaces);"^   Invalid numeral or command"
	endproc
rem [am] 2670 : rem subroutine from 2020
def procERR_RESERVED_WORD(string$, spaces)
	print space$(spaces);"^"
	print "Reserved word '";string$;"' may not occur within a term"
	endproc
2790 rem [am] subroutine from 2020
	s$(j) = s$
	t(j) = 6
	return
rem---Parse line in S$()--- : rem [am] 2890
def procPARSE_LINE
	d1 = -1
	if s$(2) = "all" then
		if t(3) <> 6 then
			procERR_MISSING_SUBJECT_TERM
		elseif t(4) <> 5 then
			procERR_MISSING_COPULA
		elseif t(5) <> 6 then
			procERR_MISSING_PREDICATE
		else
			w$(1) = s$(3)
			w$(2) = s$(5)
			d1 = 2 : rem all A is B
		endif
	elseif s$(2) = "some" then
		if t(3) <> 6 then
			procERR_MISSING_SUBJECT_TERM
		elseif t(4) <> 5 then
			procERR_MISSING_COPULA
		elseif s$(5) <> "not" then
			if t(5) <> 6 then
				procERR_MISSING_PREDICATE
			else
				w$(1) = s$(3)
				w$(2) = s$(5)
				d1 = 0 : rem Some A is B
			endif
		else
			if t(6) <> 6 then
				procERR_MISSING_PREDICATE
			else
				w$(1) = s$(3)
				w$(2) = s$(6)
				d1 = 1 : rem some A is not B
			endif
		endif
	elseif s$(2) = "no" then
		if t(3) <> 6 then
			procERR_MISSING_SUBJECT_TERM
		elseif t(4) <> 5 then
			procERR_MISSING_COPULA
		elseif t(5) <> 6 then
			procERR_MISSING_PREDICATE
			procERR_HELP
		else
			w$(1) = s$(3)
			w$(2) = s$(5)
			d1 = 3 : rem no A is B
		endif
	elseif t(2) <> 6 then
		procERR_MISSING_SUBJECT_TERM
	elseif t(3) = 5 then
		w$(1) = s$(2)
		if s$(4) <> "not" then
			if t(4) <> 6 then
			procERR_MISSING_PREDICATE
		endif
			d1 = 4 : rem a is T
			w$(2) = s$(4)
		else
			if t(5) <> 6 then
				procERR_MISSING_PREDICATE
			else
				d1 = 5 : rem a is not T
				w$(2) = s$(5)
			endif
		endif
	else
		procERR_MISSING_COPULA
	endif
	endproc
rem [am] 3330 : rem subroutine from 2890
def procERR_MISSING_COPULA
	print "** Missing copula is/are"
	procERR_HELP
	endproc
rem [am] 3350 : rem subroutine from 2890
def procERR_MISSING_SUBJECT_TERM
	print "** Subject term bad or missing"
	procERR_HELP
	endproc
rem [am] 3370 : rem subroutine from 2890
def procERR_MISSING_PREDICATE
	print "** Predicate term bad or missing"
	procERR_HELP
	endproc
rem [am] 3380 : rem subroutine from 2890
def procERR_HELP
	if msg then print "Enter SYNTAX for help with statements"
	endproc
rem---Add W$(1), W$(2) to table term_strings$()--- : rem [am] 3400
def procADD_TABLE_STRINGS
	if (d1 mod 2) <> 0 then
		negative_premise_count = negative_premise_count+1
		if negative_premise_count > 1 and msg then print "Warning: ";negative_premise_count;" negative premises"
	endif
	recent_article_types(1) = 0
	for j = 1 to 2
		w$ = w$(j)
		if d1 < 4 then g = 1 :  else if j = 1 then g = 2 :  else g = p1
		w$ = fnCONVERT_WSTR_TO_SINGULAR$(w$)
		i1 = 1
		do
	 		procSEARCH_TSTR_FOR_WSTR
			if i1 > length_of_symbol_table then
				if b1 > 0 then i1 = b1 : else length_of_symbol_table = length_of_symbol_table + 1
				term_strings$(i1) = w$
				term_type(i1) = g
				exit do
			endif
			if g = 0 then
				if not(term_type(i1) = 0 and not msg) then
					print "Note: predicate term '";w$;"'";
					print " taken as the ";g$(term_type(i1));" used earlier"
				endif
				exit do
			endif
			if term_type(i1) = 0 then
				if msg then print "Note: earlier use of '";w$;"' taken as the ";g$(g);" used here"
				if g = 2 then term_dist_count(i1) = term_occurrences(i1)
				exit do
			endif
			if g = term_type(i1) then exit do
			if msg then print "Warning: ";g$(g);" '";w$;"' has also occurred as a ";g$(3-g)
			i1 = i1+1
		loop
		if recent_article_types(j) > 0 or not (term_article(i1) > 0 or w$ = w$(j)) then
			if not (recent_article_types(j) > 0) and not (term_article(i1) > 0 or w$ = w$(j)) then
				a$ = left$(w$,1)
				if a$ = "a" or a$ = "e" or a$ = "i" or a$ = "o" or a$ = "u" then recent_article_types(j) = 2 : else recent_article_types(j) = 1
			endif
			term_article(i1) = recent_article_types(j)
		endif
		term_occurrences(i1) = term_occurrences(i1)+1
		if term_occurrences(i1) >= 3 then
			if msg then
				print "Warning: ";g$(term_type(i1));" '";w$;"' has occurred";term_occurrences(i1);"times"
			endif
		endif
		if j <> 2 then
			p(a1) = i1
			if d1 >= 2 then term_dist_count(i1) = term_dist_count(i1)+1
		else
			q(a1) = i1
			if p(a1) = q(a1) then
				if msg then print "Warning: same term occurs twice in line ";s$(1)
			endif
			if term_type(i1) = 2 then d1 = d1+2
			if d1 = 6 or d1 mod 2 then term_dist_count(i1) = term_dist_count(i1)+1
		endif
		if not (term_occurrences(i1) <> 2 or term_dist_count(i1) > 0) then
			if msg then print "Warning: undistributed middle term '";term_strings$(i1);"'"
		endif
	next j
	r(a1) = d1
	endproc
rem---Search term_strings$() for W$ from I1 to length_of_symbol_table--- : rem [am] 3950
def procSEARCH_TSTR_FOR_WSTR
	rem If found, I1 = length_of_symbol_table; else I1 = length_of_symbol_table+1. B1 set to 1st empty loc.
	b1 = 0
	do
		if i1 > length_of_symbol_table then exit do
		if term_strings$(i1) = w$ then exit do
		if term_occurrences(i1) = 0 and b1 = 0 then b1 = i1
		i1 = i1+1
	loop
	endproc
rem---Convert W$ to singular--- : rem [am] 4040
def fnCONVERT_WSTR_TO_SINGULAR$(word$)
	local x$, y$
	local my_matched_plural
	if not fnHAS_PREFIX(word$, "the ") then
		x$ = ""
		i = 1
		rem [am] n = 1
		do
			if fnIS_BLANK(word$, i) = true then
				word$ = x$
				exit do
			else
				m = fnNEXT_SPACE(word$, i)
				y$ = mid$(word$, i, m)
				my_matched_plural = false
				for k = 1 to u1
					if y$ = u$(k) then
						y$ = v$(k)
						x$ = fnAPPEND$(x$, y$)
						my_matched_plural = true
					endif
				next k
				if my_matched_plural = false then
					if fnHAS_SUFFIX(y$, "men") then
						y$ = left$(y$,len(y$)-2)+"an"
					elseif fnHAS_SUFFIX(y$, "s") then
						if not (fnHAS_SUFFIX(y$, "ss") or fnHAS_SUFFIX(y$, "us") or fnHAS_SUFFIX(y$, "is") or fnHAS_SUFFIX(y$, "'s")) then
							y$ = left$(y$,len(y$)-1)
							if fnHAS_SUFFIX(y$, "xe") then
								y$ = left$(y$,len(y$)-1)
							elseif fnHAS_SUFFIX(y$, "ie") then
								y$ = left$(y$,len(y$)-2)
								y$ = y$+"y"
							elseif fnHAS_SUFFIX(y$, "sse") or fnHAS_SUFFIX(y$, "she") or fnHAS_SUFFIX(y$, "che") then
								y$ = left$(y$,len(y$)-1)
							endif
						endif
					endif
					x$ = fnAPPEND$(x$, y$)
				endif
				rem [am] n = n+1
				i = m+i
			endif
		loop
	endif
	=word$
4530 rem---Enter line into list---
	n = val(s$(1))
	s = len(s$(1))+1
	l = len(l1$)
	l$ = mid$(l1$,s+1,l-s)
	i = 0
	do
		j1 = l(i)
		if j1 = 0 then
			gosub 4690
			exit do
		elseif n = line_numbers(j1) then
			procDECREMENT_TABLE_ENTRIES
			line_strings$(j1) = l$
			a1 = j1
			exit do
		elseif n < line_numbers(j1) then
			gosub 4690
			exit do
		else
			i = j1
		endif
	loop
	return
4690 rem subroutine from 4530
	a1 = a(a(0))
	line_strings$(a1) = l$
	line_numbers(a1) = n
	l(i) = a1
	l(a1) = j1
	a(0) = a(0)+1
	return
rem---Delete a line--- : rem [am] 4760
def procDELETE_LINE
	n = val(s$(1))
	i = 0
	do
		j1 = l(i)
		if j1 = 0 then
			print "Line ";n;" not found"
			exit do
		elseif n = line_numbers(j1) then
			a(0) = a(0) - 1
			a(a(0)) = j1
			l(i) = l(j1)
			procDECREMENT_TABLE_ENTRIES
			exit do
		else
			i = l(i)
		endif
	loop
	endproc
rem---Decrement table entries--- : rem [am] 4890
def procDECREMENT_TABLE_ENTRIES
	j(1) = p(j1)
	j(2) = q(j1)
	if r(j1) mod 2 <> 0 then
		negative_premise_count = negative_premise_count-1
		j(4) = 1
	else
		if term_type(q(j1)) = 2 then j(4) = 1 : else j(4) = 0
	endif
	if r(j1) >= 2 then j(3) = 1 : else j(3) = 0
	for k = 1 to 2
		term_occurrences(j(k)) = term_occurrences(j(k))-1
		if not (term_occurrences(j(k)) > 0) then
			term_strings$(j(k)) = ""
			term_article(j(k)) = 0
			term_type(j(k)) = 0
		endif
		term_dist_count(j(k)) = term_dist_count(j(k))-j(k+2)
	next k
	endproc
5070 rem---See if syllogism---
	j1 = 0
	v1 = 0 : rem flag for modern validity
	if l(0) then
		c = 0
		for i = 1 to length_of_symbol_table
			if not(term_occurrences(i) = 0 or term_occurrences(i) = 2) then
				if term_occurrences(i) = 1 then
					c = c+1
					c(c) = i
				else
					if j1 <> 2 then
						print "Not a syllogism:"
						j1 = 2
					endif
					print "   ";g$(term_type(i));" '";term_strings$(i);"' occurs ";term_occurrences(i);" times in premises."
				endif
			endif
		next i
		if c <> 2 then
			print "Not a syllogism:"
			j1 = 3
			if not(c > 0) then
				print "   no terms occur exactly once in premises."
			else
				print "   ";c;" terms occur exactly once in premises."
				for i = 1 to c
					print space$(6);term_strings$(c(i));" -- ";g$(term_type(c(i)))
				next i
			endif
		endif
		if j1 = 0 then
			rem [am] in spirit, should be "if not (j1)""
			i = l(0)
			l = 0
			rem [am] loop until not(i) might be more appropriate, in terms
			rem [am] of semantic translation, but it doesn't work here...
			do
				l = l+1
				k(l) = i
				i = l(i)
			loop until i = 0
			if l <> 1 then
				if term_dist_count(c(1)) = 0 and term_dist_count(c(2)) = 1 then t = c(2) : else t = c(1)
				for i = 1 to l
					for k = i to l
						if p(k(k)) = t or q(k(k)) = t then
							if p(k(k)) = t then
								t = q(k(k))
							elseif q(k(k)) = t then
								t = p(k(k))
							endif
							if k <> i then
								n = 1
								h(1) = k(i)
								for m = i to k-1
									n = 3-n
									h(n) = k(m+1)
									k(m+1) = h(3-n)
									next m
								k(i) = h(n)
							endif
							if j1 then gosub 5710
							exit for
						endif
					next k
					if k > l then
						t = q(k(i))
						if not(j1 > 0) then
							j1 = 4
							print "Not a syllogism: no way to order premises so that each premise"
							print "shares exactly one term with its successor; there is a"
						endif
						print "closed loop in the term chain within the premise set--"
						gosub 5710
					endif
				next i
			endif
			if not (j1 > 0) then
				if l1$ = "link" or l1$ = "link*" then
					print "Premises of syllogism in order of term links:"
					for i = 1 to l
						print line_numbers(k(i));" ";
						if l1$ <> "link" then
							if r(k(i)) < 6 and term_type(q(k(i))) = 2 then r(k(i)) = r(k(i))+2
							if r(k(i)) < 4 then print x$(r(k(i)));"  ";
							print term_strings$(p(k(i)));y$(r(k(i)));"  ";term_strings$(q(k(i)));z$(r(k(i)))
						else
							print line_strings$(k(i))
						endif
					next i
				endif
			endif
		endif
	else
		j1 = 1
	endif
	return
5710 rem [am] subroutine from 5070 see if syll
	print line_numbers(k(i));
	print line_strings$(k(i))
	return
5880 rem---See if conclusion possible---
	c1 = c(1)
	c2 = c(2)
	for i = 1 to length_of_symbol_table
		if term_occurrences(i) >= 2 then
			if not (term_dist_count(i) > 0) then
				if not (j1 > 0) then
					print "Undistributed middle terms:"
					j1 = 5
				endif
				print space$(5);term_strings$(i)
			endif
			if term_dist_count(i) <> 1 and term_type(i) <> 2 then v1 = i
		endif
	next i
	if negative_premise_count >= 2 then
		j1 = 6
		print "More than one negative premise:"
	endif
	if j1 > 0 then
		gosub 6180
	else
		if not (negative_premise_count = 0) then
			if not (term_dist_count(c1) > 0 or term_dist_count(c2) > 0) then
				print "Terms '";term_strings$(c1);"' and '";term_strings$(c2);"',";" one of which is"
				gosub 6150
			elseif not (term_dist_count(c1) > 0 or term_type(c2) < 2) then
				print "Term '";term_strings$(c1);"'"
				gosub 6150
			elseif not (term_dist_count(c2) > 0 or term_type(c1) < 2) then
				print "Term '";term_strings$(c2);"'"
				gosub 6150
			endif
		endif
	endif
	return
6150 rem [am] subroutine from 5880
	print "required in predicate of negative conclusion"
	print "not distributed in the premises."
	j1 = 7
	gosub 6180
	return
6180 rem [am] no possible conc. (from 5880)
	print "No possible conclusion."
	return
6200 rem---Compute conclusion---
	if l(0) = 0 then
		z$ = "A is A"
	else
		if not (negative_premise_count = 0) then
			rem negative conclusion
			if not (term_dist_count(c2) > 0) then
				z$ = "Some "+term_strings$(c2)+" is not "+article_strings$(term_article(c1))+term_strings$(c1)
			elseif not (term_dist_count(c1) > 0) then
				z$ = "Some "+term_strings$(c1)+" is not "+article_strings$(term_article(c2))+term_strings$(c2)
			elseif term_type(c1) >= 2 then
				z$ = term_strings$(c1)+" is not "+article_strings$(term_article(c2))+term_strings$(c2)
			elseif term_type(c2) >= 2 then
				z$ = term_strings$(c2)+" is not "+article_strings$(term_article(c1))+term_strings$(c1)
			elseif not (term_article(c1) > 0 or term_article(c2) = 0) then
				z$ = "No "+term_strings$(c2)+" is "+article_strings$(term_article(c1))+term_strings$(c1)
			else
				z$ = "No "+term_strings$(c1)+" is "+article_strings$(term_article(c2))+term_strings$(c2)
			endif
		else
			rem affirmative conclusion
			if not (term_dist_count(c1) = 0) then
				if term_type(c1) <> 2 then
					z$ = "All "+term_strings$(c1)+" is "+term_strings$(c2)
				else
					z$ = term_strings$(c1)+" is "+article_strings$(term_article(c2))+term_strings$(c2)
				endif
			elseif not (term_dist_count(c2) = 0) then
				if term_type(c2) <> 2 then
					z$ = "All "+term_strings$(c2)+" is "+term_strings$(c1)
				else
					z$ = term_strings$(c2)+" is "+article_strings$(term_article(c1))+term_strings$(c1)
				endif
			else
				if not (term_article(c1) > 0 or term_article(c2) = 0) then
					z$ = "Some "+term_strings$(c2)+" is "+article_strings$(term_article(c1))+term_strings$(c1)
				else
					z$ = "Some "+term_strings$(c1)+" is "+article_strings$(term_article(c2))+term_strings$(c2)
				endif
			endif
		endif
	endif
	rem PRINT  conclusion
	print "  / ";z$
	if not(v1 = 0) then
		print "  * Aristotle-valid only, i.e. on requirement that term ";
		print "'";term_strings$(v1);"' denotes."
	endif
	return
6630 rem---test offered conclusion---
	rem--conc. poss, line in s$()
	procPARSE_LINE
	if not (d1 < 0) then
		if d1 < 4 then g1 = 1 : g2 = 1 : else g1 = 2 : g2 = p1
		if g2 = 2 and d1 < 6 and d1 > 3 then d1 = d1+2
		w$ = w$(1)
		w$ = fnCONVERT_WSTR_TO_SINGULAR$(w$)
		if not (j1 = 0) then
			w$(1) = w$
		else
			for j = 1 to 2
				if w$ = term_strings$(c(j)) then
					if not (term_type(c(j)) > 0) then
						print "Note: '";term_strings$(c(j));"' used in premises taken to be ";g$(g1)
						exit for
					elseif g1 = term_type(c(j)) then
						exit for
					endif
				endif
			next j
			if j > 2 then
				print "** Conclusion may not contain ";g$(g1);" '";w$;"'."
				j = 0
			endif
		endif
		w$ = w$(2)
		w$ = fnCONVERT_WSTR_TO_SINGULAR$(w$)
		if not (j1 = 0) then
			if w$ = w$(1) then
				if d1 <> 4 or g2 = 0 then
					goto 7120
				else
					print "** Subject is a ";g$(2);", predicate is a ";g$(1);" -- but"
				endif
			endif
			gosub 6880
		else
			if not (j > 0) then
				if w$ = term_strings$(c(1)) then t2 = c(2) : else t2 = c(1)
			else
				t1 = c(j)
				t2 = c(3-j)
				if w$ = term_strings$(t2) then
					if not (term_type(t2) > 0) or (g2 = 0 or g2 = term_type(t2)) then
						if not (term_type(t2) > 0) and not (g2 = 0) then
							print "Note: '";term_strings$(t2);"' used in premises taken to be ";g$(g2)
						endif
						if not (negative_premise_count = 0 or (d1 mod 2) = 1) then
							print "** Negative conclusion required."
							goto 7370
						else
							goto 7120
						endif
					endif
				endif
				print "** Conclusion may not contain ";g$(g2);" '";w$;"';"
			endif
			print "** Conclusion must contain ";g$(term_type(t2));" '";term_strings$(t2);"'."
		endif
		goto 7370
7120	if (negative_premise_count > 0 or d1 mod 2 = 0) then
			print "** Affirmative conclusion required."
			goto 7370
		endif
		if j1 <> 1 then
			if not (term_dist_count(t1) > 0 or d1 <= 1 or d1 >= 4) then
				print "** Term '";term_strings$(t1);"' not distributed in premises"
				gosub 7180
				goto 7370
			endif
			if not (term_dist_count(t2) > 0) then
				if not (d1 mod 2 = 0 and d1 <> 6) then
					print "** Term '";term_strings$(t2);"' not distributed in premises"
					gosub 7180
					goto 7370
				endif
			endif
		endif
		print "-->  VALID!"
		if not (j1 = 0) then
			if d1 > 0 then goto 7370
			term_strings$(0) = w$
		elseif not (term_dist_count(t1) = 0 or d1 >= 2) then
			v1 = t1
		else
			if term_dist_count(t2) > 0 and d1 mod 2 = 0 and d1 <> 4 and d1 <> 6 then v1 = t2
			if v1 = 0 then goto 7370
		endif
		print "    but on Aristotelian interpretation only, i.e. on requirement"
		print "    that term '";term_strings$(v1);"' denotes."
	endif
7370 return
6880 rem [am] subroutine from 6630
	print "** Conclusion from no premises must have same subject and predic";
	print "ate."
	return
7180 rem [am] subroutine from 6630
	print "   may not be distributed in conclusion."
	return
rem---list--- : rem [am] 7460
def procLIST
	i = l(0)
	while not (i = 0)
		print line_numbers(i);" ";
		if l1$ <> "list" then
			if r(i) < 6 and term_type(q(i)) = 2 then r(i) = r(i)+2
			if r(i) < 4 then print x$(r(i));"  ";
			print term_strings$(p(i));y$(r(i));"  ";term_strings$(q(i));z$(r(i))
		else
			print line_strings$(i)
		endif
		i = l(i)
	wend
	endproc
rem---List valid inputs--- : rem [am] 7660
def procLIST_VALID_INPUTS
	cls : print "Valid commands are:"
	print "   <n>  [ <statement> ]   Insert, delete, or replace premise number  <n> "
	print space$(28);"Examples:   10  All men are mortal"
	print space$(40);"10"
	print "  DUMP";space$(15);"Prints symbol table, distribution count, etc."
	print "  HELP";space$(15);"Prints this list"
	print "  INFO";space$(15);"Gives information about syllogisms"
	print "  LIST";space$(15);"Lists premises"
	print "  LIST*";space$(14);"Same, but displays distribution analysis:"
	print space$(25);"distributed positions marked with '*', "
	print space$(25);"designators marked with '+'"
	print "  LINK";space$(15);"Lists premises in order of term-links (if possible)"
	print "  LINK*";space$(14);"Same, but in distribution-analysis format"
	print "  MSG";space$(16);"Turns on/off Printing of certain messages and warnings"
	print "  NEW";space$(16);"Erases current syllogism"
	print "  SAMPLE";space$(13);"Erases current syllogism and enters sample syllogism"
	print "  STOP";space$(15);"Stops entire program"
	print "  SUBSTITUTE";space$(9);"Allows uniform substitution of new terms in ";
	print "old premises"
	print "  SYNTAX";space$(13);"Explains statement syntax, with examples"
	print "  /";space$(18);"Asks program to draw conclusion"
	print "  /  <statement>";space$(5);"Tests  <statement>  as conclusion"
	print space$(25);"Note: this can be done even if there are no premises"
	endproc
rem--"syntax"-- : rem [am] 7960
def procSYNTAX
	cls : print "Valid statement forms:"
	print "  All    <general term #1>   is/are       <general term #2>"
	print "  Some   <general term #1>   is/are       <general term #2>"
	print "  Some   <general term #1>   is/are not   <general term #2>"
	print "  No     <general term #1>   is/are       <general term #2>"
	print
	print "   <designator>      is/are       <general term>"
	print "   <designator>      is/are not   <general term>"
	print "   <designator A>    is/are       <designator B>"
	print "   <designator A>    is/are not   <designator B>" : print
	print "Examples:"
	print "  All tall men are Greek gods             The teacher of Plato is wise"
	print "  Some cheese is tasty                    Socrates is not handsome"
	print "  Some cheese is not soft                 The teacher of Plato is Socrates"
	print "  No libertarians are cringing wimps      Socrates is not the";
	print " teacher of Thales"
	print
	print "Since e.g. 'Socrates is grunch' is ambiguous ('grunch' could be"
	print "either a designator or a general term), the program will try to"
	print "resolve the ambiguity from other uses of the term in the syllogism."
	print "The indefinite article 'sm' may be used with mass terms in predicates"
	print "(e.g. 'This puddle is sm ink') to ensure that the mass term is taken"
	print "as a general term rather than as a designator."
	endproc
rem---Info--- : rem [am] 8290
def procINFO
	cls : print "   To use this program, enter a syllogism, one line at a time,"
	print "and  THEN  test conclusions or ask the program to draw a conclusion."
	print
	print "   A syllogism as (mis)defined here is a (possibly empty) set of"
	print "numbered premises, each of a form specified in the SYNTAX list."
	print "No term may occur more than twice.  Exactly two terms must occur"
	print "exactly once: these are the two 'end' terms, which will appear in"
	print "the conclusion.  Furthermore, each premise must have exactly one"
	print "term in common with its successor, for some ordering of the premises."
	print "Example:"
	print "   10 Socrates is a Greek"
	print "   20 All men are mortal"
	print "   30 All Greeks are men"
	print "   40 No gods are mortal" : print
	print "Note: using a '/' command to draw or test a conclusion does not"
	print "require you to stop.  You can continue, adding or deleting premises"
	print "and drawing and testing more conclusions." : print
	print "Reference:  H. Gensler, 'A Simplified Decision Procedure for Categor-"
	print "   ical Syllogisms,' Notre Dame J. of Formal Logic 14 (1973) 457-466."
	endproc
8890 rem---"Dump" values of variables---
	print "Highest symbol table loc. used:";length_of_symbol_table;"  Negative premises:";negative_premise_count
	if length_of_symbol_table <> 0 then
		print "Adr. art. term";space$(48-14);"type       occurs    dist. count"
		for i = 1 to length_of_symbol_table
			rem Metal's lack of tabbing gets difficult here...
			itab = 7-len(str$(i))
			astringtab = 11-len(article_strings$(term_article(i)))-7
			tstringtab = 49-len(term_strings$(i))-11
			gtab = 60-len(str$(term_type(i)))-49
			otab = 71-len(str$(term_occurrences(i)))-60
			print i;space$(itab);article_strings$(term_article(i));space$(astringtab);term_strings$(i);space$(tstringtab);term_type(i);space$(gtab);
			print term_occurrences(i);space$(otab);term_dist_count(i)
		next i
	endif
	return
8980 rem--sample--
	for z8 = 1 to 10 : read l1$ : print l1$
	gosub 2020 : procPARSE_LINE : gosub 4530 : procADD_TABLE_STRINGS
	next z8
9004 data "10 all mortals are fools"
	data "20 all athenians are men"
	data "30 all philosophers are geniuses"
	data "40 all people with good taste are philosophers"
	data "50 richter is a diamond broker"
	data "60 richter is the most hedonistic person in florida"
	data "70 all men are mortal"
	data "80 no genius is a fool"
	data "90 all diamond brokers are people with good taste"
	data "100 the most hedonistic person in florida is a decision-theorist"
	restore 9004
	if msg then print "Suggestion: try the LINK or LINK* command."
	return
9060 rem---Substitute terms---
	do
		print "Enter address of old term; or 0 for help, -1 to exit, -2 for dump"
		input i1
		if i1 <> -1 then
			if i1 = -2 then
				gosub 8890
			else
				if i1 = 0 then
					print "   This subroutine allows a term in a syllogism to be uniformly"
					print "replaced by another term.  This is useful e.g. for finding an"
					print "interpretation which actually makes the premises true, to produce as"
					print "an obvious example of invalidity an argument having exactly the same"
					print "logical form.  The substitution does not take place in the premises"
					print "as originally entered; it takes place in the terms as stored within"
					print "the program.  Thus, the LINK and LIST commands will display the"
					print "original premises; to see the changed ones, use the LIST* and LINK*"
					print "commands."
					print "   To find the 'addresses' of the terms, enter -2 to run the DUMP."
					print "   Warning: if you replace a term with another one already occurring"
					print "in the syllogism, the result will not make much sense.  However,"
					print "this routine does not convert entered term to lower-case or singular."
				else
					if i1 > length_of_symbol_table then
						print "Address ";i1;" too large.  Symbol table only of length ";length_of_symbol_table
					else
						print "Enter new term to replace ";g$(term_type(i1));" '";term_strings$(i1);"'"
						input w$
						term_strings$(i1) = w$
						print "Replaced by '";w$;"'"
					endif
				endif
				print
			endif
		else
			exit do
		endif
	loop
	print "Exit from substitution routine"
	return
end

rem from 4040
def fnIS_BLANK(string$, index)
	local isblank
	local i
	
	isblank = true

	for i = index to len(string$)
		if mid$(w$, i, 1) <> " " then
			isblank = false
			exit for
		endif
	next i
	=isblank

rem from 4040
def fnNEXT_SPACE(string$, index)
	local m
	for m = 1 to (len(string$) - index)
		if mid$(string$, index + m, 1) = " " then exit for
	next m
	=m

rem from 4040 : rem [am] 4470
def fnAPPEND$(string$, suffix$)
	if len(string$) = 0 then
		string$ = suffix$
	else
		string$ = string$ + " " + suffix$
	endif
	=string$

def fnHAS_PREFIX(string$, fragment$)
	local l, flag
	flag = false
	len_string = len(string$)
	len_fragment = len(fragment$)
	if len_string >= len_fragment and len_fragment > 0 then
		if left$(string$, len_fragment) = fragment$ then flag = true
	endif
	=flag
end

def fnHAS_SUFFIX(string$, fragment$)
	local l, flag
	flag = false
	len_string = len(string$)
	len_fragment = len(fragment$)
	if len_string >= len_fragment and len_fragment > 0 then
		if right$(string$, len_fragment) = fragment$ then flag = true
	endif
	=flag
end
