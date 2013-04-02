#!/usr/bin/python
# -*- coding: utf-8 -*-

import os

# from the latest BASIC distribution:
#   Syllogism 1.0. November 8, 2002
#   I edited this program in 2002, for compatibility with freeware BASIC
#   interpreters for the Mac: Chipmunk BASIC 3.5.7 and Metal BASIC 1.7.3.
#   Peace. Ben Sharvy. luvnpeas99@yahoo.com
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

#### New stuff: error messages!

MSG_NEW_SYLLOGISM = 'Begin new syllogism'
MSG_USAGE_HINT = 'Enter HELP for list of commands'
MSG_STOPPED = '(Some versions support typing CONT to continue)'
MSG_NO_PREMISES = 'No premises'
MSG_LINK_SUGGEST = 'Suggestion: try the LINK or LINK* command.'

COPYRIGHT_LINES = (
    "Syllogism Program Copyright (c) 1988 Richard Sharvy",
    "Syllogism 1.0 (c) 2002 Richard Sharvy's estate",
    "Ben Sharvy: luvnpeas99@yahoo.com or bsharvy@efn.org",
)

INFO_BLURB = """   To use this program, enter a syllogism, one line at a time,
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
   ical Syllogisms,' Notre Dame J. of Formal Logic 14 (1973) 457-466."""

SYNTAX_BLURB = """Valid statement forms:
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
as a general term rather than as a designator."""

MSG_PROMPT = '>'

SAMPLE_LINES = (
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

####

articles = ("a ", "an ", "sm ")
term_type_names = ("undetermined type", "general term", "designator")


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


class Premise(object):
    def __init__(self, line):
        self.raw = line.strip()
        tokens = line.strip().split()
        self.line_number = int(tokens[0])
        self.statement = u' '.join(tokens[1:])
    
    def empty(self):
        """ Check whether this premise actually contains a statement.

        Returns
        -------
        boolean : `True` if a statement exists for this premise, `False` otherwise.
        """
        return not self.statement or len(self.statement) == 0
    
    def __repr__(self):
        return self.raw

class Rubric(object):
    """ Hold a list of premises. """
    def __init__(self):
        self.premises = []

    def __len__(self):
        """ Return the number of premises """
        return len(self.premises)

    def reset(self):
        """ Remove all premises. """
        self.premises = []

    def enter_premise(self, premise):
        """ Try to add a premise to our lookup table.
        
        Parameters
        ----------
        premise : Premise
                  A premise to add to our lookup table.

        Returns
        -------
        boolean : `True` if the premise was added successfully, `False` otherwise.
                  If this method returns `False`, odds are that an error message was printed.
        """
        if premise.empty():
            return self.remove_line(premise.line_number, silent=False)
        else:
            # Remove any lines with the same line number
            self.remove_line(premise.line_number, silent=True)
            self.premises.append(premise)
            # If a statement was included, replace the existing line
            # Otherwise, simply don't put the statement into place

            # Sort the new premises by line number
            # We only need to do this when adding lines                
            self.premises.sort(key=lambda p: p.line_number)
            return True

    def remove_line(self, line_number, silent=False):
        """ Remove a premise (identified by line number) from the lookup table.
        
        Parameters
        ----------
        line_number : integer
                      A premise line number for whose existence to check.
        silent      : boolean
                      `True` to complain if line does not exist or if no premises are entered,
                      `False` otherwise.

        Returns
        -------
        boolean : `True` if the line was removed, `False` otherwise.
        """
        if self.line_exists(line_number):
            self.premises = [p for p in self.premises if p.line_number != line_number]
            return True
        elif len(self.premises) == 0:
            # No premises have been entered
            if not silent:
                print(MSG_NO_PREMISES)
            return False
        else:
            # The premise to remove did not exist
            if not silent:
                print("Line {0} not found".format(line_number))
            return False

    def line_exists(self, line_number):
        """ Check if a premise with a given line number exists in the lookup table.
        
        Parameters
        ----------
        line_number : integer
                      A premise line number for whose existence to check.

        Returns
        -------
        boolean : `True` if a premise with the given line number exists already, `False` otherwise.
        """
        for p in self.premises:
            if p.line_number == line_number:
                return True
        return False

    def p(self, analyze=False):
        """ Printable format """
        lines = []
        last_premise = self.premises[-1:]
        if len(last_premise) > 0:
            max_padding_chars = len(str(last_premise[0].line_number))
            # Format lines with nice spacing
            premise_groups = ((p.line_number, p.statement) for p in self.premises)
            for p in premise_groups:
                if not analyze:
                    lines.append(u' {0} {1}'.format(str(p[0]).rjust(max_padding_chars), p[1]))
                else:
                    lines.append(u" [TODO] Distribution analysis listing")
        return u'\n'.join(lines)

    def __repr__(self):
        return self.p()

class Syllogism(object):

    def __init__(self):
        self.show_messages = True
        self.rubric = Rubric()

    def run(self):
        # Clear the screen
        self.cls()

        # Print copyright messages
        for c in COPYRIGHT_LINES:
            print(c)
        print
        
        # Print a usage hint
        self.print_hint()
        
        # Create a new document: Currently unnecessary
        #self.new_syllogism()
        
        # Start the request loop
        self.request_input()

    def cls(self):
        # clear the screen
        os.system("clear")

    def spaces(self, space_count):
        # print a specified number of space
        return (space_count * ' ')
        #return ' '.ljust(space_count)
        #str_list = []
        #for n in range(space_count):
        #   str_list.append(' ')
        #s = ''.join(str_list)
        #s = ''.join([' ' for n in range(space_count)])

    def print_hint(self):
        """ Print usage hint """
        self.print_message(MSG_USAGE_HINT)

    def request_input(self):
        """ Main loop """
        functions = {
            'new': self.new_syllogism,
            'sample': self.sample_syllogism,
            'help': self.print_commands,
            'syntax': self.print_syntax,
            'info': self.print_info,
            # 'dump': self.show_dump,
            'msg': self.toggle_messages,
            # 'substitute': self.substitute_terms,
            #'link': link(),
            #'link*': link(),
            'list': self.list_lines,
            'list*': self.list_lines,
        }

        line = ''
        while True:
            #print
            line = raw_input(MSG_PROMPT).lower()
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
                elif line == 'stop':
                    break
                else:
                    self.enter_line(line)
        self.print_message(MSG_STOPPED)

    def print_message(self, msg):
        """ Print a message if `show_messages` is enabled """
        if self.show_messages:
            print(msg)

    def print_indented_msg(self, msg, offset):
        print(u'{0}^   {1}'.format(self.spaces(offset + len(MSG_PROMPT) - 1), msg))

    def toggle_messages(self):
        """ Toggle the state of certain messages """
        self.show_messages = not self.show_messages
        if self.show_messages:
            state = 'on'
        else:
            state = 'off'
        print('Messages turned {0}'.format(state))

    def strip_string(self, s):
        """ Remove punctuation from the end of a string

        Parameters
        ----------
        s : string
            A string to strip.

        Returns
        -------
        string : a processed version of the original string.
        """
        punctuation = ('.', '?', '!')
        s = s.rstrip()
        while s.endswith(punctuation):
            self.print_indented_msg("Punctuation mark ignored", len(s))
            s = s[:-1].rstrip()
            #line = line.rstrip('.?!')
        s = s.lstrip()
        return s

    def print_commands(self):
        """ List valid inputs """
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
        print(SYNTAX_BLURB)
    
    def print_info(self):
        # rem---Info--- : rem [am] 8290
        self.cls()
        print(INFO_BLURB)

    def singularize(self, string):
        # divide by one or more whitespace characters
        words = string.split()
        #words_out = [plurals[word] for word in words if word in plurals.keys()]
        words_out = []
        for word in words:
            word = word.lower()
            if word in plurals.keys():
                words_out.append(plurals[word])
            else:
                # Try to make some educated guesses
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
        return u' '.join(words_out)

    ## These all work but are unused.  They are therefore commented out.
    #def show_parse_error_missing_copula(self):
    #    print("** Missing copula is/are")
    #    self.show_parse_error_help()
    #def show_parse_error_missing_subject_term(self):
    #    print("** Subject term bad or missing")
    #    self.show_parse_error_help()
    #def show_parse_error_missing_predicate(self):
    #    print("** Predicate term bad or missing")
    #    self.show_parse_error_help()
    #def show_parse_error_help(self):
    #    if self.show_messages:
    #        print("Enter SYNTAX for help with statements")

    def sample_syllogism(self):
        """ Enter a sample syllogism. """
        # 8980 rem--sample--
        self.new_syllogism(False)
        for line in SAMPLE_LINES:
            print(line)
            self.enter_line(line)
        self.print_message(MSG_LINK_SUGGEST)

    def list_lines(self, analyze=False):
        """ Cover method to list out lines, optionally in a distribution-analysis format. """
        # rem---list--- : rem [am] 7460
        if len(self.rubric) > 0:
            print(self.rubric.p(analyze))
        else:
            print(MSG_NO_PREMISES)

    def enter_line(self, line):
        """ Try to parse a string into a premise and add it to our rubric.
        
        Parameters
        ----------
        line : string
               A string to parse into a premise.
        """
        try:
            premise = Premise(line)
            self.rubric.enter_premise(premise)
        except ValueError:
            # Invalid input
            print("*** Invalid entry [am].")

    def new_syllogism(self, show_message=True):
        """ Remove all premises from the rubric. """
        if show_message:
            print(MSG_NEW_SYLLOGISM)
        self.rubric.reset()

    # This works but is unused
    #def show_error_invalid_cmd(self, i):
    #    self.print_indented_msg("Invalid numeral or command", i)


#class Premise:
#   line_num = ''
#   line_txt = ''
#   term_1 = ''
#   term_2 = ''
#   term_1_type = (-1)
#   term_2_type = (-1)
#
#   symbol_strings = []
#   symbol_types = []
#
#   def __init__(self, txt=''):
#       line_num = (-1)
#       line_txt = txt
#       term_1 = ''
#       term_2 = ''
#   
#   def parse_line(self):
#       pass
#
#   def split_line(self):
#       pass
#
#
#   def symbol_string_with_index(self, idx):
#       r = ''
#       if idx < len(symbol_strings):
#           r = symbol_strings[idx]
#       return r
#   
#   def symbol_type_with_index(self, idx):
#       r = ''
#       if idx < len(symbol_types):
#           r = symbol_types[idx]
#       return r

Syllogism().run()

s = Rubric()
#s.enter_line("10 all men are mortal")
#s.enter_line("30 all men are mortal")
#s.enter_line("30 a no men are mortal")
#s.enter_line("401 all men are mortal")
#s.enter_line("41 all men are mortal")
#s.enter_line("4 all men are mortal")
#s.enter_line("4012 all men are mortal")
#s.enter_line("50 all men are mortal")

print(s)

#test_line1 = '10 all men are mortal'
#p = Premise(test_line1)

#s.new_syllogism()


#def contains_any(str, set):
#   flag = False
#    for c in set:
#        if c in str:
#           flag = True
#           break
#    return flag
#
#def contains_all(str, set):
#   flag = True
#    for c in set:
#        if c not in str:
#           flag = False
#           break
#    return flag##