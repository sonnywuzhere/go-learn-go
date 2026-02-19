package main

import (
	"fmt"
)

// https://www.geeksforgeeks.org/go-language/class-and-object-in-golang/
// Go doesn't have classes, but it has some other ways of
// creating class-ish structures
// The first one is the struct (structures)
type Tea struct {
	flavor string 
	brand string 
	steepTime float64
	price float64
}

// Inheritance (ish)
// Go doesn't have inheritance, but it does allow embedding of structs 
type TeaShop struct {
	name string 
	location string 
	isOpen bool
	Tea 
}

// func (ts *TeaShop) String() string {
// 	return fmt.Println("")
// }
 
// a convention in Go is that interface names must end in "er"
type TS interface {
	Steep()
	ProcessOrder()
}

func (t TeaShop) Steep() {
	fmt.Println("The steep time for your", t.flavor, "is", t.steepTime, "minutes")
}

func (t TeaShop) ProcessOrder() {
	fmt.Println("Your", t.brand, t.flavor, "costs $", t.price)
	fmt.Println("Thanks for stopping by", t.name)
}

// https://www.geeksforgeeks.org/go-language/embedding-interfaces-in-golang/
type TeaShopDetailer interface {
	TeaShop
}

type Point struct {
	x int
	y int
}

func (p *Point) String() string {
	return fmt.Sprintf("Pointy y is %d", p.y)
}

// to custom format printing out a struct as a string, 
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
	return fmt.Sprintf("%s's %s is $%f", t.brand, t.flavor, t.price)
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

	// This prints out our TeaShop (note the nice formatting
	// thanks to our fancy function above)
	fmt.Println("Testing")
	fmt.Println(myTeaShop)

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

	myNextTeaOrder := &Tea {
		flavor: "Chai Latte",
		brand: "Tazo",
		steepTime: 4.0,
		price: 4.75,
   	}

	fmt.Println("My next tea order is: ", myNextTeaOrder)

}




// ● Create an object (a Person object like the one we created in lab is fine)
// ● Give the object class attributes (aka instance variables)
// ● Give the object functions
// ● Create an object that inherits from the first object (e.g., Student) with some attributes
// and functions of its own
// ● Test how to instantiate both types of objects, call their functions, and modify their
// variables.