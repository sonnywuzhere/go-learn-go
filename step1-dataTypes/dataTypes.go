package main

import (
	"fmt"
)

// Based on Go Documentation: https://go.dev/tour/welcome/1

// can "batch" vars together
var (
	// if you're wondering what to use for a number, just use int
	// because if you have a use case where you need to use something
	// really specific, then you wouldn't be wondering
	
	num int = 9223372036854775807
	
	// fancy integers and unsigned integers with their maximum
	// provided for your enjoyment
	maxInt8 int8 = 127
	maxInt16 int16 = 32767
	maxInt32 int32 = 2147483647
	maxInt64 int64 = 9223372036854775807
	// unsigned integers
	unni uint = 18446744073709551615
	maxuInt8 uint8 = 255
	maxUint16 uint16 = 65535
	maxUint32 uint32 = 4294967295
	maxUint64 uint64 = 18446744073709551615
	
	// Go also has aliases for two data types
	// a rune is an alias for an int32
	runey rune = 2147483647
	// a byte is alias for a uint8
	lilByte byte = 255
	
	// You can also initialize a pointer, but we'll come back to this later
	pointy uintptr
	
	// Floats are always important, and there are two flavors in Go 
	// Single precision
	floaty1 float32 = 3.1415926
	// Double precision 
	// Defaulted float in Go if not specified, because more precision = better
	// We're programmers after all. We want allll the control.
	floaty2 float64 = 3.141592653589793
	
)

func main() {
	fmt.Println("Time to explore variables and data types!")
	
	// naming conventions for variables
	fmt.Println("The naming convention for variables is to use MixedCaps or mixedCaps")
	fmt.Println("Go doesn't use underscores for variable names (sorry, Python fans)")

	var dont_do_this string = "the compiler won't judge you, but other programmers will"
	fmt.Println(dont_do_this)

	var _alsoAvoid string = "the compiler will run this, but you'll still get judged"
	fmt.Println(_alsoAvoid)
	
	// however, you cannot start variables with numbers or symbols 
	// the compiler will judge you in this case and crash 

	// these lines of code will crash the code 
	// var 2go string = "no go here"
	// var ^no string = "also no go here"
	// fmt.Println(2go)
	// fmt.Println(^no)

	// Note! If you delare and don't use variables, you will get 
	// an error in both your IDE and it won't compile your code 
	// because the compiler sees this as an error
	// Unused variables is something Go takes seriously, so you should too!

	// These all create integers 
	var number int = 3
	var otherNum = 5
	someNum := 10
	fmt.Println("These are all integers in Go: ", number, otherNum, someNum)
	var zeroNum int
	fmt.Println("This is an integer that was initialized but not given a value: ", zeroNum)

	// These all create booleans 
	var likesCoding bool = true 
	var likesGo = true
	spendsTooMuchTimeCoding := false
	fmt.Println("These are all bools in Go: ", likesCoding, likesGo, spendsTooMuchTimeCoding)
	var likesChocolate bool
	fmt.Println("This is a bool that was initialized but not given a value: ", likesChocolate)

	// These all create strings
	var firstName string = "Nancy"
	var lastName string = "Drew"
	occupation := "detective"
	fmt.Println("These are all strings: ", firstName, lastName, occupation)
	var nickName string
	fmt.Println("This string was initialized but not given a value, so it's an empty string: ", nickName)

	// This is a float 
	var rootBeerFloat float64 = 3.141592653589793
	fmt.Println("This is a floating point number: ", rootBeerFloat)

	// You can also initialize variables in groups
	// and use different data types
	one, two, three, four := true, "two", 3, 4.53
	fmt.Println("These were initialized together: ", one, two, three, four)


	// Now for the big question: can ints and floats add? 
	// Does go have narrowing or widening conversion?
	// i.e. does it chop off the decimal or keep the decimal? 
	fmt.Println("Go can add/multiply/divide ints and floats, and it results in a widening conversion")
	fmt.Println("This is an example of multiplying an int and a float: ", 5 * rootBeerFloat)
	fmt.Println("This means that the result becomes a float and the decimal is kept (precision for the win!)")
	fmt.Println("Look at all that precision!")


	// So how do you create arrays in Go? 
	// The Go documentation says "The type [n]T is an array of n values of type T"
	// https://go.dev/tour/moretypes/6
	// ... huh? What does that mean? 
	// Let's break it down 
	// In Go, you need to define how big the array is upfront (again, with the precision)
	// You need to declare the size of the array and what data types will be in the array

	// Let's see this in action!
	// There are two ways to create an array
	// One way is to declare the array and then add the values into it one by one
	var lilArray [4]string 
	lilArray[0] = "this"
	lilArray[1] = "is"
	lilArray[2] = "an"
	lilArray[3] = "array"
	fmt.Println(lilArray)

	moreEfficient := [4]string{"this", "is", "more", "efficient"}
	fmt.Println(moreEfficient)

	// So, if you have to specify which data type is in the array, 
	// can we put different data types into the array? 
	// Or do they have to be all the same? 
	// The data types all need to be the same, so the following 
	// lines of code are commented out otherwise the code would crash
	// var varietyPack [3]int
	// varietyPack[0] = true

	// The "dictionary" data type in Go is a map 
	// https://www.w3schools.com/go/go_maps.php
	// Like all the other data types, you can initialize a map in a couple of ways 

	// Here is the first way
	// The data is map, and then you specify the key,value pair data types as well
	// In this example, both the key and value are strings 
	var mapQuest = map[string]string{"city": "Paris", "country": "France"}
	// If you print out the map directly, the formatting is a bit funny
	fmt.Println(mapQuest)
	// If you want to access the elements in the Map, you can format it like this
	fmt.Println(mapQuest["city"], mapQuest["country"])

	// You can also create an empty map and fill it up later 
	// To do this, you need to use the make() function 
	// You still specify the map data type, then the key,value data types
	// In this example, the key is a string and the value is an boolean 
	var expensiveCities = make(map[string]bool)

	// what happens if you try to print out the empty map?
	fmt.Println("This is an empty map: ", expensiveCities)

	expensiveCities["NYC"] = true
	expensiveCities["Boston"] = true 
	expensiveCities["Decatur, IL"] = false 

	fmt.Println("This map is no longer empty: ", expensiveCities)

	// ------ OTHER DATA TYPES -------

	// Pointers 
	// Go has pointers which are the address in memory of a value 
	// This is an empty pointer 
	var pointy *int 
	// Spoiler: it points to <nil>
	fmt.Println("This is an empty pointer. I wonder what it points to: ", pointy)

	// So what do pointers do exactly?
	// Let's create a pointless statement
	pointlessStatement := "This is a pointless statement"
	// If we print this statement, it prints the string like expected
	fmt.Println(pointlessStatement)
	// But if we initialize another variable and add a funny symbol "&"
	// to the front of our pointlessStatement
	iAmPointing := &pointlessStatement
	// If we print this, now we see this complicated looking number
	// This is the pointer
	fmt.Println("This is the pointer: ", iAmPointing)
	// But if we add another funny symbol to the front of this, then 
	// we go to the memory address that the pointer points to
	fmt.Println("We're using the pointer to point to: ", *iAmPointing)

	// Slices 
	// https://go.dev/tour/moretypes/7
	// Slices provide a solution to the problem of arrays being fixed in size 
	// One of the great things about programming is having precision, 
	// but we also love flexibility in coding too 
	// A slice is just what it sounds like: a slice of an array 
	// Just like slicing up a mango, you can slice up a part of an array 
	// (and it's much easier than cutting up a mango... just saying)

	// so we start with an array of fruits
	fruits := [7]string{"strawberries", "blueberries", "mango", "mango", "mango", "dragonfruit", "kiwi"}
	// to slice out just the mangos, we create a variable that looks 
	// like creating an array, but doesn't specify the length
	// then we use [] with the lower and upper indeces with a colon 
	// separating the two numbers 
	var mangoSlices []string = fruits[2:5]
	// This slices out just the mangos from the fruit array
	fmt.Println(mangoSlices)
	// What happens if we change one of the indeces of the slice?
	mangoSlices[0] = "golden Mango"
	// this change is reflected in the slice, as expected
	fmt.Println(mangoSlices)
	// but it also is changed in our array we sliced from 
	fmt.Println(fruits)


	// Interfaces 
	// "A set of method signatures"
	// convention is to capitalize interface names 
	// the interface has the name of a method and its return type
	// https://go.dev/tour/methods/9
	// https://www.geeksforgeeks.org/go-language/interfaces-in-golang/

	// interfaces are more related to functions, so we'll revisit this 
	// later when we learn about functions in Go
	type BakedGoods interface {
		Cookies() string 
		Brownies() string 
		Cheesecake() string 
		CremeBrulee() string 
	}

	// Struct
	// https://www.w3schools.com/go/go_struct.php
	// The convention is to capitalize Structs 
	type Book struct {
		author string
		title string 
		pages int 
	}
	
	var book1 Book 
	book1.title = "Modern Operating Systems"
	book1.author = "Andrew S. Tanenbaum"
	book1.pages = 1076

	// You can print out the struct in it's entirety like this:
	fmt.Println(book1)
	// Or you can access each element individually
	fmt.Println(book1.author, "wrote the book", book1.title, "which has", book1.pages, "pages.")

	// Type conversions
	// Can you convert one data type into another?

	// You can convert an int into a float 
	var plainInt int = 10
	var fancyFloat float64 = float64(plainInt)
	fmt.Println("Converting int into a float", fancyFloat)

	// You can also convert a float into an int
	// it will just chop off the decimal point 
	var floatingAway float64 = 3.14
	var boringInt int = int(floatingAway)
	fmt.Println("Converting float into an int ", boringInt)

	// You cannot convert a string into an int directly
	// var fakeNumber string = "10"
	// var actualNumber int = int(fakeNumber)
	// fmt.Println("This is an actual int ", actualNumber)
	// However, you can use a package called strconv to do this

	// Cannot convert a bool to an int directly
	// var likesToWorkout bool = true 
	// var hoursWorkedOut int = int(likesToWorkout)
	// fmt.Println("Hours worked out: ", hoursWorkedOut)

	// Cannot convert an int to a bool directly
	// var num int = 1
	// var num1 bool = bool(num)
	// fmt.Println(num1)

	// Further experimentation!
	
	// Can you add a String and an int together? 
	// In Go, this is a no go
	// The following two lines are commented out because 
	// they'll throw and error and not compile if you try to run it
	// x = "5" + 6
	// fmt.Println("5" + 6)


	// Can you add a bool and an int?
	// Nope, also a no go
	// var likesTea bool = true 
	// var teaFlavors int = 20
	// fmt.Println("Adding a bool and int: ", likesTea + teaFlavors)

	// Can you add two bools together? 
	// This is also a no
	// var coldOutside bool = true 
	// var warmBlanketsInside bool = true 
	// var stayWarm bool = coldOutside + warmBlanketsInside
	// fmt.Println("Adding two bools: ", stayWarm)

}