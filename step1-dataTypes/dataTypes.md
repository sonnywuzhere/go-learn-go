# Step 1 - data types

1. What are the naming requirements for variables in Go?
   a. What about naming conventions? Are they enforced by the comiler/interpreter,
   or are they just standards in the community?

   The Go naming requirements are that variables are in mixedCaps.
   It's not enforced by the compiler (Go is a compiler-based language).

   However, you cannot start a variable name with either a number or symbol.
   This will crash your code.

   https://go.dev/doc/effective_go#names

2. Is your language statically or dynamically typed?

   Go is a statically typed language.

   Once a data type is specified, you can only put data of that type in the variable
   If you want to change the data type of a variable, you need to specifically do
   this through type conversion/type casting

   https://www.w3schools.com/go/go_data_types.php
   https://www.geeksforgeeks.org/go-language/type-casting-or-type-conversion-in-golang/

3. Strongly typed or weakly typed?

   Go is a strongly typed language.

   The compiler will enforce the rules of not combining data types that cannot be combined
   and will refuse to run (aka, crash).

   https://www.techtarget.com/whatis/definition/strongly-typed
   https://go.dev/ref/spec

4. If you put this line (or similar) in a program and try to print x,
   what does it do? If it doesn't compile, why? Is there something you can do to make it compile?
   x = "5" + 6

   You cannot add a string and an int in Go. It will not compile.
   However, you can use the standard library package called strconv to use type conversion between string and ints.

5. Describe the limitations (or lack thereof) of your programming language as they relate to
   the coding portion of the assignment (adding ints and floats, storing different types in lists,
   converting between data types). Are there other restrictions or pitfalls that the documentation
   mentions that you need to be aware of?

   Because Go is both statically typed and strongly typed, it's a bit picky in how you handle
   initializing your data types and having them interact with each other.

   For example, in Go, you cannot have an array of different data types. This could be a limitation
   depending on what you are trying to do.

6. Are there built-in complex data types that are commonly used in your language? (hint: they'd
   probably appear fairly early in the documentation if so)

   Some of the built-in complex data types are Maps (like dictionaries in Python), Structs,
   Interfaces, Slices, Arrays, and Pointers

   See the [step1-dataTypes Go code](https://github.com/sonnywuzhere/go-learn-go/blob/main/step1-dataTypes/dataTypes.go) for more details on each of these
