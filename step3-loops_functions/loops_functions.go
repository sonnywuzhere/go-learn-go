package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello loops and functions!")

	// Loops in Go 
	// https://go.dev/tour/flowcontrol/1

	// fun fact about Go: it only has for loops 
	// you initialize a variable, state your condition, and then iterate the variable 
	// This will print out 
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// but don't worry about not having a while keyword!
	// you can use the for keyword essentially as a while loop 
	// so that it loops through the condition while the condition is true 
	teaInventory := 10
	for teaInventory > 0 {
		fmt.Println("Order up!")
		teaInventory--
	}

	// if for some strange reason, you want a compact 
	// infinte loop without having True be the condition, 
	// Go's got you covered. 
	// uncomment the next two lines for your infinite loop
	// for {
	// }

	
	// Defer 
	// Go has a defer keyword, which means that a function will 
	// be executed after all the other functions have been returned 
	defer fmt.Println("I am printed out last")
	
	// This example calls the multiply function 
	// and multiplies two numbers together and returns the product 
	output := multiply(10, 7)
	fmt.Println("Your output is ", output)

	pourTea(16, "chai")
	buyTea(12, 5)
	listTea()
	recursiveTea(4)
	
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
	teaName, teaBrand := splittingTea(teaData[3])
	fmt.Println("The tea flavor is", teaName, "and the brand is", teaBrand)

	myTea := "Chai"
	fmt.Println("My tea is", myTea)
	attemptToStealTheTea(myTea)
	fmt.Println("My tea is", myTea)

	actuallyStealTheTea(&myTea)
	fmt.Println("My tea is: ", myTea)

}
// Functions 
// you need to specify what type you want your parameters to be 
// and you specify the type of the return value 

func pourTea(oz int, flavor string) int {
	fmt.Println("Here is your ", oz, "oz ", flavor, " tea!")
	return 1
}

// however, you don't have to specify the type of the 
// return value if you don't return anything from 
// the function 
func moreTea(oz int, flavor string)  {
	fmt.Println("Here is your ", oz, "oz ", flavor, " tea!")
}


// if the paremeters have the same type, then you 
// can do the shortened verstion of 
func buyTea(oz, price int) int {
	fmt.Println(price)
	return oz
}

// but you're not limited to just returning one thing 
// you can return multiple things 
func listTea() (string, string, string) {
	return "Multiple", "Different", "Teas"
}

func multiply(x, y int) int {
	output := x * y
	return output
}

// Recursion in Go 
// https://www.w3schools.com/go/go_function_recursion.php

// whether you love it, are intimidated by, or just like it 
// for the memes, recursion can be a part of prgramming 
// and Go does support it, so it's your lucky day!
func recursiveTea(tea int) string {
	if tea == 0 {
		return "out of stock"
	} 
	fmt.Println("Here is your order!")
	return recursiveTea(tea - 1)
}

// What if you had a string of a tea name with the tea brand 
// separated with a comma. 
// Such as the following: 
// Vanilla Caramel Chai,Tazo

// how would you split these strings into 
// the tea flavor and the tea brand? 

// https://www.geeksforgeeks.org/go-language/how-to-split-a-string-in-golang/
// first you need to make sure the strings package 
// is imported (look at the top of the file for imports)
// this package contains the split function 
// because doing this manually would be tricky because 
// every character in a string is represented by a byte(s)
// with UTF-8 encoding. So it's pretty complicated. 
// But luckily, someone has solved this problem for us
// so we can benefit from their coding
func splittingTea(tea string) (string, string) {
	splitTea := strings.Split(tea, ",")
	teaName := splitTea[0]
	teaBrand := splitTea[1]
	if strings.Contains(tea, "Tazo") {
		fmt.Println("Tazo for the win!")
	}
	return teaName, teaBrand
}

// Pass by Reference vs. Value 
// https://www.geeksforgeeks.org/go-language/function-arguments-in-golang/

// the default of Go is pass by value 
// this means that if you try to modify a variable outside 
// of a variable inside of a variable, it won't change 
// the outside variable 
func attemptToStealTheTea(tea string) bool {
	tea = "London Fog, obviously"
	return false
}

// Well that's pretty standard
// Any way we can do something more exciting with functions?

// Yes, you can pass by reference. 
// which means changing external variables inside 
// the function (sneaky!), then you can add some 
// cryptic symbols to make this happen
// for your parameter, add the * symbol before your type
// and then add it again before the parameter referenced
// inside of your function 
// when you call the function, use the & symbol 
// and this will allow your function to modify stuff
// outside of the function 
// This happens becuase you're passing the actual pointer 
// to the function, so you're both pointing to the same thing 
func actuallyStealTheTea(tea *string) bool {
	*tea = "I stole your tea!"
	return true
}