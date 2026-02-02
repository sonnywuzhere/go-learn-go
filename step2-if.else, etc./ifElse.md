# Step 2 - if/else, etc.

1.  What are the boolean values in your language? (e.g., True and False, true
    and false, 1, and 0, etc)

    The boolean values in Go are true and false (must be lowercase)

    You cannot use 1 and 0 interchangeably for true and false

2.  What types of conditional statements are available in your language? (if/else,
    if/then/else, if/elseif/else). Does your language allow for statements other than
    “if” (for example, Perl has an “unless” statement, which does the opposite of
    “if”!)

    Go uses if/else if/else for conditional statments.

    Sadly, Go doesn't have a ternary operator (? : ) for shorthand conditional statements

3.  Does your language use short-circuit evaluation? Include an example of the
    short-circuit logic working or not working (or both, if your language is like Java
    and supports both!)

    Go does use short-circuit evaluation

    Consider this example (which is also in the Go code example):

    var isWarmOutside bool = false
    var isTired bool = true

    For the AND (&&) operator, because isWamOutside is false,
    Go won't even bother to check isTired, because false AND
    true can't be true it will immediately jump to the else code block
    ... and who wants to take a walk outside when it's cold out?

    ```
    if isWarmOutside && isTired {
        fmt.Println("Go take a walk outside")
    } else {
        fmt.Println("Stay inside and make some warm tea")
    }
    ```

4.  How does your programming language deal with the “dangling else” problem?

    Go will just crash and refuse to compile if there is an else statement
    hanging out by itself on it's own line. The "else" must be on the same line
    as the closing curly bracket of the "if" or "else if" statment

5.  If your language supports switch or case statements, do you have to use
    “break” to get out of them? Can you use “continue” to have all of them
    evaluated?

    Go does support switch statements. You do not need to use "break" to get
    out of them (yay!). You can, however, still use break if you want to

    However, Go does have a "continue" called fallthrough to fall through
    (clever, right?) to the next case. It must be put last in the code block,
    so you can't write fallthough and then write more code in the case's
    code block

    You can uses any data type for the cases, but all the cases must be the
    same type as the expression, otherwise the compiler will throw an error

    https://go.dev/wiki/Switch
