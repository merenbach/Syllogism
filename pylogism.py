#! /usr/bin/env python3
    # [TODO] should be #! /usr/bin/env python3
# -*- coding: utf-8 -*-

import os
from collections import UserList, UserString

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

def spaces(space_count):
    # print a specified number of space
    return (space_count * ' ')
    #return ' '.ljust(space_count)
    #str_list = []
    #for n in range(space_count):
    #   str_list.append(' ')
    #s = ''.join(str_list)
    #s = ''.join([' ' for n in range(space_count)])

def format_indented_msg(msg, offset):
    return '{0}^   {1}'.format(spaces(offset + len(MSG_PROMPT) - 1), msg)

def print_indented_msg(msg, offset):
    print(format_indented_msg(msg, offset))

#### New stuff: error messages!

MSG_NEW_SYLLOGISM = 'Begin new syllogism'
MSG_USAGE_HINT = 'Enter HELP for list of commands'
MSG_STOPPED = '(Some versions support typing CONT to continue)'
MSG_NO_PREMISES = 'No premises'
MSG_LINK_SUGGEST = 'Suggestion: try the LINK or LINK* command.'

MSG_INDENTED_RESERVED = '\nReserved word "{0}" may not occur within a term'
MSG_INDENTED_IGNORED = 'Punctuation mark ignored'
MSG_INDENTED_INVALID = 'Invalid numeral or command'

MSG_MISSING_COPULA_IS_ARE = '** Missing copula is/are'
MSG_SUBJECT_TERM_BAD_OR_MISSING = '** Predicate term bad or missing'
MSG_PREDICATE_TERM_BAD_OR_MISSING = '** Predicate term bad or missing'

COPYRIGHT_BLURB = """Syllogism Program Copyright (c) 1988 Richard Sharvy
Syllogism 1.0 (c) 2002 Richard Sharvy's estate
Ben Sharvy: luvnpeas99@yahoo.com or bsharvy@efn.org"""

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

RESERVED_WORDS = (
    'somebody',
    'something',
    'nobody',
    'nothing',
    'someone',
    'everyone',
    'everybody',
    'everything',
)

# These are more experimental

RESERVED_TERMS_POSITIVE = (
    'all',
    'some',
)

RESERVED_TERMS_NEGATIVE = (
    'no',
    'not',
)

RESERVED_TERMS_EQUALITY = (
    'is',
    'are',
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

#class Proposition(object):
#    # Per the square of opposition
#    # http://en.wikipedia.org/wiki/Square_of_opposition
#    pass
#
#class UniversalAffirmativeProposition(Proposition):
#    # An "A" proposition in the square of opposition
#    # All S are P
#    # Every S is a P
#    pass
#
#class UniversalNegativeProposition(Proposition):
#    # An "E" proposition in the square of opposition
#    # No S is/are P
#    pass
#
#class ParticularAffirmativeProposition(Proposition):
#    # An "I" proposition in the square of opposition
#    # Some S is P
#    # Particular affirmative
#    pass
#
#class ParticularNegativeProposition(Proposition):
#    # An "O" proposition in the square of opposition
#    # Some S is not P
#    pass


class PremiseToken(UserString):
    #	rem T(): 1:line num., 2:"/", 3:quantifier, 4:no/not, 5:is/are, 6:term
    QUANTIFIER = 3
    NO_NOT = 4
    IS_ARE = 5
    TERM = 6
    
    # 2 = '/'
    # 3 = line number

#class LineNumberToken(PremiseToken):
#    # 3: line number
#    def __repr__(self):
#        return 'L=' + self.raw
#
#class SlashToken(PremiseToken):
#    # 2: "/"
#    def __repr__(self):
#        return '/=' + self.raw


class QuantifierToken(PremiseToken):
    category = PremiseToken.QUANTIFIER

    def __repr__(self):
        return 'Q=' + super().__repr__()


class NegationToken(PremiseToken):
    category = PremiseToken.NO_NOT

    def __repr__(self):
        return 'N=' + super().__repr__()


class EqualityToken(PremiseToken):
    category = PremiseToken.IS_ARE

    def __repr__(self):
        return 'E=' + super().__repr__()


class TermToken(PremiseToken):
    category = PremiseToken.TERM

    def __repr__(self):
        return 'T=' + super().__repr__()


class Premise(UserString):
    """ This object represents a parsed line. """

    def __init__(self, seq):
        seq = seq.strip()
        super().__init__(seq)
        components = seq.split()
        try:
            self.line_number = int(components[0])
        except ValueError:
            raise ValueError(format_indented_msg(MSG_INDENTED_INVALID, len(components[0])))
        self.statement = ' '.join(components[1:])
        tokens = self.parse(components[1:])
        self.validate(tokens)
        # self.tokens = tokens

    def parse(self, components):
        """ Return a list of tokens corresponding to each component passed in """
        tokens = []
        for t in components:
            # if t.isdigit():
            #     tokens.append(LineNumberToken(t))
            # elif t == '/':
            #     tokens.append(SlashToken(t))
            if t in ('all', 'some', 'no'):
                tokens.append(QuantifierToken(t))
            elif t in ('is', 'are'):
                tokens.append(EqualityToken(t))
            elif t in ('not'):
                tokens.append(NegationToken(t))
            else:
                tokens.append(TermToken(t))
        return self.elide_terms(tokens)
    
    def elide_terms(self, tokens):
        # Elide consecutive terms into single terms
        elided_tokens = []
        prev_term = None
        for t in tokens:
            if t.category == PremiseToken.TERM:
                if prev_term:
                    prev_term += ' ' + t
                else:
                    elided_tokens.append(t)
                    prev_term = t
            else:
                elided_tokens.append(t)
                prev_term = None
        return elided_tokens

    def validate(self, tokens):
        """ Validate the tokens """
        #All    <general term #1>   is/are       <general term #2>
        #Some   <general term #1>   is/are       <general term #2>
        #Some   <general term #1>   is/are not   <general term #2>
        #No     <general term #1>   is/are       <general term #2>
        #
        # <designator>      is/are       <general term>
        # <designator>      is/are not   <general term>
        # <designator A>    is/are       <designator B>
        # <designator A>    is/are not   <designator B>
        print(tokens)
        #valid_forms = (
        #    (PremiseToken.QUANTIFIER_ALL, PremiseToken.TERM, PremiseToken.IS_ARE, PremiseToken.TERM),
        #    (PremiseToken.QUANTIFIER_SOME, PremiseToken.TERM, PremiseToken.IS_ARE, PremiseToken.TERM),
        #    (PremiseToken.QUANTIFIER_SOME, PremiseToken.TERM, PremiseToken.IS_ARE, PremiseToken.NO_NOT, PremiseToken.TERM),
        #    (PremiseToken.QUANTIFIER_NO, PremiseToken.TERM, PremiseToken.IS_ARE, PremiseToken.TERM),
        #    (PremiseToken.TERM, PremiseToken.IS_ARE, PremiseToken.TERM),
        #    (PremiseToken.TERM, PremiseToken.IS_ARE, PremiseToken.NO_NOT, PremiseToken.TERM),
        #)
        #
        #is_valid = True
        #for f in valid_forms:
        #    is_valid = True
        #    for z in zip(f, tokens):
        #        if z[0] != z[1].category:
        #            is_valid = False
        #            break
        #    if is_valid:
        #        break
        #print('isvalid = ' + str(is_valid))
        
        #proposition_types = {
        #    'all': UNIVERSAL_AFFIRMATIVE,
        #    'no': UNIVERSAL_NEGATIVE,
        #    'some': PARTICULAR_AFFIRMATIVE,
        #    'some': PARTICULAR_NEGATIVE,
        #}
        
        # elif t in RESERVED_TERMS_POSITIVE:
        #     if tj == 6:
        #         raise ValueError(format_indented_msg(MSG_INDENTED_RESERVED.format(t), len_covered + len(t)))
        #     else:
        #         tj = 3
        # elif t in RESERVED_TERMS_NEGATIVE:
        #     pass
        # elif t in RESERVED_TERMS_EQUALITY:
        #     pass
        
        # statement_type = None
        #len_covered = len(MSG_PROMPT) # account for prompt
        ##	rem T(): 1:line num., 2:"/", 3:quantifier, 4:no/not, 5:is/are, 6:term
        #is_valid = True
        #valid_next_tokens = (PremiseToken.QUANTIFIER, PremiseToken.TERM)
        #for t in tokens:
        #    if t.category in valid_next_tokens:
        #        if t.category == PremiseToken.QUANTIFIER:
        #            valid_next_tokens = (PremiseToken.TERM,)
        #        elif t.category == PremiseToken.IS_ARE:
        #            valid_next_tokens = (PremiseToken.TERM, PremiseToken.NO_NOT)
        #        elif t.category == PremiseToken.NO_NOT:
        #            valid_next_tokens = (PremiseToken.TERM,)
        #        elif t.category == PremiseToken.TERM:
        #            valid_next_tokens = (PremiseToken.IS_ARE,)
        #            if t.raw in RESERVED_WORDS:
        #                raise ValueError(format_indented_msg(MSG_INDENTED_RESERVED.format(t), len_covered + len(t)))
        #        len_covered += len(t)
        #    else:
        #        # Invalid pattern
        #        if PremiseToken.IS_ARE in valid_next_tokens:
        #            raise ValueError(MSG_MISSING_COPULA_IS_ARE)
        #        is_valid = False
        #        #raise ValueError(MSG_SUBJECT_TERM_BAD_OR_MISSING)
        #        #raise ValueError(MSG_PREDICATE_TERM_BAD_OR_MISSING)
        #print 'is valid = ' + str(is_valid)


class Rubric(UserList):
    """ This object represents a collection of parsed lines. """

    def enter_premise(self, premise):
        """ Try to add a premise to our lookup table.
        
        Parameters
        ----------
        premise : Premise
                  A premise to add to our lookup table.
        """
        if not premise.statement:
            self.remove_line(premise.line_number)
        else:
            # Remove any lines with the same line number
            try:
                self.remove_line(premise.line_number)
            except ValueError:
                # Let it slide: We don't need to worry if the line doesn't exist
                pass
            self.append(premise)
            # If a statement was included, replace the existing line
            # Otherwise, simply don't put the statement into place

            # Sort the new premises by line number
            # We only need to do this when adding lines                
            self.sort(key=lambda p: p.line_number)

    def remove_line(self, line_number):
        """ Remove a premise (identified by line number) from the lookup table.
        
        Parameters
        ----------
        line_number : integer
                      A premise line number for whose existence to check.
        
        [TODO] could also use OrderedDict...

        """
        if self.line_exists(line_number):
            for p in self:
                if p.line_number == line_number:
                    self.remove(p)
        elif len(self) == 0:
            # No premises have been entered
            raise ValueError(MSG_NO_PREMISES)
        else:
            # The premise to remove did not exist
            raise ValueError("Line {0} not found".format(line_number))

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
        for p in self:
            if p.line_number == line_number:
                return True
        return False

    def p(self, analyze=False):
        """ Printable format """
        lines = []
        last_premise = self[-1:]
        if len(last_premise) > 0:
            max_padding_chars = len(str(last_premise[0].line_number))
            # Format lines with nice spacing
            premise_groups = ((p.line_number, p.statement) for p in self)
            for p in premise_groups:
                if not analyze:
                    lines.append(' {0} {1}'.format(str(p[0]).rjust(max_padding_chars), p[1]))
                else:
                    lines.append(u" [TODO] Distribution analysis listing")
        return '\n'.join(lines)

    def __repr__(self):
        return self.p()

class Syllogism(object):

    def __init__(self):
        self.show_messages = True
        self.rubric = Rubric()

    def run(self):
        # Clear the screen
        self.cls()

        # Print copyright message
        print(COPYRIGHT_BLURB)
        print
        
        # Print a usage hint
        self.print_hint()
        
        # Create a new document: Currently unnecessary
        #self.new_syllogism(silent=True)
        
        # Start the request loop
        self.request_input()

    def cls(self):
        # clear the screen
        os.system("clear")

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
            line = input(MSG_PROMPT).lower()
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
                    try:
                        self.enter_line(line)
                    except ValueError as e:
                        print(e)
        self.print_message(MSG_STOPPED)

    def print_message(self, msg):
        """ Print a message if `show_messages` is enabled """
        if self.show_messages:
            print(msg)

    def toggle_messages(self):
        """ Toggle the state of certain messages """
        self.show_messages = not self.show_messages
        if self.show_messages:
            state = 'on'
        else:
            state = 'off'
        print('Messages turned {0}'.format(state))

    def strip_string(self, s):
        """ Remove punctuation from the end of a string.  Used in input processing.

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
            print_indented_msg(MSG_INDENTED_IGNORED, len(s))
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
        print(spaces(28) + "Examples:   10  All men are mortal")
        print(spaces(40) + "10")
        print("  DUMP" + spaces(15) + "Prints symbol table, distribution count, etc.")
        print("  HELP" + spaces(15) + "Prints this list")
        print("  INFO" + spaces(15) + "Gives information about syllogisms")
        print("  LIST" + spaces(15) + "Lists premises")
        print("  LIST*" + spaces(14) + "Same, but displays distribution analysis:")
        print(spaces(25) + "distributed positions marked with '*', ")
        print(spaces(25) + "designators marked with '+'")
        print("  LINK" + spaces(15) + "Lists premises in order of term-links (if possible)")
        print("  LINK*" + spaces(14) + "Same, but in distribution-analysis format")
        print("  MSG" + spaces(16) + "Turns on/off Printing of certain messages and warnings")
        print("  NEW" + spaces(16) + "Erases current syllogism")
        print("  SAMPLE" + spaces(13) + "Erases current syllogism and enters sample syllogism")
        print("  STOP" + spaces(15) + "Stops entire program")
        print("  SUBSTITUTE" + spaces(9) + "Allows uniform substitution of new terms in old premises")
        print("  SYNTAX" + spaces(13) + "Explains statement syntax, with examples")
        print("  /" + spaces(18) + "Asks program to draw conclusion")
        print("  /  <statement>" + spaces(5) + "Tests  <statement>  as conclusion")
        print(spaces(25) + "Note: this can be done even if there are no premises")
    
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
        return ' '.join(words_out)

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
        self.new_syllogism(silent=True)
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
        if line.startswith('/'):
            # [TODO] Evaluate
            pass
        else:
            premise = Premise(line)
            self.rubric.enter_premise(premise)

    def new_syllogism(self, silent=False):
        """ Remove all premises from the rubric. """
        if not silent:
            print(MSG_NEW_SYLLOGISM)
        self.rubric.clear()

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
