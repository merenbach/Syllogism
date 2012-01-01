import os

# from the latest BASIC distribution:
#	Syllogism 1.0. November 8, 2002
#	I edited this program in 2002, for compatibility with freeware BASIC
#	interpreters for the Mac: Chipmunk BASIC 3.5.7 and Metal BASIC 1.7.3.
#	Peace. Ben Sharvy. luvnpeas99@yahoo.com
#
# This port is under active development by Andrew Merenbach

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
	def __init__(self):
		self.intro()

	def intro(self):
		self.cls()
		print "Syllogism Program Copyright (c) 1988 Richard Sharvy"
		print "Syllogism 1.0 (c) 2002 Richard Sharvy's estate"
		print "Ben Sharvy: luvnpeas99@yahoo.com or bsharvy@efn.org"
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

	def printCommands(self):
		# rem---List valid inputs--- : rem [am] 7660
		self.cls()
		print "Valid commands are:"
		print "   <n>  [ <statement> ]   Insert, delete, or replace premise number  <n> "
		print self.spaces(28) + "Examples:   10  All men are mortal"
		print self.spaces(40) + "10"
		print "  DUMP" + self.spaces(15) + "Prints symbol table, distribution count, etc."
		print "  HELP" + self.spaces(15) + "Prints this list"
		print "  INFO" + self.spaces(15) + "Gives information about syllogisms"
		print "  LIST" + self.spaces(15) + "Lists premises"
		print "  LIST*" + self.spaces(14) + "Same, but displays distribution analysis:"
		print self.spaces(25) + "distributed positions marked with '*', "
		print self.spaces(25) + "designators marked with '+'"
		print "  LINK" + self.spaces(15) + "Lists premises in order of term-links (if possible)"
		print "  LINK*" + self.spaces(14) + "Same, but in distribution-analysis format"
		print "  MSG" + self.spaces(16) + "Turns on/off Printing of certain messages and warnings"
		print "  NEW" + self.spaces(16) + "Erases current syllogism"
		print "  SAMPLE" + self.spaces(13) + "Erases current syllogism and enters sample syllogism"
		print "  STOP" + self.spaces(15) + "Stops entire program"
		print "  SUBSTITUTE" + self.spaces(9) + "Allows uniform substitution of new terms in old premises"
		print "  SYNTAX" + self.spaces(13) + "Explains statement syntax, with examples"
		print "  /" + self.spaces(18) + "Asks program to draw conclusion"
		print "  /  <statement>" + self.spaces(5) + "Tests  <statement>  as conclusion"
		print self.spaces(25) + "Note: this can be done even if there are no premises"
	
	def printSyntax(self):
		# rem--"syntax"-- : rem [am] 7960
		self.cls()
		print "Valid statement forms:"
		print "  All    <general term #1>   is/are       <general term #2>"
		print "  Some   <general term #1>   is/are       <general term #2>"
		print "  Some   <general term #1>   is/are not   <general term #2>"
		print "  No     <general term #1>   is/are       <general term #2>"
		print 
		print "   <designator>      is/are       <general term>"
		print "   <designator>      is/are not   <general term>"
		print "   <designator A>    is/are       <designator B>"
		print "   <designator A>    is/are not   <designator B>"
		print 
		print "Examples:"
		print "  All tall men are Greek gods             The teacher of Plato is wise"
		print "  Some cheese is tasty                    Socrates is not handsome"
		print "  Some cheese is not soft                 The teacher of Plato is Socrates"
		print "  No libertarians are cringing wimps      Socrates is not the teacher of Thales"
		print 
		print "Since e.g. 'Socrates is grunch' is ambiguous ('grunch' could be"
		print "either a designator or a general term), the program will try to"
		print "resolve the ambiguity from other uses of the term in the syllogism."
		print "The indefinite article 'sm' may be used with mass terms in predicates"
		print "(e.g. 'This puddle is sm ink') to ensure that the mass term is taken"
		print "as a general term rather than as a designator."
	
	def printInfo(self):
		# rem---Info--- : rem [am] 8290
		self.cls()
		print "   To use this program, enter a syllogism, one line at a time,"
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
		print "   40 No gods are mortal"
		print
		print "Note: using a '/' command to draw or test a conclusion does not"
		print "require you to stop.  You can continue, adding or deleting premises"
		print "and drawing and testing more conclusions."
		print
		print "Reference:  H. Gensler, 'A Simplified Decision Procedure for Categor-"
		print "   ical Syllogisms,' Notre Dame J. of Formal Logic 14 (1973) 457-466."

	def singularize(self, string):
		# divide by whitespace and remove blank items
		words = filter(None, string.split())
		#words_out = [plurals[word] for word in words if word in plurals.keys()]
		words_out = []
		for word in words:
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

#s = Syllogism()
