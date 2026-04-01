package main

import (
	"fmt"
)

// Go doesn't have classes, but it has some other ways of
// creating class-ish structures
// The first one is the struct (structures)
// This creates a structure/blueprint with fields that
// groups these attributes together
// https://www.geeksforgeeks.org/go-language/class-and-object-in-golang/

// For example, we can create a Tea struct with different
// things we'd want to keep track of about the tea
type Tea struct {
	flavor string 
	brand string 
	steepTime float64
	price float64
}

// Inheritance (ish)
// Go doesn't have inheritance, but it does allow embedding of structs 
// which means you can put a struct as a field of a struct (kind of meta)
// Here, we create a TeaShop struct and embed the Tea struct within this
type TeaShop struct {
	name string 
	location string 
	isOpen bool
	Tea 
}

// Go also has interfaces, which are blueprints like structs 
// but they are for methods (or the behavior )
// a convention in Go is that interface names ends in "er"
// and matches the name of a struct to help you keep track
// of what goes with what (at least that's the theory)
type TeaShoper interface {
	Steep()
	ProcessOrder()
}

// You initialize the methods in functions outside of the interface
func (t TeaShop) Steep() {
	fmt.Println("The steep time for your", t.flavor, "is", t.steepTime, "minutes")
}

func (t TeaShop) ProcessOrder() {
	fmt.Println("Your", t.brand, t.flavor, "costs $", t.price)
	fmt.Println("Thanks for stopping by", t.name)
}


// To custom format printing out a struct as a string, 
// you declare a function that takes in the struct as an 
// argument, but specifically the pointer (so add the * before 
// the name of your struct). Then you use the String() method in Go.
// As usual, you define what your return type is (string obviously
// because what you're putting in all this hard work to return a 
// better formatted string after all.)
// Inside the function, you return your print statement. It can be handy 
// here to use the Sprintf() function to help format your string. 
// https://fullstackdojo.medium.com/the-python-str-method-and-the-go-string-method-a-comparison-75770c78c4d8
func (t *Tea) String() string {
	// https://www.geeksforgeeks.org/go-language/fmt-sprintf-function-in-golang-with-examples/
	// The Sprintf() function in Go is a way to format strings when printing 
	// You use %s for a string, %d for an integer, %f for a float 
	return fmt.Sprintf("%s's %s, Price $%f", t.brand, t.flavor, t.price)
}

func main() {
	fmt.Println("Hello Classes and Inheritance (ish)!")

	// Let's create an instance of our TeaShop struct 
	myTeaShop := TeaShop {
		name: "The Daily Steep",
		location: "Boston, MA",
		Tea: Tea {
			flavor: "London Fog",
			brand: "Tazo",
			steepTime: 4.5,
			price: 4.00,
		},
	}

	// Test out our steep and process order from our interface 
	myTeaShop.Steep()
	myTeaShop.ProcessOrder()

	// We can also create just an instance of our Tea struct
	myTeaOrder := Tea {
		flavor: "London Fog",
		brand: "Tazo",
		steepTime: 4.5,
		price: 4.00,
   	}

	// since we did the fancy formatting above with 
	// printing out the struct as a prettier string, 
	// printing it out will have it print what we specified 
	// above, right?
    fmt.Println("My tea order is: ", myTeaOrder)

	// Nope. There's another little detail to access our 
	// beautiful formatting. It's adding the & symbol 
	// so we can access the memory address (remember pass by reference?)
	// this is when that comes in handy.

	myActualTeaOrder := &Tea {
		flavor: "Chai Latte",
		brand: "Tazo",
		steepTime: 4.0,
		price: 4.75,
   	}

	fmt.Println("My next tea order is: ", myActualTeaOrder)

	// What if you want to change the flavor of your tea order? 
	// We use what's called dot notation, where you use your 
	// struct with your attribute following a .
	// Then you can reasign your attributes
	myActualTeaOrder.flavor = "Strawberry Matcha Latte"
	myActualTeaOrder.steepTime = 1.0
	myActualTeaOrder.price = 7.84

	// When we print out our order, it will be updated now
	fmt.Println("My updated tea order is: ", myActualTeaOrder)

}
