# Missing digit kata

The goal of this kata is to find exactly one missing digit in
a given equation string. An equation string contains exactly
two numbers, one mathematical operand (+, -, * or /), an equation
sign (=) and one number with a missing digit. The number with a
missing digit can be identified by the "x" it contains. An example
input would be "2x + 5 = 30". The corresponding output would be
5 (as a number) because 25 + 5 equals 30.

# This solution

## Go vs. languages with eval (or macros)
There are various way to approach this problem. In most languages
which include an "eval" function (a function that takes a statement written
in this language and returns the result of this statement) the solution is
trivial.
In go where there is no such function. This is why this solution is a bit more involved.

## Methodology
The chosen solution was developed using TDD (although I got sloppy at the end) and extracts
tokens from the equation based on whitespaces.
It has a constant runtime (I mention this even though I cannot imagine a solution which has a
non-constant runtime) and 100% test coverage.

If you want to check how I use TDD you can go through the version history of this repo.
Note that I tried to follow TPP as defined by uncle Bob during implementation and test writing.
However, there are still some "jumps" or discontinuities which I didn't like.

## Closing remarks
I think the solution is far from perfect and could most likely be improved by functions operating
on an equation struct instead of passing around slices and returning lots of variables. This is
especially true for numbersWithFilledDigit which returns too many variables.

# Note
This kata wasn't invented by me. It has been on the internet for quite some time (as far as I can tell).
I worked on this during a code screening. However, the solution was created solely by me.
