package main

import "fmt"

func main() {
	fmt.Println("Time to learn about if/else, etc... aka, the programmy stuff of programming")

	// let's create some boolean values to use in our if/else statements
	var isCoding bool = true 
	isThirsty := true 
	isTired := true
	isWarmOutside := false
	
	// Note, you have to use true or false (lowercase) 
	// and cannot use 0 or 1 
	// this code will crash the code
	// var doesNotWork bool = 0
	// var noCanDo bool = True

	// If statements 

	// a simple, boring, vanilla if statment 
	// If the condition is true, the if block runs 
	// otherwise, the else code block runs 
	// Go doesn't require parenthesis around the condition
	// but the compiler won't complain if you do put parentheses 
	if isWarmOutside {
		fmt.Println("Go take a walk")
	} else {
		fmt.Println("Stay inside and get some tea to warm up!")
	}

	// What happens if we create two if statments here 
	// and then create an else statment by itself
	// This is called the "dangling else" problem 
	
	// showing parenthesis around the condition here 
	// to show that the compiler won't complain 
	// the ! is called a Logical NOT and makes whatever 
	// the boolean is the opposite (if true, then false
	// and vice versa)
	if (!isCoding) {
		fmt.Println("You should be coding")
	}

	if !isThirsty {
		fmt.Println("Great job hydrating today!")
	}


	// having a "dangling else" statment just hanging out here
	// will crash the code. The compiler doesn't even try to 
	// associate it with a previous if statement. 
	// else {
	// 	fmt.Println("Generic statement")
	// }

	// to use multiple conditions in an if statement, 
	// go uses else if 
	if isTired || isThirsty {
		fmt.Println("You should take a break")
	} else if isTired && isCoding {
		fmt.Println("You really need to take a break")
	} else if isCoding && isThirsty {
		fmt.Println("Go make a London fog")
	} else if isTired && isWarmOutside {
		fmt.Println("Take a walk outside")
	} else {
		fmt.Println("Get yourself your favorite tea!")
	}
	
	// Switch Statments 

	// a switch statement is an elegant way to evaluate 
	// many different cases for an expression 
	// The great thing about Go is that it only runs the code 
	// block for the matched case. 
	// This means no need for a break statement! 
	// Kudos to Go for not breaking the Flow of Program Control! 
	
	// first initialize a variable that will be evaluated 
	// Since tea selection doesn't have a value, it will print the default
	// upon initialization 
	var teaSelection int 
	// But if we give it a value, then it will print that case 
	teaSelection = 4
	
	switch teaSelection {
	case 1:
		fmt.Println("Strawberry matcha latte")
	case 2: 
		fmt.Println("Matcha latte")
	case 3: 
		fmt.Println("Chai tea")
	case 4: 
		fmt.Println("Candy cane lane")
	case 5: 
		fmt.Println("Elderberry ginger")
	case 6: 
		fmt.Println("Hibiscus Rose")
	case 7: 
		fmt.Println("Sweet Cinnamon")
	// all the case values must be the same type as the 
	// switch expression, so we couldn't just add this
	// to the mix 
	// case "I Love tea":
	// 	fmt.Println("Join our tea club!")
	default: 
		// the default case is the "catch all" if none 
		// of the other cases are run
		// ... and let's be honest, London Fogs are q
		// fantastic go-to tea 
		fmt.Println("London fog")
	}

	// You don't have to use just integers for switch 
	// statements. You can use strings, bools, or floats, 
	// it just has to be consistent for the swtich statement 
	var teaSuggestions string = "chai"
	switch teaSuggestions {
		case "chai":
			fmt.Println("Try rose chai or chocolate chai for a unique twist") 
			// if you use fallthrough, it will go to the next case automatically
			// so this will run the next code block as well
			fallthrough
		case "hibiscus":
			fmt.Println("Try hibiscus with coconut")
		case "green":
			fmt.Println("Try a ginger mojito")
		default: 
			fmt.Println("Try a London fog with lavender honey syrup")
	}

	// Short circuit logic means that in programming, 
	// there's a shortcut for evaluating logic expressions
	// because again, programmers love effeciency 

	// this can show up in two different scenarios 
	// So you don't have to scroll up to remember, 
	// isWarmOutside is false and isTired is true 
	// because this is an AND (&&) operator 
	// Because isWamOutside is false, Go won't even bother 
	// to check isTired, because false AND true can't be true
	// it will immediately jump to the else code block 
	// ... and who wants to take a walk outside when it's 
	// not warm out?  
	if isWarmOutside && isTired {
		fmt.Println("Go take a walk outside")
	} else {
		fmt.Println("Stay inside and make some warm tea")
	}

	// this works if the conditions are flipped as well
	// because isTired is true, Go then evaluates the second
	// conditional as well. Since it's false, it goes to the 
	// else block
	if isTired && isWarmOutside {
		fmt.Println("How about taking a walk outside?")
	} else {
		fmt.Println("Make sure to stay inside and make some warm tea")
	}

	// the other scenario is in an OR logic expression 
	// here, isThirsty is true and isCoding is true 
	// because it's an OR expression, either isThirsty OR 
	// isCoding needs to be true to run the code block
	// Because isThirsty is true, Go doesn't need to check 
	// isCoding because one of the conditions is true
	if isThirsty || isCoding {
		fmt.Println("You should make a ", teaSelection)
	} else {
		fmt.Println("Make some ")
	}

	// this will still be the case if isCoding is false 
	// because only one of the conditions needs to be true
	isCoding = false 
	if isThirsty || isCoding {
		fmt.Println("You should make your favorite tea")
	} else {
		fmt.Println("Probably still make your favorite tea :)")
	}
	
}