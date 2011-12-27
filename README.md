[Syllogism](http://members.efn.org/~bsharvy/rsharvy/syll.html). an interactive BASIC program developed by Richard Sharvy, parses a user-input set of logical premises and then attempts to determine their validity.  The source code is available online and should work, at least with minimal modifications, in a variety of freely-available interpreters.  As linked above, the website for the official distribution of the program is currently at:

<http://members.efn.org/~bsharvy/rsharvy/syll.html>

For some reason, I really like this admittely-historical program.  In an effort to future-proof the program—and perhaps simply for the personal satisfaction—I'm attempting to making a port to Python.  In the process, I will be committing to this repository an updated BASIC version (with changes such as subroutines and/or named functions) that should hopefully compile (again, with minimal modifications) in other interpreters.  Testing for this particular variant will be performed under Mac OS X 10.7 in [Bas](http://www.moria.de/~michael/bas/) 2.2, an open-source BASIC interpreter.

The license—taken from the Syllogism website—is as follows:

> Syllogism 1.0, as source code or compiled program, costs zero dollars.
> You may distribute it. You may distribute altered versions provided:
> 
> 1. they work,
> 2. there is clear and prominent notification that you altered the 
>    program and how,
> 3. there is clear and prominent notification of the original
>    authorship.