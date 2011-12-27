rem Syllogism 1.0. November 8, 2002
rem I edited this program in 2002, for compatibility with freeware BASIC
rem interpreters for the Mac: Chipmunk BASIC 3.5.7 and Metal BASIC 1.7.3. 
rem I hope compatibility with two indpendently developed BASICs should
rem assure some universality regardless of platform. The following
rem notes should help anybody with a similar project.
rem Summary of changes.....
rem * Metal doesn't support PRINT TAB(N). It supports the command HTAB, but a
rem bug makes it useless for formatting more than one column of text. The only
rem standard BASIC solution is the " " character, implemented with tb$ and Left$()
rem * Metal doesn't support MID$() as a command. Both support LCASE$().
rem * Metal crashes when it reads an empty ("") DATA item. Had to hack it.
rem * Metal requires ; between PRINT items. Also, it doesn't add a space
rem when the ; separates a number and a string.
rem Chipmunk buggy with IF...ELSE on one line--use colons (and GOTOs)
rem Chipmunk requires quotes around DATA items
rem Chipmunk buggy when using integer variables (J%), use floating point (J)
rem Neither support PRINT CHR$(27) as a way to clear the screen; used cls
rem Peace. Ben Sharvy. luvnpeas99@yahoo.com
rem Metal doesn't support PRINT TAB(N)...
tb$ = "                                                  "
cls : print "Syllogism Program Copyright (c) 1988 Richard Sharvy"
print "Syllogism 1.0 (c) 2002 Richard Sharvy's estate"
print "Ben Sharvy: luvnpeas99@yahoo.com or bsharvy@efn.org" : print
dim a(63),c(63),d(63),g(63),l(63),n(63),o(63),p(63),q(63)
dim r(63),b(63),k(63),j(4),t(7),e(2),h(2)
dim a$(3),l$(63),t$(65)
dim g$(2),s$(6),w$(2),x$(7),y$(7),z$(7)
read a$(1),a$(2),a$(3),g$(0),g$(1),g$(2)
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
770 i = i+1
771  read u$(i),v$(i)
780  if u$(i) <> "ZZZZZ" then 770
790 u1 = i-1
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
1070 if msg then print "Enter HELP for list of commands"
1080 rem---Input line---
	print ">";
	line input l1$
	l = len(l1$)
	if l = 0 then 1070
1120 l2$ = right$(l1$,1)
		if l2$ = " " then 1160
			if l2$ <> "." and l2$ <> "?" and l2$ <> "!" then 1181
			print left$(tb$,l);" ^   Punctuation mark ignored"
1160	if l = 1 then 1080
		l = l-1 : l1$ = left$(l1$,l) : goto 1120
1181 if left$(l1$,1) <> " " then 1190
		if l = 1 then 1080
		l = l-1 : l1$ = right$(l1$,l) : goto 1181
1190 rem / FOR I = 1 TO L
		rem / V = ASC(MID$(L1$,I,1))
		rem / IF V >= 65 AND V <= 90  THEN  MID$(L1$,I,1) = CHR$(V+32)
		rem / NEXT
	rem Metal doesn't support mid$ as command, but lcase$() is well supported...
	l1$ = lcase$(l1$)
	if l1$ <> "stop" then 1230
		if msg then print "(Some versions support typing CONT to continue)"
	stop
		goto 1080
1230 if l1$ <> "new" then 1270
		print "Begin new syllogism"
		gosub 1840
		goto 1080
1270 if l1$ <> "sample" then 1310
		gosub 1840
		gosub 8980
		goto 1080
1310 if l1$ <> "help" then 1340
		gosub 7660
		goto 1080
1340 if l1$ <> "syntax" then 1370
		gosub 7960
		goto 1080
1370 if l1$ <> "info" then 1430
		gosub 8290
		goto 1080
1430 if l1$ <> "dump" then 1460
		gosub 8890
		goto 1080
1460 if l1$ <> "msg" then 1500
		msg = not (msg)
		print "Messages turned ";
		if msg then print "on" : else print "off"
		goto 1080
1500 if l1$ <> "substitute" then 1540
		if l(0) = 0 then 1612
		gosub 9060
		goto 1080
1540 if l1$ <> "link" and l1$ <> "link*" then 1561
		if l(0) = 0 then 1612
		gosub 5070
		goto 1080
1561 if l1$ <> "list" and l1$ <> "list*" then 1570
		if l(0) = 0 then 1612
		gosub 7460
		goto 1080
1570 rem--scan line L1$ into array S$()
	gosub 2020
	if t(1) <> 1 then 1745
		if t(2) then 1640
			if l(0) then 1620
1612			print "No premises" : goto 1080
1620		gosub 4760 : rem delete line
			goto 1080
1640	gosub 2890 : rem parse the line in S$()
		if d1 < 0 then 1080
		gosub 4530 : rem enter line into list
		gosub 3400 : rem add terms to symbol table
		goto 1080
1745 if t(1) = 0 then 1070
	rem draw/test conclusion
	gosub 5070 : rem is it a syl?
	if j1 > 1 then 1080
	if j1 = 0 then gosub 5880 : rem poss. conclusion?
		if j1 > 1 then 1080
		if t(2) then gosub 6630 : else gosub 6200 : rem test/draw conclusion
	goto 1080
1840 rem---New---
	if l(0) = 0 then 2010
	for i = 1 to l1
		d(i) = 0 : t$(i) = "" : b(i) = 0 : o(i) = 0 : g(i) = 0
	next i
	l1 = 0 : n1 = 0
	j = l(0)
1960 a(0) = a(0)-1
		a(a(0)) = j
		j = l(j)
		if j > 0 then 1960
	l(0) = 0
2010 return
2020 rem---L1$ into array S$()---
	rem T(): 1:line num., 2:"/", 3:quantifier, 4:no/not, 5:is/are, 6:term
	rem                     10 SOME  FRIED COCONUTS   ARE  NOT  TASTY
	rem                      1   3        6            5    4     6
	for j = 1 to 6 : s$(j) = "" : t(j) = 0 : next j
	p1 = 0 : e(2) = 0 : j = 1 : i = 1
	l = len(l1$)
2180 if i > l then 2885
		s$ = mid$(l1$,i,1)
		if s$ = " " then i = i+1 : goto 2180
	k = 1
2240 if i+k > l then 2290
		s$ = mid$(l1$,i+k,1)
		if s$ <> " " then k = k+1 : goto 2240
2290 s$ = mid$(l1$,i,k) : rem S$ is set to next word
	if j > 1 then 2520
		if s$ <> "/" then 2400
		t(1) = 2
		goto 2840
2400 n = len(s$)
	if n > 4 then 2460
		n = 1
2430	t$ = mid$(s$,n,1)
		if asc(t$) <= 57 and asc(t$) >= 48 then 2480
2460		print left$(tb$,i+n);"^   Invalid numeral or command"
			goto 2885
2480	n = n+1
		if n <= len(s$) then 2430
	t(1) = 1
	goto 2840
2520 rem Scan
	if s$ = "somebody" or s$ = "something" or s$ = "nobody" or s$ = "nothing" then 2670
	if s$ = "someone" or s$ = "everyone" or s$ = "everybody" or s$ = "everything" then 2670
	if s$ <> "all" and s$ <> "some" then 2570
		if t(j) = 6 then 2670
		t(j) = 3
		goto 2840
2570 if s$ <> "no" and s$ <> "not" then 2610
		if t(j) = 6 then 2670
		t(j) = 4
		goto 2840
2610 if s$ <> "is" and s$ <> "are" then 2710
		if t(j) <> 6 then 2670
		if t(j-1) = 5 or t(j-2) = 5 then 2670
			j = j+1
			t(j) = 5
			goto 2840
2670	print left$(tb$,i+k-1);"^"
		print "Reserved word '";s$;"' may not occur within a term"
		t(1) = 0
		goto 2885
2710 if t(j) = 6 then 2820
	if t(j-1) <> 5 and t(j-2) <> 5 then 2790
	if s$ <> "a" and s$ <> "an" and s$ <> "sm" then 2780
		if i = l then 2790
		if s$ = "a" then e(2) = 1 :  else if s$ = "an" then e(2) = 2 :  else e(2) = 3
		p1 = 1
		goto 2860
2780 if s$ = "the" then p1 = 2
2790 s$(j) = s$
	t(j) = 6
	goto 2860
2820 s$(j) = s$(j)+" "+s$
	goto 2860
2840 s$(j) = s$
	j = j+1
2860 i = k+i
	if j <= 6 then 2180
2885 return
2890 rem---Parse line in S$()---
	d1 = -1
	if s$(2) <> "all" then 2990
		if t(3) <> 6 then 3350
		if t(4) <> 5 then 3330
		if t(5) <> 6 then 3370
		w$(1) = s$(3)
		w$(2) = s$(5)
		d1 = 2 : rem all A is B
		goto 3390
2990 if s$(2) <> "some" then 3130
		if t(3) <> 6 then 3350
		if t(4) <> 5 then 3330
		if s$(5) = "not" then 3080
			if t(5) <> 6 then 3370
			w$(1) = s$(3)
			w$(2) = s$(5)
			d1 = 0 : rem Some A is B
			goto 3390
3080	if t(6) <> 6 then 3370
			w$(1) = s$(3)
			w$(2) = s$(6)
				d1 = 1 : rem some A is not B
				goto 3390
3130 if s$(2) <> "no" then 3210
		if t(3) <> 6 then 3350
		if t(4) <> 5 then 3330
		if t(5) <> 6 then 3370
		w$(1) = s$(3)
		w$(2) = s$(5)
		d1 = 3 : rem no A is B
		goto 3390
3210 if t(2) <> 6 then 3350
	if t(3) <> 5 then 3330
	w$(1) = s$(2)
	if s$(4) = "not" then 3290
		if t(4) <> 6 then 3370
		d1 = 4 : rem a is T
		w$(2) = s$(4)
		goto 3390
3290 if t(5) <> 6 then 3370
		d1 = 5 : rem a is not T
		w$(2) = s$(5)
		goto 3390
3330 print "** Missing copula is/are"
	goto 3380
3350 print "** Subject term bad or missing"
	goto 3380
3370 print "** Predicate term bad or missing"
3380 if msg then print "Enter SYNTAX for help with statements"
3390 return
3400 rem---Add W$(1), W$(2) to table T$()---
	if (d1 mod 2) = 0 then 3440
		n1 = n1+1
		if n1 > 1 and msg then print "Warning: ";n1;" negative premises"
3440 e(1) = 0
	for j = 1 to 2
	w$ = w$(j)
	if d1 < 4 then g = 1 :  else if j = 1 then g = 2 :  else g = p1
	gosub 4040
	i1 = 1
3500 gosub 3950
	if i1 <= l1 then 3550
		if b1 > 0 then i1 = b1 : else l1 = l1+1
		t$(i1) = w$
		goto 3720
3550 if g = 0 then 3660
	if g(i1) = 0 then 3620
	if g = g(i1) then 3730
		if not msg then 3600
			print "Warning: ";g$(g);" '";w$;"' has also occurred as a ";g$(3-g)
3600	i1 = i1+1
	goto 3500
3620 if not msg then 3710
		print "Note: earlier use of '";w$;"' taken as the ";g$(g);" used here"
	goto 3710
3660 if g(i1) = 0 and  not msg then 3730
		print "Note: predicate term '";w$;"'";
		print " taken as the ";g$(g(i1));" used earlier"
	goto 3730
3710 if g = 2 then d(i1) = o(i1)
3720 g(i1) = g
3730 if e(j) > 0 then 3770
		if b(i1) > 0 or w$ = w$(j) then 3780
			a$ = left$(w$,1)
	if a$ = "a" or a$ = "e" or a$ = "i" or a$ = "o" or a$ = "u" then e(j) = 2 : else e(j) = 1
3770		b(i1) = e(j)
3780 o(i1) = o(i1)+1
	if o(i1) < 3 then 3810
	if not msg then 3810
		print "Warning: ";g$(g(i1));" '";w$;"' has occurred";o(i1);"times"
3810 if j = 2 then 3850
		p(a1) = i1
		if d1 >= 2 then d(i1) = d(i1)+1
		goto 3900
3850 q(a1) = i1
		if p(a1) <> q(a1) then 3880
			if msg then print "Warning: same term occurs twice in line ";s$(1)
3880	if g(i1) = 2 then d1 = d1+2
		if d1 = 6 or d1 mod 2 then d(i1) = d(i1)+1
3900 if o(i1) <> 2 or d(i1) > 0 then 3920
		if msg then print "Warning: undistributed middle term '";t$(i1);"'"
3920 next j
	r(a1) = d1
	return
3950 rem---Search T$() for W$ from I1 to L1---
	rem If found, I1 = L1; else I1 = L1+1. B1 set to 1st empty loc.
	b1 = 0
3980 if i1 > l1 then 4030
	if t$(i1) = w$ then 4030
		if o(i1) = 0 and b1 = 0 then b1 = i1
		i1 = i1+1
		goto 3980
4030 return
4040 rem---Convert W$ to singular---
	l = len(w$)
	if l < 4 then 4090
		s$ = left$(w$,4)
		if s$ = "the " then 4520
4090 x$ = ""
	i = 1
	n = 1
4120 if i > l then 4510
		s$ = mid$(w$,i,1)
		if s$ <> " " then 4170
			i = i+1
			goto 4120
4170 m = 1
4180 if i+m > l then 4230
		s$ = mid$(w$,i+m,1)
		if s$ = " " then 4230
			m = m+1
			goto 4180
4230 s$ = mid$(w$,i,m)
	y$ = s$
	k = 1
4260 if y$ <> u$(k) then 4290
		y$ = v$(k)
		goto 4470
4290	k = k+1
		if k <= u1 then 4260
4302 if len(y$) < 3 then 4310
		if right$(y$,3) <> "men" then 4310
		y$ = left$(y$,len(y$)-2)+"an"
		goto 4470
4310 l$ = right$(y$,1)
4320 if l$ <> "s" then 4470
		if len(y$) > 1 then l$ = right$(y$,2) :  else goto 4470
		if l$ = "ss" or l$ = "us" or l$ = "is" or l$ = "'s" then 4470
			y$ = left$(y$,len(y$)-1)
			if len(y$) > 1 then l$ = right$(y$,2) :  else goto 4470
4370 if l$ <> "xe" then 4400
		y$ = left$(y$,len(y$)-1)
		goto 4470
4400 if l$ <> "ie" or len(y$) <= 3 then 4440
		y$ = left$(y$,len(y$)-2)
		y$ = y$+"y"
		goto 4470
4440 if len(y$) > 2 then l$ = right$(y$,3) :  else goto 4470
	if l$ <> "sse" and l$ <> "she" and l$ <> "che" then 4470
		y$ = left$(y$,len(y$)-1)
4470 if len(x$) = 0 then x$ = y$ :  else x$ = x$+" "+y$
	n = n+1
	i = m+i
	goto 4120
4510 w$ = x$
4520 return
4530 rem---Enter line into list---
	n = val(s$(1))
	s = len(s$(1))+1
	l = len(l1$)
	l$ = mid$(l1$,s+1,l-s)
	i = 0
4590 j1 = l(i)
	if j1 = 0 then 4690
		if n <> n(j1) then 4660
			gosub 4890
			l$(j1) = l$
			a1 = j1
			goto 4750
4660	if n < n(j1) then 4690
			i = j1
			goto 4590
4690 a1 = a(a(0))
	l$(a1) = l$
	n(a1) = n
	l(i) = a1
	l(a1) = j1
	a(0) = a(0)+1
4750 return
4760 rem---Delete a line---
	n = val(s$(1))
	i = 0
4790 j1 = l(i)
		if j1 = 0 then print "Line ";n;" not found" : goto 4880
			if n = n(j1) then 4840
				i = l(i)
				goto 4790
4840		a(0) = a(0)-1
			a(a(0)) = j1
			l(i) = l(j1)
			gosub 4890
4880 return
4890 rem---Decrement table entries---
	j(1) = p(j1)
	j(2) = q(j1)
	if r(j1) mod 2 = 0 then 4960
		n1 = n1-1
		j(4) = 1
		goto 4970
4960 if g(q(j1)) = 2 then j(4) = 1 : else j(4) = 0
4970 if r(j1) >= 2 then j(3) = 1 : else j(3) = 0
	for k = 1 to 2
		o(j(k)) = o(j(k))-1
		if o(j(k)) > 0 then 5040
			t$(j(k)) = ""
			b(j(k)) = 0
			g(j(k)) = 0
5040	d(j(k)) = d(j(k))-j(k+2)
	next k
	return
5070 rem---See if syllogism---
	j1 = 0
	v1 = 0 : rem flag for modern validity
	if l(0) then 5140
		j1 = 1 : goto 5870
5140 c = 0
	for i = 1 to l1
		if o(i) = 0 or o(i) = 2 then 5250
			if o(i) <> 1 then 5210
				c = c+1
				c(c) = i
				goto 5250
5210		if j1 = 2 then 5240
				print "Not a syllogism:"
				j1 = 2
5240		print "   ";g$(g(i));" '";t$(i);"' occurs ";o(i);" times in premises."
5250	next i
	if c = 2 then 5360
		print "Not a syllogism:"
		j1 = 3
		if c > 0 then 5320
			print "   no terms occur exactly once in premises."
			goto 5360
5320	print "   ";c;" terms occur exactly once in premises."
	for i = 1 to c
		print left$(tb$,6);t$(c(i));" -- ";g$(g(c(i)))
	next i
5360 if j1 then 5870
	i = l(0)
	l = 0
5390	l = l+1
		k(l) = i
		i = l(i)
		if i then 5390
	if l = 1 then 5750
	if d(c(1)) = 0 and d(c(2)) = 1 then t = c(2) : else t = c(1)
	i = 1
5460 k = i
5470	if p(k(k)) <> t then 5500
			t = q(k(k))
			goto 5520
5500	if q(k(k)) <> t then 5620
			t = p(k(k))
5520		if k = i then 5610
				n = 1
				h(1) = k(i)
				for m = i to k-1
					n = 3-n
					h(n) = k(m+1)
					k(m+1) = h(3-n)
					next m
				k(i) = h(n)
5610		if j1 then 5710 else goto 5730
5620	k = k+1
		if k <= l then 5470
		t = q(k(i))
		if j1 > 0 then 5700
			j1 = 4
			print "Not a syllogism: no way to order premises so that each premise"
			print "shares exactly one term with its successor; there is a"
5700		print "closed loop in the term chain within the premise set--"
5710		print n(k(i));
			print l$(k(i))
5730		i = i+1
			if i <= l then 5460
5750 if j1 > 0 then 5870
	 if l1$ <> "link" and l1$ <> "link*" then 5870
		print "Premises of syllogism in order of term links:"
		for i = 1 to l
			print n(k(i));" ";
				if l1$ = "link" then 5850
				if r(k(i)) < 6 and g(q(k(i))) = 2 then r(k(i)) = r(k(i))+2
				if r(k(i)) < 4 then print x$(r(k(i)));"  ";
				print t$(p(k(i)));y$(r(k(i)));"  ";t$(q(k(i)));z$(r(k(i)))
				goto 5860
5850			print l$(k(i))
5860	next i
5870 return
5880 rem---See if conclusion possible---
	c1 = c(1)
	c2 = c(2)
	for i = 1 to l1
		if o(i) < 2 then 6000
			if d(i) > 0 then 5980
				if j1 > 0 then 5970
					print "Undistributed middle terms:"
					j1 = 5
5970			print left$(tb$,5);t$(i)
5980		if d(i) = 1 or g(i) = 2 then 6000
				v1 = i
6000	next i
	if n1 < 2 then 6040
		j1 = 6
		print "More than one negative premise:"
6040 if j1 > 0 then 6180
	if n1 = 0 then 6190
	if d(c1) > 0 or d(c2) > 0 then 6100
		print "Terms '";t$(c1);"' and '";t$(c2);"',";" one of which is"
		goto 6150
6100 if d(c1) > 0 or g(c2) < 2 then 6130
		print "Term '";t$(c1);"'"
		goto 6150
6130 if d(c2) > 0 or g(c1) < 2 then 6190
		print "Term '";t$(c2);"'"
6150 print "required in predicate of negative conclusion"
	print "not distributed in the premises."
	j1 = 7
6180 print "No possible conclusion."
6190 return
6200 rem---Compute conclusion---
	if l(0) = 0 then z$ = "A is A" : goto 6580
	if n1 = 0 then 6400
	rem negative conclusion
	if d(c2) > 0 then 6260
		z$ = "Some "+t$(c2)+" is not "+a$(b(c1))+t$(c1)
		goto 6390
6260 if d(c1) > 0 then 6290
		z$ = "Some "+t$(c1)+" is not "+a$(b(c2))+t$(c2)
		goto 6390
6290 if g(c1) < 2 then 6320
		z$ = t$(c1)+" is not "+a$(b(c2))+t$(c2)
		goto 6390
6320 if g(c2) < 2 then 6350
		z$ = t$(c2)+" is not "+a$(b(c1))+t$(c1)
		goto 6390
6350 if b(c1) > 0 or b(c2) = 0 then 6380
		z$ = "No "+t$(c2)+" is "+a$(b(c1))+t$(c1)
		goto 6390
6380	z$ = "No "+t$(c1)+" is "+a$(b(c2))+t$(c2)
6390 goto 6570
6400 rem affirmative conclusion
	if d(c1) = 0 then 6470
		if g(c1) = 2 then 6450
			z$ = "All "+t$(c1)+" is "+t$(c2)
			goto 6570
6450		z$ = t$(c1)+" is "+a$(b(c2))+t$(c2)
			goto 6570
6470 if d(c2) = 0 then 6530
		if g(c2) = 2 then 6510
			z$ = "All "+t$(c2)+" is "+t$(c1)
			goto 6570
6510		z$ = t$(c2)+" is "+a$(b(c1))+t$(c1)
			goto 6570
6530 if b(c1) > 0 or b(c2) = 0 then 6560
		z$ = "Some "+t$(c2)+" is "+a$(b(c1))+t$(c1)
		goto 6570
6560	z$ = "Some "+t$(c1)+" is "+a$(b(c2))+t$(c2)
6570 rem PRINT  conclusion
6580 print "  / ";z$
	if v1 = 0 then 6620
		print "  * Aristotle-valid only, i.e. on requirement that term ";
		print "'";t$(v1);"' denotes."
6620 return
6630 rem---test offered conclusion---
	rem--conc. poss, line in s$()
	gosub 2890
	if d1 < 0 then 7370
	if d1 < 4 then g1 = 1 : g2 = 1 : else g1 = 2 : g2 = p1
	if g2 = 2 and d1 < 6 and d1 > 3 then d1 = d1+2
	w$ = w$(1)
	gosub 4040
	if j1 = 0 then 6750
		w$(1) = w$
		goto 6840
6750 for j = 1 to 2
		if w$ <> t$(c(j)) then 6810
			if g(c(j)) > 0 then 6800
				print "Note: '";t$(c(j));"' used in premises taken to be ";g$(g1)
				goto 6840
6800		if g1 = g(c(j)) then 6840
6810	next j
	print "** Conclusion may not contain ";g$(g1);" '";w$;"'."
	j = 0
6840 w$ = w$(2)
	gosub 4040
	if j1 = 0 then 6940
		if w$ = w$(1) then 6910
6880		print "** Conclusion from no premises must have same subject and predic";
				print "ate."
			goto 7370
6910	if d1 <> 4 or g2 = 0 then 7120
			print "** Subject is a ";g$(2);", predicate is a ";g$(1);" -- but"
			goto 6880
6940 if j > 0 then 6970
		if w$ = t$(c(1)) then t2 = c(2) : else t2 = c(1)
		goto 7070
6970 t1 = c(j)
	t2 = c(3-j)
	if w$ <> t$(t2) then 7060
		if g(t2) > 0 then 7040
			if g2 = 0 then 7090
			print "Note: '";t$(t2);"' used in premises taken to be ";g$(g2)
			goto 7090
7040	if g2 = 0 then 7090
		if g2 = g(t2) then 7090
7060 print "** Conclusion may not contain ";g$(g2);" '";w$;"';"
7070 print "** Conclusion must contain ";g$(g(t2));" '";t$(t2);"'."
	 goto 7370
7090 if n1 = 0 or (d1 mod 2) = 1 then 7120
		print "** Negative conclusion required."
		goto 7370
7120 if n1 > 0 or d1 mod 2 = 0 then 7150
		print "** Affirmative conclusion required."
		goto 7370
7150 if j1 = 1 then 7250
	if d(t1) > 0 or d1 <= 1 or d1 >= 4 then 7200
		print "** Term '";t$(t1);"' not distributed in premises"
7180	print "   may not be distributed in conclusion."
		goto 7370
7200 if d(t2) > 0 then 7250
		if d1 mod 2 = 0 and d1 <> 6 then 7250
		print "** Term '";t$(t2);"' not distributed in premises"
		goto 7180
7250 print "-->  VALID!"
	if j1 = 0 then 7300
		if d1 > 0 then 7370
		t$(0) = w$
		goto 7350
7300 if d(t1) = 0 or d1 >= 2 then 7330
		v1 = t1
		goto 7350
7330 if d(t2) > 0 and d1 mod 2 = 0 and d1 <> 4 and d1 <> 6 then v1 = t2
	if v1 = 0 then 7370
7350 print "    but on Aristotelian interpretation only, i.e. on requirement"
	print "    that term '";t$(v1);"' denotes."
7370 return
7460 rem---list---
	i = 0
7540 i = l(i)
	if i = 0 then 7650
	print n(i);" ";
	if l1$ = "list" then 7630
		if r(i) < 6 and g(q(i)) = 2 then r(i) = r(i)+2
		if r(i) < 4 then print x$(r(i));"  ";
		print t$(p(i));y$(r(i));"  ";t$(q(i));z$(r(i))
		goto 7540
7630 print l$(i)
	goto 7540
7650 return
7660 rem---List valid inputs---
	cls : print "Valid commands are:"
	print "   <n>  [ <statement> ]   Insert, delete, or replace premise number  <n> "
	print left$(tb$,28);"Examples:   10  All men are mortal"
	print left$(tb$,40);"10"
	print "  DUMP";left$(tb$,15);"Prints symbol table, distribution count, etc."
	print "  HELP";left$(tb$,15);"Prints this list"
	print "  INFO";left$(tb$,15);"Gives information about syllogisms"
	print "  LIST";left$(tb$,15);"Lists premises"
	print "  LIST*";left$(tb$,14);"Same, but displays distribution analysis:"
	print left$(tb$,25);"distributed positions marked with '*', "
	print left$(tb$,25);"designators marked with '+'"
	print "  LINK";left$(tb$,15);"Lists premises in order of term-links (if possible)"
	print "  LINK*";left$(tb$,14);"Same, but in distribution-analysis format"
	print "  MSG";left$(tb$,16);"Turns on/off Printing of certain messages and warnings"
	print "  NEW";left$(tb$,16);"Erases current syllogism"
	print "  SAMPLE";left$(tb$,13);"Erases current syllogism and enters sample syllogism"
	print "  STOP";left$(tb$,15);"Stops entire program"
	print "  SUBSTITUTE";left$(tb$,9);"Allows uniform substitution of new terms in ";
	print "old premises"
	print "  SYNTAX";left$(tb$,13);"Explains statement syntax, with examples"
	print "  /";left$(tb$,18);"Asks program to draw conclusion"
	print "  /  <statement>";left$(tb$,5);"Tests  <statement>  as conclusion"
	print left$(tb$,25);"Note: this can be done even if there are no premises"
	return
7960 rem--"syntax"--
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
	return
8290 rem---Info---
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
	return
8890 rem---"Dump" values of variables---
	print "Highest symbol table loc. used:";l1;"  Negative premises:";n1
	if l1 = 0 then 8970
	print "Adr. art. term";left$(tb$,48-14);"type       occurs    dist. count"
	for i = 1 to l1
		rem Metal's lack of tabbing gets difficult here...
		itab = 7-len(str$(i))
		astringtab = 11-len(a$(b(i)))-7
		tstringtab = 49-len(t$(i))-11
		gtab = 60-len(str$(g(i)))-49
		otab = 71-len(str$(o(i)))-60
		print i;left$(tb$,itab);a$(b(i));left$(tb$,astringtab);t$(i);left$(tb$,tstringtab);g(i);left$(tb$,gtab);
		print o(i);left$(tb$,otab);d(i)
		next i
8970 return
8980 rem--sample--
	for z8 = 1 to 10 : read l1$ : print l1$
	gosub 2020 : gosub 2890 : gosub 4530 : gosub 3400
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
9070 print "Enter address of old term; or 0 for help, -1 to exit, -2 for dump"
	input i1
	if i1 = -1 then 9470
	if i1 <> -2 then 9130
		gosub 8890
		goto 9070
9130 if i1 > 0 then 9340
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
	goto 9455
9340 if i1 <= l1 then 9370
		print "Address ";i1;" too large.  Symbol table only of length ";l1
		goto 9455
9370 print "Enter new term to replace ";g$(g(i1));" '";t$(i1);"'"
		input w$
		t$(i1) = w$
		print "Replaced by '";w$;"'"
9455 print
		goto 9070
9470 print "Exit from substitution routine"
	return
end
