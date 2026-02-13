# Loops & Functions

1. Does Go include multiple types of loops (while, do/while, for, foreach)? If
   so, what are they and how do they differ from each other?

   Go only has one loop: a for loop. This for loop can be used as either a traditional
   for loop (i.e. looping through a fixed number of something) or as a while loop
   (i.e. looping until a condition is no longer true... if you write your function
   correctly, that is, otherwise you've got an infinite loop)

2. What is the syntax for declaring a function in Go?

   To declare a function in Go, you use the func keyword, followed by the name
   of your function and then parentheses ().
   Inside the parentheses, you write your parameters, specifying their types.
   The type comes after the parameter.
   Then right next to the outside of the parentheses, you specify the return type.
   If you don't specify anything here, you can't return anything inside the function.
   Add some curly brakets, and voila! You have a Go function.

   The following function takes in a value of type int (integer). It them returns the argument,
   following the type that is required for a return value (integer).

```
   func pointlessFunction(someValue int) int {
        return someValue
   }
```

3. Are there any rules about where the function has to be placed in your code file so that
   it can run?

   Functions using the func keyword can't be placed inside of each other.
   If a function was a box, then all functions need to be in their own box.
   You can't nest the boxes inside of each other.

   Non-main functions can be placed either before or after the main function.
   There must be a main function in your code for it to run properly.

4. Does Go support recursive functions?

   Yes, Go supports recursive functions. See the loops_functions.go file for an example!

5. Can functions in Go accept multiple parameters? Can they be of different
   data types?

   Functions in Go do accept multiple parameters and they can be of different data types.
   See loops_functions.go file for examples.

6. Can functions in Go return multiple values at the same time? How is that
   implemented? If not, are there ways around that problem? What are they?

   Go can retun multiple values at the same time. Theoretically, you could return as many
   as you wanted and there's no limit... except the limit to computer memory, your patience,
   or your ability to leverage AI to do stuff for you

7. Is Go pass-by reference or value? Check your code against outside
   sources in case there is anything tricky going on (like in Perl).

   Go is pass-by value by default, but you can also pass-by reference if you so desire.
   See loops_functions.go file for examples of both.

8. Are there any other aspects of functions in Go that aren't specifically asked
   about here, but that are important to know in order to write one? What are they?

   You can't declare a function inside another function, but you can be sneaky and create a variable
   that is a function. For example, this code would not work:

```
   func main() {
        func notWorking() {
            fmt.Println("This isn't working")
        }
   }
```

    But if you have the function be a variable, then it does work

```
    func main() {
        var working = func() {
            fmt.Println("Now I'm working")
        }

        working()
    }
```

https://forum.golangbridge.org/t/writing-function-inside-and-outside-of-func-main/13268
