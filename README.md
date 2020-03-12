# Syllogism

[Syllogism](http://members.efn.org/~bsharvy/rsharvy/syll.html), an interactive BASIC program developed by Richard Sharvy, parses a user-input set of logical premises and then attempts to determine their validity.  The source code is available online and should work, at least with minimal modifications, in a variety of freely-available interpreters.  As linked above, the website for the official distribution of the program is currently at:

<http://members.efn.org/~bsharvy/rsharvy/syll.html>


## Why does this repository exist?

I really like this admittely-historical program, which despite its choice of language (BASIC) contained a fascinatingly effective English language processor and tokenizer.  In an effort to future-proof the program---and perhaps simply for the personal satisfaction---I'm attempting to making a port to Golang or Python.

The original license---taken from the Syllogism website---is as follows:

> Syllogism 1.0, as source code or compiled program, costs zero dollars.
> You may distribute it. You may distribute altered versions provided:
> 
> 1. they work,
> 2. there is clear and prominent notification that you altered the 
>    program and how,
> 3. there is clear and prominent notification of the original
>    authorship.

For those requiring a more formal license, Ben Sharvy has graciously offered the source code under the GNU GPL v3.  Any derivatives of the original code base that I place here, regardless of coding language, are released under the same.

## How does this program differ from the original program?

1. The program is written in Golang, not BASIC.
2. My name is appended to the end of the copyright notice to designate that I ported the program to Golang.
3. Program execution flow (e.g., `STOP`/`CONT`) is supplanted by modern terminal program flow (e.g., when you exit, the program exits completely). Press `ctrl-z` to suspend the program without closing in a particular terminal session.
4. Some slight changes to help printouts, primarily spacing.
