1 rem Syllogism 1.0. November 8, 2002
2 rem I edited this program in 2002, for compatibility with freeware BASIC
4 rem interpreters for the Mac: Chipmunk BASIC 3.5.7 and Metal BASIC 1.7.3. 
6 rem I hope compatibility with two indpendently developed BASICs should
8 rem assure some universality regardless of platform. The following
10 rem notes should help anybody with a similar project.
12 rem Summary of changes.....
14 rem * Metal doesn't support PRINT TAB(N). It supports the command HTAB, but a
16 rem bug makes it useless for formatting more than one column of text. The only
18 rem standard BASIC solution is the " " character, implemented with tb$ and Left$()
20 rem * Metal doesn't support MID$() as a command. Both support LCASE$().
22 rem * Metal crashes when it reads an empty ("") DATA item. Had to hack it.
24 rem * Metal requires ; between PRINT items. Also, it doesn't add a space
26 rem when the ; separates a number and a string.
28 rem Chipmunk buggy with IF...ELSE on one line--use colons (and GOTOs)
30 rem Chipmunk requires quotes around DATA items
32 rem Chipmunk buggy when using integer variables (J%), use floating point (J)
34 rem Neither support PRINT CHR$(27) as a way to clear the screen; used cls
36 rem Peace. Ben Sharvy. luvnpeas99@yahoo.com
100 rem Metal doesn't support PRINT TAB(N)...
102 tb$ = "                                                  "
110 cls : print "Syllogism Program Copyright (c) 1988 Richard Sharvy"
112 print "Syllogism 1.0 (c) 2002 Richard Sharvy's estate"
113 print "Ben Sharvy: luvnpeas99@yahoo.com or bsharvy@efn.org" : print
480 dim a(63),c(63),d(63),g(63),l(63),n(63),o(63),p(63),q(63)
490 dim r(63),b(63),k(63),j(4),t(7),e(2),h(2)
500 dim a$(3),l$(63),t$(65)
510 dim g$(2),s$(6),w$(2),x$(7),y$(7),z$(7)
520 read a$(1),a$(2),a$(3),g$(0),g$(1),g$(2)
530 data "a ","an ","sm ","undetermined type","general term","designator"
550 for i = 0 to 7
551 read x$(i),y$(i),z$(i)
552 rem Metal crashes when "" is in DATA statement...
554 if x$(i) = "<null>" then x$(i) = ""
555 if y$(i) = "<null>" then y$(i) = ""
556 if z$(i) = "<null>" then z$(i) = ""
560 next i
570 data "some","  is","<null>","some","  is not","*"
580 data "all","*  is","<null>","no","*  is","*"
590 data "<null>","+  is","<null>","<null>","+  is not","*"
600 data "<null>","+  = ","+","<null>","+   = / = ","*"
602 rem /error check/ for err = 0 to 7 : print x$(err),y$(err),z$(err) : next err
740 dim u$(75),v$(75)
760 i = 0
770 i = i+1
771  read u$(i),v$(i)
780  if u$(i) <> "ZZZZZ" then 770
790 u1 = i-1
820 data "socrates","socrates","parmenides","parmenides","epimenides","epimenides"
840 data "mice","mouse","lice","louse","geese","goose"
850 data "children","child","oxen","ox","people","person","teeth","tooth"
860 data "wolves","wolf","wives","wife","selves","self","lives","life","leaves","leaf"
870 data "shelves","shelf","elves","elf","dwarves","dwarf","knives","knife","thieves","thief"
880 data "neckties","necktie","hippies","hippie","yippies","yippie","yuppies","yuppie"
890 data "moonies","moonie","druggies","druggie","cookies","cookie","commies","commie"
891 data "groupies","groupie","tomatoes","tomato"
910 data "alcibiades","alcibiades","thales","thales","aries","aries","athens","athens"
920 data "species","species","feces","feces","geniuses","genius","sorites","sorites"
930 data "crises","crisis","emphases","emphasis","memoranda","memorandum","theses","thesis"
940 data "automata","automaton","formulae","formula","stigmata","stigma","lemmata","lemma"
950 data "vertices","vertex","vortices","vortex","indices","index","codices","codex"
960 data "matrices","matrix"
970 data "gasses","gas","gases","gas","buses","bus","aches","ache","headaches","headache"
980 data "grits","grits","molasses","molasses","gas","gas","christmas","christmas"
990 data "mathematics","mathematics","semantics","semantics","physics","physics"
1000 data "metaphysics","metaphysics","ethics","ethics","linguistics","linguistics"
1010 data "kiwis","kiwi","israelis","israeli"
1020 data "goyim","goy","seraphim","seraph","cherubim","cherub"
1025 data "semen","semen","amen","amen"
1030 data "ZZZZZ","ZZZZZ"
1050 msg = -1 : a(0) = 1 : for i = 1 to 63 : a(i) = i : next i
1070 if msg then print "Enter HELP for list of commands"
1080 rem---Input line---
1085 print ">";
1090 line input l1$
1100 l = len(l1$)
1110 if l = 0 then 1070
1120 l2$ = right$(l1$,1)
1121  if l2$ = " " then 1160
1130   if l2$ <> "." and l2$ <> "?" and l2$ <> "!" then 1181
1140   print left$(tb$,l);" ^   Punctuation mark ignored"
1160  if l = 1 then 1080
1165  l = l-1 : l1$ = left$(l1$,l) : goto 1120
1181 if left$(l1$,1) <> " " then 1190
1182  if l = 1 then 1080
1184  l = l-1 : l1$ = right$(l1$,l) : goto 1181
1190 rem / FOR I = 1 TO L
1191  rem / V = ASC(MID$(L1$,I,1))
1192  rem / IF V >= 65 AND V <= 90  THEN  MID$(L1$,I,1) = CHR$(V+32)
1193  rem / NEXT
1194 rem Metal doesn't support mid$ as command, but lcase$() is well supported...
1195 l1$ = lcase$(l1$)
1220 if l1$ <> "stop" then 1230
1223  if msg then print "(Some versions support typing CONT to continue)"
1224 stop
1227  goto 1080
1230 if l1$ <> "new" then 1270
1240  print "Begin new syllogism"
1250  gosub 1840
1260  goto 1080
1270 if l1$ <> "sample" then 1310
1280  gosub 1840
1290  gosub 8980
1300  goto 1080
1310 if l1$ <> "help" then 1340
1320  gosub 7660
1330  goto 1080
1340 if l1$ <> "syntax" then 1370
1350  gosub 7960
1360  goto 1080
1370 if l1$ <> "info" then 1430
1380  gosub 8290
1390  goto 1080
1430 if l1$ <> "dump" then 1460
1440  gosub 8890
1450  goto 1080
1460 if l1$ <> "msg" then 1500
1465  msg = not (msg)
1470  print "Messages turned ";
1475  if msg then print "on" : else print "off"
1480  goto 1080
1500 if l1$ <> "substitute" then 1540
1510  if l(0) = 0 then 1612
1520  gosub 9060
1530  goto 1080
1540 if l1$ <> "link" and l1$ <> "link*" then 1561
1545  if l(0) = 0 then 1612
1550  gosub 5070
1560  goto 1080
1561 if l1$ <> "list" and l1$ <> "list*" then 1570
1562  if l(0) = 0 then 1612
1563  gosub 7460
1564  goto 1080
1570 rem--scan line L1$ into array S$()
1575 gosub 2020
1580 if t(1) <> 1 then 1745
1600  if t(2) then 1640
1610   if l(0) then 1620
1612    print "No premises" : goto 1080
1620   gosub 4760 : rem delete line
1630   goto 1080
1640  gosub 2890 : rem parse the line in S$()
1650  if d1 < 0 then 1080
1660  gosub 4530 : rem enter line into list
1670  gosub 3400 : rem add terms to symbol table
1680  goto 1080
1745 if t(1) = 0 then 1070
1750 rem draw/test conclusion
1755 gosub 5070 : rem is it a syl?
1760 if j1 > 1 then 1080
1770 if j1 = 0 then gosub 5880 : rem poss. conclusion?
1810  if j1 > 1 then 1080
1820  if t(2) then gosub 6630 : else gosub 6200 : rem test/draw conclusion
1830 goto 1080
1840 rem---New---
1850 if l(0) = 0 then 2010
1860 for i = 1 to l1
1870  d(i) = 0 : t$(i) = "" : b(i) = 0 : o(i) = 0 : g(i) = 0
1920  next i
1930 l1 = 0 : n1 = 0
1950 j = l(0)
1960 a(0) = a(0)-1
1970  a(a(0)) = j
1980  j = l(j)
1990  if j > 0 then 1960
2000 l(0) = 0
2010 return
2020 rem---L1$ into array S$()---
2030 rem T(): 1:line num., 2:"/", 3:quantifier, 4:no/not, 5:is/are, 6:term
2040 rem                     10 SOME  FRIED COCONUTS   ARE  NOT  TASTY
2050 rem                      1   3        6            5    4     6
2080 for j = 1 to 6 : s$(j) = "" : t(j) = 0 : next j
2120 p1 = 0 : e(2) = 0 : j = 1 : i = 1
2160 l = len(l1$)
2180 if i > l then 2885
2190  s$ = mid$(l1$,i,1)
2200  if s$ = " " then i = i+1 : goto 2180
2230 k = 1
2240 if i+k > l then 2290
2250  s$ = mid$(l1$,i+k,1)
2260  if s$ <> " " then k = k+1 : goto 2240
2290 s$ = mid$(l1$,i,k) : rem S$ is set to next word
2310 if j > 1 then 2520
2320  if s$ <> "/" then 2400
2360  t(1) = 2
2370  goto 2840
2400 n = len(s$)
2410 if n > 4 then 2460
2420  n = 1
2430  t$ = mid$(s$,n,1)
2440  if asc(t$) <= 57 and asc(t$) >= 48 then 2480
2460   print left$(tb$,i+n);"^   Invalid numeral or command"
2470   goto 2885
2480  n = n+1
2490  if n <= len(s$) then 2430
2500 t(1) = 1
2510 goto 2840
2520 rem Scan
2525 if s$ = "somebody" or s$ = "something" or s$ = "nobody" or s$ = "nothing" then 2670
2527 if s$ = "someone" or s$ = "everyone" or s$ = "everybody" or s$ = "everything" then 2670
2530 if s$ <> "all" and s$ <> "some" then 2570
2540  if t(j) = 6 then 2670
2550  t(j) = 3
2560  goto 2840
2570 if s$ <> "no" and s$ <> "not" then 2610
2580  if t(j) = 6 then 2670
2590  t(j) = 4
2600  goto 2840
2610 if s$ <> "is" and s$ <> "are" then 2710
2620  if t(j) <> 6 then 2670
2630  if t(j-1) = 5 or t(j-2) = 5 then 2670
2640   j = j+1
2650   t(j) = 5
2660   goto 2840
2670  print left$(tb$,i+k-1);"^"
2680  print "Reserved word '";s$;"' may not occur within a term"
2690  t(1) = 0
2700  goto 2885
2710 if t(j) = 6 then 2820
2720 if t(j-1) <> 5 and t(j-2) <> 5 then 2790
2730 if s$ <> "a" and s$ <> "an" and s$ <> "sm" then 2780
2740  if i = l then 2790
2750  if s$ = "a" then e(2) = 1 :  else if s$ = "an" then e(2) = 2 :  else e(2) = 3
2760  p1 = 1
2770  goto 2860
2780 if s$ = "the" then p1 = 2
2790 s$(j) = s$
2800 t(j) = 6
2810 goto 2860
2820 s$(j) = s$(j)+" "+s$
2830 goto 2860
2840 s$(j) = s$
2850 j = j+1
2860 i = k+i
2870 if j <= 6 then 2180
2885 return
2890 rem---Parse line in S$()---
2900 d1 = -1
2910 if s$(2) <> "all" then 2990
2920  if t(3) <> 6 then 3350
2930  if t(4) <> 5 then 3330
2940  if t(5) <> 6 then 3370
2950  w$(1) = s$(3)
2960  w$(2) = s$(5)
2970  d1 = 2 : rem all A is B
2980  goto 3390
2990 if s$(2) <> "some" then 3130
3000  if t(3) <> 6 then 3350
3010  if t(4) <> 5 then 3330
3020  if s$(5) = "not" then 3080
3030   if t(5) <> 6 then 3370
3040   w$(1) = s$(3)
3050   w$(2) = s$(5)
3060   d1 = 0 : rem Some A is B
3070   goto 3390
3080  if t(6) <> 6 then 3370
3090   w$(1) = s$(3)
3100   w$(2) = s$(6)
3110    d1 = 1 : rem some A is not B
3120    goto 3390
3130 if s$(2) <> "no" then 3210
3140  if t(3) <> 6 then 3350
3150  if t(4) <> 5 then 3330
3160  if t(5) <> 6 then 3370
3170  w$(1) = s$(3)
3180  w$(2) = s$(5)
3190  d1 = 3 : rem no A is B
3200  goto 3390
3210 if t(2) <> 6 then 3350
3220 if t(3) <> 5 then 3330
3230 w$(1) = s$(2)
3240 if s$(4) = "not" then 3290
3250  if t(4) <> 6 then 3370
3260  d1 = 4 : rem a is T
3270  w$(2) = s$(4)
3280  goto 3390
3290 if t(5) <> 6 then 3370
3300  d1 = 5 : rem a is not T
3310  w$(2) = s$(5)
3320  goto 3390
3330 print "** Missing copula is/are"
3340 goto 3380
3350 print "** Subject term bad or missing"
3360 goto 3380
3370 print "** Predicate term bad or missing"
3380 if msg then print "Enter SYNTAX for help with statements"
3390 return
3400 rem---Add W$(1), W$(2) to table T$()---
3410 if (d1 mod 2) = 0 then 3440
3420  n1 = n1+1
3430  if n1 > 1 and msg then print "Warning: ";n1;" negative premises"
3440 e(1) = 0
3450 for j = 1 to 2
3460 w$ = w$(j)
3470 if d1 < 4 then g = 1 :  else if j = 1 then g = 2 :  else g = p1
3480 gosub 4040
3490 i1 = 1
3500 gosub 3950
3510 if i1 <= l1 then 3550
3520  if b1 > 0 then i1 = b1 : else l1 = l1+1
3530  t$(i1) = w$
3540  goto 3720
3550 if g = 0 then 3660
3570 if g(i1) = 0 then 3620
3580 if g = g(i1) then 3730
3590  if not msg then 3600
3592   print "Warning: ";g$(g);" '";w$;"' has also occurred as a ";g$(3-g)
3600  i1 = i1+1
3610 goto 3500
3620 if not msg then 3710
3630  print "Note: earlier use of '";w$;"' taken as the ";g$(g);" used here"
3640 goto 3710
3660 if g(i1) = 0 and  not msg then 3730
3670  print "Note: predicate term '";w$;"'";
3680  print " taken as the ";g$(g(i1));" used earlier"
3690 goto 3730
3710 if g = 2 then d(i1) = o(i1)
3720 g(i1) = g
3730 if e(j) > 0 then 3770
3740  if b(i1) > 0 or w$ = w$(j) then 3780
3750   a$ = left$(w$,1)
3760 if a$ = "a" or a$ = "e" or a$ = "i" or a$ = "o" or a$ = "u" then e(j) = 2 : else e(j) = 1
3770   b(i1) = e(j)
3780 o(i1) = o(i1)+1
3790 if o(i1) < 3 then 3810
3800 if not msg then 3810
3805  print "Warning: ";g$(g(i1));" '";w$;"' has occurred";o(i1);"times"
3810 if j = 2 then 3850
3820  p(a1) = i1
3830  if d1 >= 2 then d(i1) = d(i1)+1
3840  goto 3900
3850 q(a1) = i1
3860  if p(a1) <> q(a1) then 3880
3870   if msg then print "Warning: same term occurs twice in line ";s$(1)
3880  if g(i1) = 2 then d1 = d1+2
3890  if d1 = 6 or d1 mod 2 then d(i1) = d(i1)+1
3900 if o(i1) <> 2 or d(i1) > 0 then 3920
3910  if msg then print "Warning: undistributed middle term '";t$(i1);"'"
3920 next j
3930 r(a1) = d1
3940 return
3950 rem---Search T$() for W$ from I1 to L1---
3960 rem If found, I1 = L1; else I1 = L1+1. B1 set to 1st empty loc.
3970 b1 = 0
3980 if i1 > l1 then 4030
3990 if t$(i1) = w$ then 4030
4000  if o(i1) = 0 and b1 = 0 then b1 = i1
4010  i1 = i1+1
4020  goto 3980
4030 return
4040 rem---Convert W$ to singular---
4050 l = len(w$)
4060 if l < 4 then 4090
4070  s$ = left$(w$,4)
4080  if s$ = "the " then 4520
4090 x$ = ""
4100 i = 1
4110 n = 1
4120 if i > l then 4510
4130  s$ = mid$(w$,i,1)
4140  if s$ <> " " then 4170
4150   i = i+1
4160   goto 4120
4170 m = 1
4180 if i+m > l then 4230
4190  s$ = mid$(w$,i+m,1)
4200  if s$ = " " then 4230
4210   m = m+1
4220   goto 4180
4230 s$ = mid$(w$,i,m)
4240 y$ = s$
4250 k = 1
4260 if y$ <> u$(k) then 4290
4270  y$ = v$(k)
4280  goto 4470
4290  k = k+1
4291  if k <= u1 then 4260
4302 if len(y$) < 3 then 4310
4304  if right$(y$,3) <> "men" then 4310
4306  y$ = left$(y$,len(y$)-2)+"an"
4308  goto 4470
4310 l$ = right$(y$,1)
4320 if l$ <> "s" then 4470
4330  if len(y$) > 1 then l$ = right$(y$,2) :  else goto 4470
4340  if l$ = "ss" or l$ = "us" or l$ = "is" or l$ = "'s" then 4470
4350   y$ = left$(y$,len(y$)-1)
4360   if len(y$) > 1 then l$ = right$(y$,2) :  else goto 4470
4370 if l$ <> "xe" then 4400
4380  y$ = left$(y$,len(y$)-1)
4390  goto 4470
4400 if l$ <> "ie" or len(y$) <= 3 then 4440
4410  y$ = left$(y$,len(y$)-2)
4420  y$ = y$+"y"
4430  goto 4470
4440 if len(y$) > 2 then l$ = right$(y$,3) :  else goto 4470
4450 if l$ <> "sse" and l$ <> "she" and l$ <> "che" then 4470
4460  y$ = left$(y$,len(y$)-1)
4470 if len(x$) = 0 then x$ = y$ :  else x$ = x$+" "+y$
4480 n = n+1
4490 i = m+i
4500 goto 4120
4510 w$ = x$
4520 return
4530 rem---Enter line into list---
4540 n = val(s$(1))
4550 s = len(s$(1))+1
4560 l = len(l1$)
4570 l$ = mid$(l1$,s+1,l-s)
4580 i = 0
4590 j1 = l(i)
4600 if j1 = 0 then 4690
4610  if n <> n(j1) then 4660
4620   gosub 4890
4630   l$(j1) = l$
4640   a1 = j1
4650   goto 4750
4660  if n < n(j1) then 4690
4670   i = j1
4680   goto 4590
4690 a1 = a(a(0))
4700 l$(a1) = l$
4710 n(a1) = n
4720 l(i) = a1
4730 l(a1) = j1
4740 a(0) = a(0)+1
4750 return
4760 rem---Delete a line---
4770 n = val(s$(1))
4780 i = 0
4790 j1 = l(i)
4800  if j1 = 0 then print "Line ";n;" not found" : goto 4880
4810   if n = n(j1) then 4840
4820    i = l(i)
4830    goto 4790
4840   a(0) = a(0)-1
4850   a(a(0)) = j1
4860   l(i) = l(j1)
4870   gosub 4890
4880 return
4890 rem---Decrement table entries---
4900 j(1) = p(j1)
4910 j(2) = q(j1)
4920 if r(j1) mod 2 = 0 then 4960
4930  n1 = n1-1
4940  j(4) = 1
4950  goto 4970
4960 if g(q(j1)) = 2 then j(4) = 1 : else j(4) = 0
4970 if r(j1) >= 2 then j(3) = 1 : else j(3) = 0
4980 for k = 1 to 2
4990  o(j(k)) = o(j(k))-1
5000  if o(j(k)) > 0 then 5040
5010   t$(j(k)) = ""
5020   b(j(k)) = 0
5030   g(j(k)) = 0
5040  d(j(k)) = d(j(k))-j(k+2)
5050  next k
5060 return
5070 rem---See if syllogism---
5080 j1 = 0
5090 v1 = 0 : rem flag for modern validity
5100 if l(0) then 5140
5120  j1 = 1 : goto 5870
5140 c = 0
5150 for i = 1 to l1
5160  if o(i) = 0 or o(i) = 2 then 5250
5170   if o(i) <> 1 then 5210
5180    c = c+1
5190    c(c) = i
5200    goto 5250
5210   if j1 = 2 then 5240
5220    print "Not a syllogism:"
5230    j1 = 2
5240   print "   ";g$(g(i));" '";t$(i);"' occurs ";o(i);" times in premises."
5250  next i
5260 if c = 2 then 5360
5270  print "Not a syllogism:"
5280  j1 = 3
5290  if c > 0 then 5320
5300   print "   no terms occur exactly once in premises."
5310   goto 5360
5320  print "   ";c;" terms occur exactly once in premises."
5330 for i = 1 to c
5340  print left$(tb$,6);t$(c(i));" -- ";g$(g(c(i)))
5350  next i
5360 if j1 then 5870
5370 i = l(0)
5380 l = 0
5390  l = l+1
5400  k(l) = i
5410  i = l(i)
5420  if i then 5390
5430 if l = 1 then 5750
5440 if d(c(1)) = 0 and d(c(2)) = 1 then t = c(2) : else t = c(1)
5450 i = 1
5460 k = i
5470  if p(k(k)) <> t then 5500
5480   t = q(k(k))
5490   goto 5520
5500  if q(k(k)) <> t then 5620
5510   t = p(k(k))
5520   if k = i then 5610
5530    n = 1
5540    h(1) = k(i)
5550    for m = i to k-1
5560     n = 3-n
5570     h(n) = k(m+1)
5580     k(m+1) = h(3-n)
5590     next m
5600    k(i) = h(n)
5610   if j1 then 5710 : else goto 5730
5620  k = k+1
5630  if k <= l then 5470
5640  t = q(k(i))
5650  if j1 > 0 then 5700
5660   j1 = 4
5670   print "Not a syllogism: no way to order premises so that each premise"
5690   print "shares exactly one term with its successor; there is a"
5700   print "closed loop in the term chain within the premise set--"
5710   print n(k(i));
5720   print l$(k(i))
5730   i = i+1
5740   if i <= l then 5460
5750 if j1 > 0 then 5870
5760 if l1$ <> "link" and l1$ <> "link*" then 5870
5770  print "Premises of syllogism in order of term links:"
5780  for i = 1 to l
5790   print n(k(i));" ";
5800    if l1$ = "link" then 5850
5810    if r(k(i)) < 6 and g(q(k(i))) = 2 then r(k(i)) = r(k(i))+2
5820    if r(k(i)) < 4 then print x$(r(k(i)));"  ";
5830    print t$(p(k(i)));y$(r(k(i)));"  ";t$(q(k(i)));z$(r(k(i)))
5840    goto 5860
5850    print l$(k(i))
5860  next i
5870 return
5880 rem---See if conclusion possible---
5890 c1 = c(1)
5900 c2 = c(2)
5910 for i = 1 to l1
5920  if o(i) < 2 then 6000
5930   if d(i) > 0 then 5980
5940    if j1 > 0 then 5970
5950     print "Undistributed middle terms:"
5960     j1 = 5
5970    print left$(tb$,5);t$(i)
5980   if d(i) = 1 or g(i) = 2 then 6000
5990    v1 = i
6000  next i
6010 if n1 < 2 then 6040
6020  j1 = 6
6030  print "More than one negative premise:"
6040 if j1 > 0 then 6180
6050 if n1 = 0 then 6190
6060 if d(c1) > 0 or d(c2) > 0 then 6100
6070  print "Terms '";t$(c1);"' and '";t$(c2);"',";" one of which is"
6090  goto 6150
6100 if d(c1) > 0 or g(c2) < 2 then 6130
6110  print "Term '";t$(c1);"'"
6120  goto 6150
6130 if d(c2) > 0 or g(c1) < 2 then 6190
6140  print "Term '";t$(c2);"'"
6150 print "required in predicate of negative conclusion"
6160 print "not distributed in the premises."
6170 j1 = 7
6180 print "No possible conclusion."
6190 return
6200 rem---Compute conclusion---
6201 if l(0) = 0 then z$ = "A is A" : goto 6580
6210 if n1 = 0 then 6400
6220 rem negative conclusion
6230 if d(c2) > 0 then 6260
6240  z$ = "Some "+t$(c2)+" is not "+a$(b(c1))+t$(c1)
6250  goto 6390
6260 if d(c1) > 0 then 6290
6270  z$ = "Some "+t$(c1)+" is not "+a$(b(c2))+t$(c2)
6280  goto 6390
6290 if g(c1) < 2 then 6320
6300  z$ = t$(c1)+" is not "+a$(b(c2))+t$(c2)
6310  goto 6390
6320 if g(c2) < 2 then 6350
6330  z$ = t$(c2)+" is not "+a$(b(c1))+t$(c1)
6340  goto 6390
6350 if b(c1) > 0 or b(c2) = 0 then 6380
6360  z$ = "No "+t$(c2)+" is "+a$(b(c1))+t$(c1)
6370  goto 6390
6380  z$ = "No "+t$(c1)+" is "+a$(b(c2))+t$(c2)
6390 goto 6570
6400 rem affirmative conclusion
6410 if d(c1) = 0 then 6470
6420  if g(c1) = 2 then 6450
6430   z$ = "All "+t$(c1)+" is "+t$(c2)
6440   goto 6570
6450   z$ = t$(c1)+" is "+a$(b(c2))+t$(c2)
6460   goto 6570
6470 if d(c2) = 0 then 6530
6480  if g(c2) = 2 then 6510
6490   z$ = "All "+t$(c2)+" is "+t$(c1)
6500   goto 6570
6510   z$ = t$(c2)+" is "+a$(b(c1))+t$(c1)
6520   goto 6570
6530 if b(c1) > 0 or b(c2) = 0 then 6560
6540  z$ = "Some "+t$(c2)+" is "+a$(b(c1))+t$(c1)
6550  goto 6570
6560  z$ = "Some "+t$(c1)+" is "+a$(b(c2))+t$(c2)
6570 rem PRINT  conclusion
6580 print "  / ";z$
6590 if v1 = 0 then 6620
6600  print "  * Aristotle-valid only, i.e. on requirement that term ";
6610  print "'";t$(v1);"' denotes."
6620 return
6630 rem---test offered conclusion---
6640 rem--conc. poss, line in s$()
6650 gosub 2890
6660 if d1 < 0 then 7370
6670 if d1 < 4 then g1 = 1 : g2 = 1 : else g1 = 2 : g2 = p1
6690 if g2 = 2 and d1 < 6 and d1 > 3 then d1 = d1+2
6700 w$ = w$(1)
6710 gosub 4040
6720 if j1 = 0 then 6750
6730  w$(1) = w$
6740  goto 6840
6750 for j = 1 to 2
6760  if w$ <> t$(c(j)) then 6810
6770   if g(c(j)) > 0 then 6800
6780    print "Note: '";t$(c(j));"' used in premises taken to be ";g$(g1)
6790    goto 6840
6800   if g1 = g(c(j)) then 6840
6810  next j
6820 print "** Conclusion may not contain ";g$(g1);" '";w$;"'."
6830 j = 0
6840 w$ = w$(2)
6850 gosub 4040
6860 if j1 = 0 then 6940
6870  if w$ = w$(1) then 6910
6880   print "** Conclusion from no premises must have same subject and predic";
6882    print "ate."
6900   goto 7370
6910  if d1 <> 4 or g2 = 0 then 7120
6920   print "** Subject is a ";g$(2);", predicate is a ";g$(1);" -- but"
6930   goto 6880
6940 if j > 0 then 6970
6950  if w$ = t$(c(1)) then t2 = c(2) : else t2 = c(1)
6960  goto 7070
6970 t1 = c(j)
6980 t2 = c(3-j)
6990 if w$ <> t$(t2) then 7060
7000  if g(t2) > 0 then 7040
7010   if g2 = 0 then 7090
7020   print "Note: '";t$(t2);"' used in premises taken to be ";g$(g2)
7030   goto 7090
7040  if g2 = 0 then 7090
7050  if g2 = g(t2) then 7090
7060 print "** Conclusion may not contain ";g$(g2);" '";w$;"';"
7070 print "** Conclusion must contain ";g$(g(t2));" '";t$(t2);"'."
7080 goto 7370
7090 if n1 = 0 or (d1 mod 2) = 1 then 7120
7100  print "** Negative conclusion required."
7110  goto 7370
7120 if n1 > 0 or d1 mod 2 = 0 then 7150
7130  print "** Affirmative conclusion required."
7140  goto 7370
7150 if j1 = 1 then 7250
7160 if d(t1) > 0 or d1 <= 1 or d1 >= 4 then 7200
7170  print "** Term '";t$(t1);"' not distributed in premises"
7180  print "   may not be distributed in conclusion."
7190  goto 7370
7200 if d(t2) > 0 then 7250
7210  if d1 mod 2 = 0 and d1 <> 6 then 7250
7220  print "** Term '";t$(t2);"' not distributed in premises"
7230  goto 7180
7250 print "-->  VALID!"
7260 if j1 = 0 then 7300
7270  if d1 > 0 then 7370
7280  t$(0) = w$
7290  goto 7350
7300 if d(t1) = 0 or d1 >= 2 then 7330
7310  v1 = t1
7320  goto 7350
7330 if d(t2) > 0 and d1 mod 2 = 0 and d1 <> 4 and d1 <> 6 then v1 = t2
7340 if v1 = 0 then 7370
7350 print "    but on Aristotelian interpretation only, i.e. on requirement"
7360 print "    that term '";t$(v1);"' denotes."
7370 return
7460 rem---list---
7530 i = 0
7540 i = l(i)
7550 if i = 0 then 7650
7570 print n(i);" ";
7580 if l1$ = "list" then 7630
7590  if r(i) < 6 and g(q(i)) = 2 then r(i) = r(i)+2
7600  if r(i) < 4 then print x$(r(i));"  ";
7610  print t$(p(i));y$(r(i));"  ";t$(q(i));z$(r(i))
7620  goto 7540
7630 print l$(i)
7640 goto 7540
7650 return
7660 rem---List valid inputs---
7670 cls : print "Valid commands are:"
7680 print "   <n>  [ <statement> ]   Insert, delete, or replace premise number  <n> "
7700 print left$(tb$,28);"Examples:   10  All men are mortal"
7704 print left$(tb$,40);"10"
7706 print "  DUMP";left$(tb$,15);"Prints symbol table, distribution count, etc."
7707 print "  HELP";left$(tb$,15);"Prints this list"
7708 print "  INFO";left$(tb$,15);"Gives information about syllogisms"
7710 print "  LIST";left$(tb$,15);"Lists premises"
7730 print "  LIST*";left$(tb$,14);"Same, but displays distribution analysis:"
7750 print left$(tb$,25);"distributed positions marked with '*', "
7760 print left$(tb$,25);"designators marked with '+'"
7830 print "  LINK";left$(tb$,15);"Lists premises in order of term-links (if possible)"
7850 print "  LINK*";left$(tb$,14);"Same, but in distribution-analysis format"
7855 print "  MSG";left$(tb$,16);"Turns on/off Printing of certain messages and warnings"
7860 print "  NEW";left$(tb$,16);"Erases current syllogism"
7865 print "  SAMPLE";left$(tb$,13);"Erases current syllogism and enters sample syllogism"
7867 print "  STOP";left$(tb$,15);"Stops entire program"
7870 print "  SUBSTITUTE";left$(tb$,9);"Allows uniform substitution of new terms in ";
7880 print "old premises"
7900 print "  SYNTAX";left$(tb$,13);"Explains statement syntax, with examples"
7943 print "  /";left$(tb$,18);"Asks program to draw conclusion"
7946 print "  /  <statement>";left$(tb$,5);"Tests  <statement>  as conclusion"
7948 print left$(tb$,25);"Note: this can be done even if there are no premises"
7950 return
7960 rem--"syntax"--
7970 cls : print "Valid statement forms:"
7980 print "  All    <general term #1>   is/are       <general term #2>"
7990 print "  Some   <general term #1>   is/are       <general term #2>"
8000 print "  Some   <general term #1>   is/are not   <general term #2>"
8010 print "  No     <general term #1>   is/are       <general term #2>"
8020 print
8030 print "   <designator>      is/are       <general term>"
8040 print "   <designator>      is/are not   <general term>"
8050 print "   <designator A>    is/are       <designator B>"
8060 print "   <designator A>    is/are not   <designator B>" : print
8080 print "Examples:"
8090 print "  All tall men are Greek gods             The teacher of Plato is wise"
8110 print "  Some cheese is tasty                    Socrates is not handsome"
8130 print "  Some cheese is not soft                 The teacher of Plato is Socrates"
8150 print "  No libertarians are cringing wimps      Socrates is not the";
8160 print " teacher of Thales"
8170 print
8180 print "Since e.g. 'Socrates is grunch' is ambiguous ('grunch' could be"
8190 print "either a designator or a general term), the program will try to"
8200 print "resolve the ambiguity from other uses of the term in the syllogism."
8210 print "The indefinite article 'sm' may be used with mass terms in predicates"
8220 print "(e.g. 'This puddle is sm ink') to ensure that the mass term is taken"
8230 print "as a general term rather than as a designator."
8240 return
8290 rem---Info---
8293 cls : print "   To use this program, enter a syllogism, one line at a time,"
8296 print "and  THEN  test conclusions or ask the program to draw a conclusion."
8298 print
8300 print "   A syllogism as (mis)defined here is a (possibly empty) set of"
8310 print "numbered premises, each of a form specified in the SYNTAX list."
8320 print "No term may occur more than twice.  Exactly two terms must occur"
8330 print "exactly once: these are the two 'end' terms, which will appear in"
8340 print "the conclusion.  Furthermore, each premise must have exactly one"
8350 print "term in common with its successor, for some ordering of the premises."
8360 print "Example:"
8370 print "   10 Socrates is a Greek"
8380 print "   20 All men are mortal"
8390 print "   30 All Greeks are men"
8395 print "   40 No gods are mortal" : print
8400 print "Note: using a '/' command to draw or test a conclusion does not"
8410 print "require you to stop.  You can continue, adding or deleting premises"
8415 print "and drawing and testing more conclusions." : print
8420 print "Reference:  H. Gensler, 'A Simplified Decision Procedure for Categor-"
8430 print "   ical Syllogisms,' Notre Dame J. of Formal Logic 14 (1973) 457-466."
8440 return
8890 rem---"Dump" values of variables---
8900 print "Highest symbol table loc. used:";l1;"  Negative premises:";n1
8910 if l1 = 0 then 8970
8920 print "Adr. art. term";left$(tb$,48-14);"type       occurs    dist. count"
8930 for i = 1 to l1
8931 rem Metal's lack of tabbing gets difficult here...
8932 itab = 7-len(str$(i))
8933 astringtab = 11-len(a$(b(i)))-7
8934 tstringtab = 49-len(t$(i))-11
8935 gtab = 60-len(str$(g(i)))-49
8936 otab = 71-len(str$(o(i)))-60
8940  print i;left$(tb$,itab);a$(b(i));left$(tb$,astringtab);t$(i);left$(tb$,tstringtab);g(i);left$(tb$,gtab);
8950  print o(i);left$(tb$,otab);d(i)
8960  next i
8970 return
8980 rem--sample--
9001 for z8 = 1 to 10 : read l1$ : print l1$
9002 gosub 2020 : gosub 2890 : gosub 4530 : gosub 3400
9003 next z8
9004 data "10 all mortals are fools"
9005 data "20 all athenians are men"
9006 data "30 all philosophers are geniuses"
9007 data "40 all people with good taste are philosophers"
9008 data "50 richter is a diamond broker"
9009 data "60 richter is the most hedonistic person in florida"
9010 data "70 all men are mortal"
9011 data "80 no genius is a fool"
9012 data "90 all diamond brokers are people with good taste"
9013 data "100 the most hedonistic person in florida is a decision-theorist"
9030 restore 9004
9040 if msg then print "Suggestion: try the LINK or LINK* command."
9050 return
9060 rem---Substitute terms---
9070 print "Enter address of old term; or 0 for help, -1 to exit, -2 for dump"
9080 input i1
9090 if i1 = -1 then 9470
9100 if i1 <> -2 then 9130
9110  gosub 8890
9120  goto 9070
9130 if i1 > 0 then 9340
9140 print "   This subroutine allows a term in a syllogism to be uniformly"
9150 print "replaced by another term.  This is useful e.g. for finding an"
9160 print "interpretation which actually makes the premises true, to produce as"
9170 print "an obvious example of invalidity an argument having exactly the same"
9180 print "logical form.  The substitution does not take place in the premises"
9190 print "as originally entered; it takes place in the terms as stored within"
9200 print "the program.  Thus, the LINK and LIST commands will display the"
9210 print "original premises; to see the changed ones, use the LIST* and LINK*"
9215 print "commands."
9220 print "   To find the 'addresses' of the terms, enter -2 to run the DUMP."
9230 print "   Warning: if you replace a term with another one already occurring"
9240 print "in the syllogism, the result will not make much sense.  However,"
9245 print "this routine does not convert entered term to lower-case or singular."
9250 goto 9455
9340 if i1 <= l1 then 9370
9350  print "Address ";i1;" too large.  Symbol table only of length ";l1
9360  goto 9455
9370 print "Enter new term to replace ";g$(g(i1));" '";t$(i1);"'"
9380  input w$
9440  t$(i1) = w$
9450  print "Replaced by '";w$;"'"
9455 print
9460  goto 9070
9470 print "Exit from substitution routine"
9480 return
9999 end
