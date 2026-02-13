package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello loops and functions!")

	teaData := [11]string{
		"Refresh Mint Biologique,Tazo", 
		"Earl Grey Biologique,Tazo",
		"Joy,Tazo",
		"Pumpkin Spice Chai,Tazo",
		"Glazed Lemon Loaf,Tazo",
		"Passion,Tazo",
		"Green Ginger,Tazo",
		"Peachy Green,Tazo",
		"Vanilla Caramel Chai,Tazo",
		"Vanilla Bean Macaron,Tazo",
		"Chai,Tazo",
	}

	// Loops in Go 
	// https://go.dev/tour/flowcontrol/1

	// fun fact about Go: it only has for loops 
	// you initialize a variable, state your condition, and then iterate the variable 
	// This will loop through and print out the tea in the teaData array above
	for i := 0; i < len(teaData); i++ {
		fmt.Println(i, teaData[i])
	}

	// Are you feeling sad that there's no while loop?
	// Good news: you can use the for keyword essentially as a while loop 
	// so that it loops through the condition while the condition is true 
	teaInventory := len(teaData)
	for teaInventory > 0 {
		fmt.Println("Order up!", teaData[teaInventory - 1])
		teaInventory--
	}

	// the for keyword can also be used as a for each loop 
	// this uses the range keyword and loops through 
	// can be used to loop through each element in a list
	// https://golangtutorial.dev/tips/foreach-loop-go/
	for index, element := range teaData {
		fmt.Println(index, element)
	}

	// Go doesn't have a do-while loop, but you can do some tricks
	// to turn the for loop into one. See the below article.
	// https://gosamples.dev/do-while/

	// If for some strange reason, you want a compact 
	// infinte loop without having True be the condition, 
	// Go's got you covered. 
	// uncomment the next two lines for your infinite loop
	// for {
	// }

	// Functions 
	// fun fact about Go: to declare a function inside of another function, 
	// you have to be sneaky and have it be in a variable like so
	var sneakyFunction = func() string {
		return "I am sneaky"
	}
	disguise := sneakyFunction()
	fmt.Println(disguise)

	// Defer 
	// Go has a defer keyword, which means that a function will wait to
	// be executed after all the other functions have been returned 
	var politeFunction = func() {
		fmt.Println("I am printed out last")
	}
	
	defer politeFunction()

	// The following code is examples of calling functions that are written 
	// outside the main function below 

	// This example calls the pourTea function 
	// and shows that you can pass different types of variables
	// as arguments into a function 
	pourTea(16, "chai")

	// This example of using shorthand when you pass two 
	// variables of the same type into a function 
	buyTea(12, 5)

	// This is an example of returning more than one value from 
	// a function 
	tea, moreTea, evenMoreTea := listTea()
	fmt.Println(tea, moreTea, evenMoreTea)

	// This example calls the multiply function 
	// and multiplies two numbers together and returns the product 
	output := multiply(10, 7)
	fmt.Println("Your output is ", output)

	// This function is an example of recursion
	recursiveTea(3)
	
	// This code calls the splittingTea function which 
	// takes in a string and returns a split string 
	teaName, teaBrand := splittingTea(teaData[3])
	fmt.Println("The tea flavor is", teaName, "and the brand is", teaBrand)

	// This code is an example of pass-by-value which is 
	// the default way in Go where myTea is not modified 
	// after being passed into the function 
	myTea := "Chai"
	fmt.Println("My tea", myTea)
	attemptToStealTheTea(myTea)
	fmt.Println("My tea", myTea)

	// This code is an example of pass-by-reference
	// where myTea is modified after being passed into 
	// the function because it points to the actual address
	// of the value and makes it capable of changing it
	actuallyStealTheTea(&myTea)
	fmt.Println("My tea: ", myTea)

}

// Functions 
// To create a function, you use the func keyword followed by the name
// of your function, parentheses () 
// you need to specify what type you want your parameters to be 
// and you specify the type of the return value 
func pourTea(oz int, flavor string) int {
	fmt.Println("Here is your ", oz, "oz ", flavor, " tea!")
	// this function only returns 1 because it needs to return an int
	return 1
}

// However, the return type is optional and you can decide not to 
// return anything from your function 
func moreTea(oz int, flavor string)  {
	fmt.Println("Here is your ", oz, "oz ", flavor, " tea!")
}

// if the paremeters have the same type, then you can do some 
// fancy shorthand where separate the parameters with commas, 
// and then put the shared type 
func buyTea(oz, price int) int {
	fmt.Println("The price of your", oz, "oz tea is $", price)
	return oz
}

// but you're not limited to just returning one thing 
// you can return multiple things 
// this function returns three strings, but it doesn't have 
// to be the same types. You could vary your types.
func listTea() (string, string, string) {
	return "Tea", "More Tea", "Even More Tea"
}

// This is a classic multiplication function 
// ... in case you don't want to open up Google 
// or your phone app or your brain's working memory to 
// multiple your two numbers 
func multiply(x, y int) int {
	output := x * y
	return output
}

// Recursion in Go 
// https://www.w3schools.com/go/go_function_recursion.php

// whether you love it, are intimidated by, or just like it 
// for the memes, recursion is a part of programming 
// and Go does support it, so it's your lucky day!
func recursiveTea(tea int) string {
	if tea == 0 {
		return "out of stock"
	} 
	fmt.Println("Here is your order!")
	return recursiveTea(tea - 1)
}

// Function that splits a string  

// What if you had a string of a tea name with the tea brand 
// separated with a comma. Such as the following: 
// Vanilla Caramel Chai,Tazo

// how would you split these strings into 
// the tea flavor and the tea brand? 

// first you need to make sure the strings package 
// is imported (look at the top of the file for imports)
// This package contains the split function 
// because doing this manually would be tricky because 
// every character in a string is represented by a byte(s)
// with UTF-8 encoding... it's pretty complicated. 
// But luckily, someone has solved this problem for us
// so we can benefit from their blood, sweat, and tears
// in coding this up just for us
// https://www.geeksforgeeks.org/go-language/how-to-split-a-string-in-golang/
func splittingTea(tea string) (string, string) {
	splitTea := strings.Split(tea, ",")
	teaName := splitTea[0]
	teaBrand := splitTea[1]
	return teaName, teaBrand
}

// Pass by Reference vs. Value 
// https://www.geeksforgeeks.org/go-language/function-arguments-in-golang/

// In functions, the default of Go is pass-by value 
// This means that if you try to modify a variable outside 
// of a function inside a function, it won't change 
// the outside variable 
func attemptToStealTheTea(tea string) bool {
	tea = "I'm trying to steal your tea!"
	return false
}

// Well that's pretty standard
// Any way we can do something more exciting with functions?

// Yes, you can do something called pass by reference. 
// This means the ability to change external variables inside 
// the function (sneaky!). To do this, you need to add some 
// cryptic symbols to make this happen.
// For your parameter, add the * symbol before your type
// and then add it again before the parameter referenced
// inside of your function 
// When you call the function, use the & symbol in front of 
// your argument.
// This will allow your function to modify that argument's 
// value outside of the function 
// This happens becuase you're passing the actual pointer 
// to the function, so you're both pointing to the same address 
// in memory and this way can change the actual value
func actuallyStealTheTea(tea *string) bool {
	*tea = "I stole your tea!"
	return true
}