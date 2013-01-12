#!/usr/bin/python
# -*- coding: utf-8 -*-

import os



# from the latest BASIC distribution:
#	Syllogism 1.0. November 8, 2002
#	I edited this program in 2002, for compatibility with freeware BASIC
#	interpreters for the Mac: Chipmunk BASIC 3.5.7 and Metal BASIC 1.7.3.
#	Peace. Ben Sharvy. luvnpeas99@yahoo.com
#
# This port is under active development by Andrew Merenbach

#xyz = list()
#xyz.append("some", "  is", "<null>", "some", "  is not", "*")
#xyz.append("all", "*  is", "<null>", "no", "*  is", "*")
#xyz.append("<null>", "+  is", "<null>", "<null>", "+  is not", "*")
#xyz.append("<null>", "+  = ", "+", "<null>", "+   = / = ", "*")

#rem l1 => length_of_symbol_table
#rem n1 => negative_premise_count
#rem b(63) => term_article(63)
#rem d(63) => term_dist_count(63)
#rem e(2) => recent_article_types(2)
#rem g(63) => term_type(63)
#rem k(63) => link_order(63) : rem this may be right
#rem n(63) => line_numbers(63)
#rem o(63) => term_occurrences(63)
#rem t(7) => recent_symbol_types(7) : rem this should be right--see (former) line 2020
#rem a$(3) => article_strings[3)
#rem g$(2) => term_type_name$(2) : rem hopefully this is right
#rem l$(63) => line_strings$(63)
#rem s$(6) => recent_symbol_strings[6] : rem hopefully this is right
#rem t$(65) => term_strings$(65)
#rem w$( 2 ) => recent_term_strings

prompt = '>'

articles = ("a ", "an ", "sm ")
term_type_names = ("undetermined type", "general term", "designator")

sample_lines = (
	"10 all mortals are fools",
	"20 all athenians are men",
	"30 all philosophers are geniuses",
	"40 all people with good taste are philosophers",
	"50 richter is a diamond broker",
	"60 richter is the most hedonistic person in florida",
	"70 all men are mortal",
	"80 no genius is a fool",
	"90 all diamond brokers are people with good taste",
	"100 the most hedonistic person in florida is a decision-theorist",
)

x_array = ('some', 'some', 'all', 'no', '', '', '', '')
y_array = ("  is", "  is not", "*  is", "*  is", "+  is", "+  is not", "+  = ", "+   = / = ")
z_array = ("", "*", "", "*", "", "*", "+", "*")

# prepopulated plurals that might otherwise confuse the program
plurals = dict(
	[
		("socrates", "socrates"),
		("parmenides", "parmenides"),
		("epimenides", "epimenides"),
		("mice", "mouse"),
		("lice", "louse"),
		("geese", "goose"),
		("children", "child"),
		("oxen", "ox"),
		("people", "person"),
		("teeth", "tooth"),
		("wolves", "wolf"),
		("wives", "wife"),
		("selves", "self"),
		("lives", "life"),
		("leaves", "leaf"),
		("shelves", "shelf"),
		("elves", "elf"),
		("dwarves", "dwarf"),
		("knives", "knife"),
		("thieves", "thief"),
		("neckties", "necktie"),
		("hippies", "hippie"),
		("yippies", "yippie"),
		("yuppies", "yuppie"),
		("moonies", "moonie"),
		("druggies", "druggie"),
		("cookies", "cookie"),
		("commies", "commie"),
		("groupies", "groupie"),
		("tomatoes", "tomato"),
		("alcibiades", "alcibiades"),
		("thales", "thales"),
		("aries", "aries"),
		("athens", "athens"),
		("species", "species"),
		("feces", "feces"),
		("geniuses", "genius"),
		("sorites", "sorites"),
		("crises", "crisis"),
		("emphases", "emphasis"),
		("memoranda", "memorandum"),
		("theses", "thesis"),
		("automata", "automaton"),
		("formulae", "formula"),
		("stigmata", "stigma"),
		("lemmata", "lemma"),
		("vertices", "vertex"),
		("vortices", "vortex"),
		("indices", "index"),
		("codices", "codex"),
		("matrices", "matrix"),
		("gasses", "gas"),
		("gases", "gas"),
		("buses", "bus"),
		("aches", "ache"),
		("headaches", "headache"),
		("grits", "grits"),
		("molasses", "molasses"),
		("gas", "gas"),
		("christmas", "christmas"),
		("mathematics", "mathematics"),
		("semantics", "semantics"),
		("physics", "physics"),
		("metaphysics", "metaphysics"),
		("ethics", "ethics"),
		("linguistics", "linguistics"),
		("kiwis", "kiwi"),
		("israelis", "israeli"),
		("goyim", "goy"),
		("seraphim", "seraph"),
		("cherubim", "cherub"),
		("semen", "semen"),
		("amen", "amen"),
		("ZZZZZ", "ZZZZZ")
	]
)

class Syllogism:
	show_messages = True
	premise_list = []


	line_numbers_arranged = []		# l()
	line_strings = []		# l$()
	term_article = []		# b()
	term_strings = []		# t$()
	term_types = []			# g()
	conclusion_terms = []	# c()
	neg_premises = 0		# n1
	modern_valid = False	# v1
	symbol_count = 0		# l1
	lowest_line = 0			# l(0)

	#recent_term_strings = []	# w$()
	recent_term_1 = ''
	recent_term_2 = ''
	recent_term_type_1 = (-1)
	recent_term_type_2 = (-1)
	recent_symbol_types = []	# t()
	recent_symbol_strings = []	# s$()

	syllogism_form = (-1)		# d1

	a_array_0 = 0	# a(0)
	a_array = []	# a()

	q_array = []
	r_array = []

	# length_of_symbol_table = len(term_strings)
	#im a(63),c(63),term_dist_count(63),term_type(63),l(63),line_numbers(63),term_occurrences(63),p(63),q(63)
	#dim r(63),term_article(63),k(63),j(4),recent_symbol_types[7],recent_article_types(2),h(2)
	#dim article_strings[3),line_strings$(63),term_strings$(65)
	#dim g$(2),recent_symbol_strings[6],recent_term_strings[MY_TWO],x$(7),y$(7),z(7)

	def __init__(self):
		self.a_array = range(64)
		self.a_array_0 = 0

		self.line_numbers_arranged = [0] * 64		# l()
		self.line_strings = [''] * 64		# l$()
		self.term_article = [0] * 64		# b()
		self.term_strings = [''] * 64		# t$()
		self.term_types = [0] * 64			# g()
		self.recent_symbol_types = [0] * 64	# t()
		self.recent_symbol_strings = [''] * 64	# s$()


		q_array = [0] * 64
		r_array = [0] * 64
		
		self.main()
		#pass

	def main(self):
		self.intro()
		self.print_hint()
		self.new_syllogism()
		self.request_input()

	def intro(self):
		self.cls()
		print("Syllogism Program Copyright (c) 1988 Richard Sharvy")
		print("Syllogism 1.0 (c) 2002 Richard Sharvy's estate")
		print("Ben Sharvy: luvnpeas99@yahoo.com or bsharvy@efn.org")
		print

	def cls(self):
		# clear the screen
		os.system("clear")

	def spaces(self, space_count):
		# print a specified number of space
		return (space_count * ' ')
		#return ' '.ljust(space_count)
		#str_list = []
		#for n in range(space_count):
		#	str_list.append(' ')
		#s = ''.join(str_list)
		#s = ''.join([' ' for n in range(space_count)])

	def print_hint(self):
		if self.show_messages:
			print("Enter HELP for list of commands")

	def request_input(self):
		functions = {
			'new': self.new_syllogism,
			'sample': self.sample_syllogism,
			'help': self.print_commands,
			'syntax': self.print_syntax,
			'info': self.print_info,
			'dump': self.show_dump,
			'msg': self.toggle_messages,
			'substitute': self.substitute_terms,
			#'link': link(),
			#'link*': link(),
			'list': self.list_lines,
			'list*': self.list_lines,
		}

		line = ''
		while line != 'stop':
			print
			line = raw_input(prompt).lower()
			line = self.strip_string(line)
			if line == '':
				self.print_hint()
			else:
				if line in functions.keys():
					function = functions[line]
					if not line.endswith('*'):
						function()
					else:
						function(True)
				else:
					self.scan_line(line)
					
		if self.show_messages:
			print("(Some versions support typing CONT to continue)")
		print
	
	def toggle_messages(self):
		self.show_messages = not self.show_messages
		state = ''
		if self.show_messages:
			state = 'on'
		else:
			state = 'off'
		print('Messages turned {0}'.format(state))

	def strip_string(self, string):
		punctuation = ('.', '?', '!')
		string = string.rstrip()
		while string[-1:] in punctuation:
			print(self.spaces(len(string)) + "^   Punctuation mark ignored")
			#line = line.rstrip('.?!')
			string = string[:-1]
			string = string.rstrip()
		string = string.lstrip()
		return string

	def print_commands(self):
		# rem---List valid inputs--- : rem [am] 7660
		self.cls()
		print("Valid commands are:")
		print("   <n>  [ <statement> ]   Insert, delete, or replace premise number  <n> ")
		print(self.spaces(28) + "Examples:   10  All men are mortal")
		print(self.spaces(40) + "10")
		print("  DUMP" + self.spaces(15) + "Prints symbol table, distribution count, etc.")
		print("  HELP" + self.spaces(15) + "Prints this list")
		print("  INFO" + self.spaces(15) + "Gives information about syllogisms")
		print("  LIST" + self.spaces(15) + "Lists premises")
		print("  LIST*" + self.spaces(14) + "Same, but displays distribution analysis:")
		print(self.spaces(25) + "distributed positions marked with '*', ")
		print(self.spaces(25) + "designators marked with '+'")
		print("  LINK" + self.spaces(15) + "Lists premises in order of term-links (if possible)")
		print("  LINK*" + self.spaces(14) + "Same, but in distribution-analysis format")
		print("  MSG" + self.spaces(16) + "Turns on/off Printing of certain messages and warnings")
		print("  NEW" + self.spaces(16) + "Erases current syllogism")
		print("  SAMPLE" + self.spaces(13) + "Erases current syllogism and enters sample syllogism")
		print("  STOP" + self.spaces(15) + "Stops entire program")
		print("  SUBSTITUTE" + self.spaces(9) + "Allows uniform substitution of new terms in old premises")
		print("  SYNTAX" + self.spaces(13) + "Explains statement syntax, with examples")
		print("  /" + self.spaces(18) + "Asks program to draw conclusion")
		print("  /  <statement>" + self.spaces(5) + "Tests  <statement>  as conclusion")
		print(self.spaces(25) + "Note: this can be done even if there are no premises")
	
	def print_syntax(self):
		# rem--"syntax"-- : rem [am] 7960
		self.cls()
		print("Valid statement forms:")
		print("  All    <general term #1>   is/are       <general term #2>")
		print("  Some   <general term #1>   is/are       <general term #2>")
		print("  Some   <general term #1>   is/are not   <general term #2>")
		print("  No     <general term #1>   is/are       <general term #2>")
		print
		print("   <designator>      is/are       <general term>")
		print("   <designator>      is/are not   <general term>")
		print("   <designator A>    is/are       <designator B>")
		print("   <designator A>    is/are not   <designator B>")
		print
		print("Examples:")
		print("  All tall men are Greek gods             The teacher of Plato is wise")
		print("  Some cheese is tasty                    Socrates is not handsome")
		print("  Some cheese is not soft                 The teacher of Plato is Socrates")
		print("  No libertarians are cringing wimps      Socrates is not the teacher of Thales")
		print
		print("Since e.g. 'Socrates is grunch' is ambiguous ('grunch' could be")
		print("either a designator or a general term), the program will try to")
		print("resolve the ambiguity from other uses of the term in the syllogism.")
		print("The indefinite article 'sm' may be used with mass terms in predicates")
		print("(e.g. 'This puddle is sm ink') to ensure that the mass term is taken")
		print("as a general term rather than as a designator.")
	
	def print_info(self):
		# rem---Info--- : rem [am] 8290
		self.cls()
		print("   To use this program, enter a syllogism, one line at a time,")
		print("and  THEN  test conclusions or ask the program to draw a conclusion.")
		print
		print("   A syllogism as (mis)defined here is a (possibly empty) set of")
		print("numbered premises, each of a form specified in the SYNTAX list.")
		print("No term may occur more than twice.  Exactly two terms must occur")
		print("exactly once: these are the two 'end' terms, which will appear in")
		print("the conclusion.  Furthermore, each premise must have exactly one")
		print("term in common with its successor, for some ordering of the premises.")
		print("Example:")
		print("   10 Socrates is a Greek")
		print("   20 All men are mortal")
		print("   30 All Greeks are men")
		print("   40 No gods are mortal")
		print
		print("Note: using a '/' command to draw or test a conclusion does not")
		print("require you to stop.  You can continue, adding or deleting premises")
		print("and drawing and testing more conclusions.")
		print
		print("Reference:  H. Gensler, 'A Simplified Decision Procedure for Categor-")
		print("   ical Syllogisms,' Notre Dame J. of Formal Logic 14 (1973) 457-466.")

	def singularize(self, string):
		# divide by whitespace and remove blank items
		words = filter(None, string.split())
		#words_out = [plurals[word] for word in words if word in plurals.keys()]
		words_out = []
		for word in words:
			word = word.lower()
			if word in plurals.keys():
				words_out.append(plurals[word])
			else:
				if word.endswith('men'):
					word = word[:-2] + 'an'
				elif word.endswith('s'):
					if not word.endswith('ss') and not word.endswith('us') and not word.endswith('is') and not word.endswith("'s"):
						word = word[:-1]
						if word.endswith('xe'):
							word = word[:-1]
						elif word.endswith('ie'):
							word = word[:-2] + 'y'
						elif word.endswith('sse') or word.endswith('she') or word.endswith('che'):
							y = word[:-1]
				words_out.append(word);
		return ' '.join(words_out)

	def show_error_no_premises(self):
		print("No premises")

	def show_parse_error_missing_copula(self):
		print("** Missing copula is/are")
		self.show_parse_error_help()
	def show_parse_error_missing_subject_term(self):
		print("** Subject term bad or missing")
		self.show_parse_error_help()
	def show_parse_error_missing_predicate(self):
		print("** Predicate term bad or missing")
		self.show_parse_error_help()
	def show_parse_error_help(self):
		if self.show_messages:
			print("Enter SYNTAX for help with statements")

	def sample_syllogism(self):
		# 8980 rem--sample--
		self.new_syllogism()
		for line in sample_lines:
			print(line)
			self.split_line(line)
			self.parse_line()
			self.enter_line(line)
			self.insert_terms()
			
			# new
			#p = Premise(line)
			#self.premise_list.append(p)

		if self.show_messages:
			print("Suggestion: try the LINK or LINK* command.")
			
	def show_dump(self):
		# 8890 rem---"Dump" values of variables---
		print("Highest symbol table loc. used: {0}  Negative premises: {1}".format(self.symbol_count, self.neg_premises))
		if self.symbol_count > 0:
			print("Adr. art. term {0} type       occurs    dist. count".format(self.spaces(48-14)))
			for i in range(self.symbol_count):
				# rem Metal's lack of tabbing gets difficult here...
				itab = 7-len(str(i))
				astringtab = 11-len(self.article_strings[self.term_article[i]])-7
				tstringtab = 49-len(self.term_strings[i])-11
				gtab = 60-len(str(self.term_type[i]))-49
				otab = 71-len(str(self.term_occurrences[i]))-60
				print(i + self.spaces(itab) + article_strings[term_article[i]] + self.spaces(astringtab) + term_strings[i] + self.spaces(tstringtab) + term_type[i] + self.spaces(gtab))
				print(self.term_occurrences[i] + self.spaces(otab) + self.term_dist_count[i])

	# should work, but not fully tested
	def substitute_terms(self):
		#9060 rem---Substitute terms---
		address = 0
		while address != (-1):
			skip = False
			print('Enter address of old term; or 0 for help, -1 to exit, -2 for dump')
			address = raw_input(prompt)
			try:
				address = int(address)
			except:
				skip = True
			if address != -1 and skip == False:
				if address == -2:
					self.show_dump()
				else:
					if address == 0:
						print("   This subroutine allows a term in a syllogism to be uniformly")
						print("replaced by another term.  This is useful e.g. for finding an")
						print("interpretation which actually makes the premises true, to produce as")
						print("an obvious example of invalidity an argument having exactly the same")
						print("logical form.  The substitution does not take place in the premises")
						print("as originally entered; it takes place in the terms as stored within")
						print("the program.  Thus, the LINK and LIST commands will display the")
						print("original premises; to see the changed ones, use the LIST* and LINK*")
						print("commands.")
						print("   To find the 'addresses' of the terms, enter -2 to run the DUMP.")
						print("   Warning: if you replace a term with another one already occurring")
						print("in the syllogism, the result will not make much sense.  However,")
						print("this routine does not convert entered term to lower-case or singular.")
					else:
						if address >= self.symbol_count:
							print("Address {0} too large.  Symbol table only of length ".format(address, self.symbol_count))
						else:
							print("Enter new term to replace {0} '".format(term_type_names[self.term_type[address]], self.term_strings[address]))
							new_term = raw_input(prompt)
							self.term_strings[address] = new_term
							print('Replaced by "{0}"'.format(new_term))
					print
		print("Exit from substitution routine")
	
	# experimental; need sanity-checks
	# should work, though!
	def compute_conclusion(self):
		#rem---Compute conclusion--- : rem 6200
		c1 = self.conclusion_terms[0]
		c2 = self.conclusion_terms[1]
		term_article_c1 = self.term_article[c1]
		term_article_c2 = self.term_article[c2]
		term_strings_c1 = self.term_strings[c1]
		term_strings_c2 = self.term_strings[c2]
		article_strings_c1 = self.article_strings[term_article_c1]
		article_strings_c2 = self.article_strings[term_article_c2]
		term_dist_count_c1 = self.term_dist_count[c1]
		term_dist_count_c2 = self.term_dist_count[c2]

		z = 'A is A'
		if self.lowest_line > 0:
			if self.neg_premises > 0:
				# negative conclusion
				if term_dist_count_c2 == 0:
					z = "Some {0} is not {1}{2}".format(term_strings_c2, article_strings_c1, term_strings_c1)
				elif term_dist_count_c1 == 0:
					z = "Some {0} is not {1}{2}".format(term_strings_c1, article_strings_c2, term_strings_c2)
				elif self.term_type[c1] >= 2:
					z = "{0} is not {1}{2}".format(term_strings_c1, article_strings_c2, term_strings_c2)
				elif self.term_type[c2] >= 2:
					z = "{0} is not {1}{2}".format(term_strings_c2, article_strings_c1, term_strings_c1)
				elif term_article_c1 == 0 and term_article_c2 > 0:
					z = "No {0} is {1}{2}".format(term_strings_c2, article_strings_c1, term_strings_c1)
				else:
					z = "No {0} is {1}{2}".format(term_strings_c1, article_strings_c2, term_strings_c2)
			else:
				# affirmative conclusion
				if term_dist_count_c1 > 0:
					if self.term_type[c1] != 2:
						z = "All {0} is {1}".format(term_strings_c1, term_strings_c2)
					else:
						z = "{0} is {1}{2}".format(term_strings_c1, article_strings_c2, term_strings_c2)
				elif term_dist_count_c2 > 0:
					if self.term_type[c2] != 2:
						z = "All {0} is {1}".format(term_strings_c2, term_strings_c1)
					else:
						z = "{0} is {1}{2}".format(term_strings_c2, article_strings_c1, term_strings_c1)
				else:
					if term_article_c1 == 0 and term_article_c2 > 0:
						z = "Some {0} is {1}{2}".format(term_strings_c2, article_strings_c1, term_strings_c1)
					else:
						z = "Some {0} is {1}{2}".format(term_strings_c1, article_strings_c2, term_strings_c2)
		# PRINT  conclusion
		print('  / ' + z)
		if modern_valid:
			print('  * Aristotle-valid only, i.e. on requirement that term "{0}" denotes.'.format(self.term_strings[v1]))

	def enter_line(self, line=''):
		# rem---Enter line into list--- : rem 4530
		if (self.recent_symbol_strings[0].isdigit()):
			n = int(self.recent_symbol_strings[0])
			s = len(self.recent_symbol_strings[1]) + 1
			l = len(line)
			l_string = line[s + 1 : l - s]
			i = self.lowest_line
			while True:
				j1 = self.line_numbers_arranged[i]
				if j1 == 0:
					self.sub_enter_line(a1, j1)
					break
				elif n == self.line_numbers_arranged[j1]:
					self.decrement_table_entries()
					self.line_strings[j1] = l_string
					a1 = j1
					break
				elif n < self.line_numbers_arranged[j1]:
					self.sub_enter_line(a1, j1)
					break
				else:
					i = j1

	def delete_line(self):
		# rem---Delete a line--- : rem [am] 4760
		if (self.recent_symbol_strings[0].isdigit()):
			n = int(self.recent_symbol_strings[0])
			i = self.lowest_line
			while True:
				j1 = self.line_numbers_arranged[i]
				if j1 == 0:
					print("Line " + n + " not found")
					break
				elif n == self.line_numbers_arranged[j1]:
					self.a_array_0 -= 1
					self.a_array[a_array_0] = j1
					self.line_numbers_arranged[i] = self.line_numbers_arranged[j1]
					self.decrement_table_entries()
					break
				else:
					i = self.line_numbers_arranged[i]

	def decrement_table_entries(self):
		# rem---Decrement table entries--- : rem [am] 4890
		# local
		j_array = [0] * 4

		j_array[0] = p(j1)
		j_array[1] = q(j1)
		if self.r_array[j1] % 2 != 0:
			self.neg_premises -= 1
			j_array[3] = 1
		else:
			if self.term_types[q_array[j1]] == 2:
				j_array[3] = 1
			else:
				j_array[3] = 0
		if self.r_array[j1] >= 2:
			j_array[2] = 1
		else:
			j_array[2] = 0
		for k in range(2):
			self.term_occurrences[j_array[k]] -= 1
			if self.term_occurrences[j_array[k]] == 0:
				self.term_strings[j_array[k]] = ""
				self.term_article[j_array[k]] = 0
				self.term_type[j_array[k]] = 0
			endif
			self.term_dist_count[j_array[k]] -= j_array[k+2]

	def sub_enter_line(self, a1, j1):
		a1 = self.a_array[self.a_array_0]
		self.line_strings[a1] = l_string
		self.line_numbers[a1] = n
		self.line_numbers_arranged[i] = a1
		self.line_numbers_arranged[a1] = j1
		self.a_array_0 += 1

	def list_lines(self, analyze=False):
		# rem---list--- : rem [am] 7460
		out_list = []
		i = self.lowest_line
		while i > 0:
			out = ''
			out += self.line_numbers_arranged[i] + " "
			if analyze:
				# distribution analysis format
				if r_array[i] < 6 and self.term_type[q_array[i]] == 2:
					r_array[i] += 2
				if r_array[i] < 4:
					out += x_array[r_array[i]] + "  "
				out += self.term_strings[p_array[i]] + y_array[r_array[i]] + "  " + self.term_strings[q_array[i]] + z_array[r_array[i]]
			else:
				out += line_strings[i]
			out_list.append(out)
			i += 1
		if len(out_list) > 0:
			print("\n".join(out_list))

	#def list_premises(self, analyze=False):
	#	for p in self.premise_list:
	#		print p

	# implemented but a_array is not entirely clear
	def new_syllogism(self):
		if self.lowest_line > 0:
			for i in range(self.symbol_count):
				self.term_dist_count[i] = 0
				self.term_strings[i] = ''
				self.term_article[i] = 0
				self.term_occurrences[i] = 0
				self.term_type[i] = 0
			self.symbol_count = 0
			self.neg_premises = 0
			
			#self.term_dist_count = []
			#self.term_strings = []
			#self.term_article = []
			#self.term_occurrences = []
			#self.term_type = []
			j = self.lowest_line
			while j > 0:
				self.a_array_0 -= 1
				self.a_array[self.a_array_0] = j
				j = self.line_numbers_arranged[j]
			#self.line_numbers_arranged = []
			self.lowest_line = 0

	def insert_terms(self):
		# rem---Add self.recent_term_1, self.recent_term_2 to table term_strings$()--- : rem [am] 3400
		pass

	def show_error_invalid_cmd(self, i):
		print(self.spaces(spaces) + "^   Invalid numeral or command")

	def split_line(self, line=''):
		pass
		# rem--scan line L1$ into array S$() : rem 2020
		# 2020 rem---L1$ into array S$()---
		#	# rem T(): 1:line num., 2:"/", 3:quantifier, 4:no/not, 5:is/are, 6:term
		#	# rem                     10 SOME  FRIED COCONUTS   ARE  NOT  TASTY
		#	# rem                      1   3        6            5    4     6
		#	self.term_strings = filter(None, line.split())
		#	self.term_types = list()
		recent_symbol_strings = [''] * 6
		recent_symbol_types = [0] * 6
		for word in self.term_strings:
			if len(self.term_strings) == 1:
				if word == '/':
					self.recent_symbol_types[0] = 2
					pass
				else:
					if len(word) > 4:
						#show_error_invalid_cmd(i + n)
						break
					else:
						if (has_invalid_chars):
							#show_error_invalid_cmd(i + n)
							pass
						else:
							#self.term_types[1] = 1
							pass
						pass
					n = len(recent_symbol_strings)
					#if n > 4 then
					#	procERROR_INVALID_CMD(i + n)
					#	goto 2885
					#else
					#	for n in range(len(recent_symbol_strings)):
					#		t$ = mid$(s$,n,1)
					#		if asc(t$) > 57 or asc(t$) < 48 then
					#			procERROR_INVALID_CMD(i + n)
					#			goto 2885
					#		endif
					#	next n
					#	t(1) = 1
					#endif
					#if len(self.term_strings) == 6:
					#	
					#		pass
	
	
		#	#for j = 1 to 6
		#	#	s$(j) = ""
		#	#	t(j) = 0
		#	#next j
		#	p1 = 0
		#	recent_article_types(2) = 0
		#	j = 1
		#	i = 1
		#	l = len(l1$)
		#	do
		#		do
		#			if i > l then 2885
		#			s$ = mid$(l1$,i,1)
		#			if s$ != " " then exit do
		#			i = i+1
		#		loop
		#		for k = 1 to (l - i)
		#			s$ = mid$(l1$,i+k,1)
		#			if s$ = " " then exit for
		#		next k
		#		s$ = mid$(l1$,i,k) : rem S$ is set to next word
		#		if j <= 1 then
		#			if s$ = "/" then
		#				t(1) = 2
		#			else
		#				n = len(s$)
		#				if n > 4 then
		#					procERROR_INVALID_CMD(i + n)
		#					goto 2885
		#				else
		#					for n = 1 to len(s$)
		#						t$ = mid$(s$,n,1)
		#						if asc(t$) > 57 or asc(t$) < 48 then
		#							procERROR_INVALID_CMD(i + n)
		#							goto 2885
		#						endif
		#					next n
		#					t(1) = 1
		#				endif
		#			endif
		#			goto 2840
		#		endif
		#rem Scan : rem [am] 2520
		#		if s$ = "somebody" or s$ = "something" or s$ = "nobody" or s$ = "nothing" then
		#			procERR_RESERVED_WORD(s$, i + k - 1)
		#			t(1) = 0
		#			goto 2885
		#		elseif s$ = "someone" or s$ = "everyone" or s$ = "everybody" or s$ = "everything" then
		#			procERR_RESERVED_WORD(s$, i + k - 1)
		#			t(1) = 0
		#			goto 2885
		#		elseif s$ = "all" or s$ = "some" then
		#			if t(j) = 6 then
		#				procERR_RESERVED_WORD(s$, i + k - 1)
		#				t(1) = 0
		#				goto 2885
		#			else
		#				t(j) = 3
		#				goto 2840
		#			endif
		#		elseif s$ = "no" or s$ = "not" then
		#			if t(j) = 6 then
		#				procERR_RESERVED_WORD(s$, i + k - 1)
		#				t(1) = 0
		#				goto 2885
		#			else
		#				t(j) = 4
		#				goto 2840
		#			endif
		#		elseif s$ = "is" or s$ = "are" then
		#			if t(j) = 6 then
		#				if not (t(j-1) = 5 or t(j-2) = 5) then
		#					j = j+1
		#					t(j) = 5
		#					goto 2840
		#				endif
		#			endif
		#			procERR_RESERVED_WORD(s$, i + k - 1)
		#			t(1) = 0
		#			goto 2885
		#		elseif t(j) <> 6 then
		#			if t(j-1) = 5 or t(j-2) = 5 then
		#				if s$ = "a" or s$ = "an" or s$ = "sm" then
		#					if i <> l then
		#						if s$ = "a" then
		#							recent_article_types(2) = 1
		#						elseif s$ = "an" then
		#							recent_article_types(2) = 2
		#						else
		#							recent_article_types(2) = 3
		#						endif
		#						p1 = 1
		#					else
		#						gosub 2790
		#					endif
		#				else
		#					if s$ = "the" then p1 = 2
		#					gosub 2790
		#				endif
		#			else
		#				gosub 2790
		#			endif
		#		else
		#			s$(j) = s$(j)+" "+s$
		#		endif
		#		goto 2860
		#2840	s$(j) = s$
		#		j = j+1
		#2860	i = k+i
		#	loop until j > 6
		#2885 return


	# array indices need adjustment
	# rem---Parse line in S$()--- : rem [am] 2890
	def parse_line(self):
		self.syllogism_form = -1
		if self.term_strings[1] == "all":
			if self.term_strings[2] != 6:
				self.show_parse_error_missing_subject_term()
			elif self.term_types[3] != 5:
				self.show_parse_error_missing_copula()
			elif self.term_types[4] != 6:
				self.show_parse_error_missing_predicate()
			else:
				self.recent_term_1 = self.term_strings[2]
				self.recent_term_2 = self.term_strings[4]
				# rem all A is B
				self.syllogism_form = 2
		elif self.term_strings[1] == "some":
			if self.term_types[2] != 6:
				self.show_parse_error_missing_subject_term()
			elif self.term_types[3] != 5:
				self.show_parse_error_missing_copula()
			elif self.term_strings[4] != "not":
				if self.term_types[4] != 6:
					self.show_parse_error_missing_predicate()
				else:
					self.recent_term_1 = self.term_strings[2]
					self.recent_term_2 = self.term_strings[4]
					# rem Some A is B
					self.syllogism_form = 0
			else:
				if self.term_types[5] != 6:
					self.show_parse_error_missing_predicate()
				else:
					self.recent_term_1 = self.term_strings[2]
					self.recent_term_2 = self.term_strings[5]
					# rem some A is not B
					self.syllogism_form = 1
		elif self.term_strings[1] == "no":
			if self.term_types[2] != 6:
				self.show_parse_error_missing_subject_term()
			elif self.term_types[3] != 5:
				self.show_parse_error_missing_copula()
			elif self.term_types[4] != 6:
				self.show_parse_error_missing_predicate()
				self.show_parse_error_help()
			else:
				# rem no A is B
				self.recent_term_1 = self.term_strings[2]
				self.recent_term_2 = self.term_strings[4]
				self.syllogism_form = 3
		elif self.term_types[1] != 6:
			self.show_parse_error_missing_subject_term()
		elif self.term_types[2] == 5:
			self.recent_term_1 = self.term_strings[1]
			if self.term_strings[3] != "not":
				if self.term_types[3] != 6:
					self.show_parse_error_missing_predicate()
				# rem a is T
				self.syllogism_form = 4
				self.recent_term_2 = self.term_strings[3]
			else:
				if self.term_types[4] != 6:
					self.show_parse_error_missing_predicate()
				else:
					# rem a is not T
					self.syllogism_form = 5
					self.recent_term_2 = self.term_strings[4]
		else:
			self.show_parse_error_missing_copula()

	def test_offered_conclusion(self):
		# 6630 rem---test offered conclusion---
		# rem--conc. poss, line in s$()
		pass
	
	def see_if_conclusion_possible(self):
		# 5880 rem---See if conclusion possible---
		pass
	
	def see_if_syllogism(self):
		# 5070 rem---See if syllogism---
		return (-1)

	# still being reworked
	def scan_line(self, line=''):
		# 1570 rem--scan line L1$ into array S$()
		self.split_line(line)
		if self.recent_term_type_1 == 1:
			if self.recent_term_type_2 > 0:
				# rem parse the line in S$()
				self.parse_line()
				if syllogism_form >= 0:
					# enter line into list
					self.enter_line()
					# add terms to symbol tablosue
					self.insert_terms()
			else:
				if self.lowest_line > 0:
					# delete line
					self.delete_line()
				else:
					self.show_error_no_premises()
		else:
			if self.recent_term_type_1 == 0:
				self.print_hint()
			else:
				# draw/test conclusion
				# is it a syl?
				j1 = self.see_if_syllogism()
				if j1 <= 1:
					if j1 == 0:
						# poss. conclusion?
						self.see_if_conclusion_possible()
					if j1 <= 1:
						if self.recent_term_type_2:
							self.test_offered_conclusion()
						else:
							# test/draw conclusion
							self.compute_conclusion()


#class Premise:
#	line_num = ''
#	line_txt = ''
#	term_1 = ''
#	term_2 = ''
#	term_1_type = (-1)
#	term_2_type = (-1)
#
#	symbol_strings = []
#	symbol_types = []
#
#	def __init__(self, txt=''):
#		line_num = (-1)
#		line_txt = txt
#		term_1 = ''
#		term_2 = ''
#	
#	def parse_line(self):
#		pass
#
#	def split_line(self):
#		pass
#
#
#	def symbol_string_with_index(self, idx):
#		r = ''
#		if idx < len(symbol_strings):
#			r = symbol_strings[idx]
#		return r
#	
#	def symbol_type_with_index(self, idx):
#		r = ''
#		if idx < len(symbol_types):
#			r = symbol_types[idx]
#		return r

s = Syllogism()
#test_line1 = '10 all men are mortal'
#p = Premise(test_line1)

#s.new_syllogism()


#def contains_any(str, set):
#	flag = False
#    for c in set:
#        if c in str:
#	        flag = True
#	        break
#    return flag
#
#def contains_all(str, set):
#	flag = True
#    for c in set:
#        if c not in str:
#	        flag = False
#	        break
#    return flag##