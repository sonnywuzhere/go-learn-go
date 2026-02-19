# Classes & Inheritance

1. Does your language support objects or something similar (e.g., structs)?

   Go doesn't have classes, but can support an object-oriented programming style.
   To achieve this, there are structs and interfaces.

   Structs are "typed collections of fields". This is to group some kind of data
   in the same place, similar to attributes in classes.

   Interfaces are "named collections of method signatures". This is a fancy way of saying
   that it groups together methods (think behavior) that haven't been defined yet as "blueprints",
   which are then implemented later using functions. It's like telling Go, "by the way,
   I want to group these things together, so save them for me, and I'll define them later."

   Naming conventions for structs follow the Go convention of starting with lower case with
   camelCase if needed.

   Interfaces are a little more particular in the Go convention. The names of Interfaces must
   end in "er". The name used before the "er" must match the name of the type that implements
   the matching "er" interface.

   In the example below, TeaShoper is the interface that goes with TeaShop.

   ```
   type TeaShop struct {
      name string
      location string
      isOpen bool
   }

   type TeaShoper interface {
      ProcessOrder()
   }

   func (t TeaShop) ProcessOrder() {
   	fmt.Println("Your", t.brand, t.flavor, "costs $", t.price)
   	fmt.Println("Thanks for stopping by", t.name)
   }

   ```

   However, this is just a naming convention in Go, so if you change TeaShoper to be TS,
   the code will still run.

   ```
   type TeaShop struct {
      name string
      location string
      isOpen bool
   }

   // If you break convention, the code will still run
   type TS interface {
      Steep()
      ProcessOrder()
   }

   func (t TeaShop) ProcessOrder() {
   	fmt.Println("Your", t.brand, t.flavor, "costs $", t.price)
   	fmt.Println("Thanks for stopping by", t.name)
   }
   ```

   However, if you want to be nice to the coders who will have to try to understand your
   code and follow convention in Go, stick to the "er" endings, even if they give your
   interfaces weird sounding names. It's ok. Nobody's paying you to be clever about naming
   your interfaces... that I know of, at least.

   https://www.geeksforgeeks.org/go-language/class-and-object-in-golang/
   https://gobyexample.com/interfaces
   https://gobyexample.com/structs
   https://go.dev/doc/faq#Is_Go_an_object-oriented_language

2. Does your language have standard methods for functions that serve a similar purpose
   across all objects? For example, toString() in Java and **str** in Python allow
   information about the objects to be printed. Are there similar functions in your
   language?

   The equivalent of this in Go is how to print out structs and interfaces in a better
   format.

   If you try to print out a struct, it doesn't look very nice and it just
   prints all the values. Pretty unhelpful.

   However, you can declare a String() function that returns a string that you can then
   format to your liking. There are a few details to keep in mind on this. When you
   declare the String() function, you need to make sure to have the parameter be the
   pointer of the struct (using the \* symbol) and when you initialize your struct, you
   need to use the & symbol. See the classes_inheritance.go file for an example.

   https://fullstackdojo.medium.com/the-python-str-method-and-the-go-string-method-a-comparison-75770c78c4d8

3. How does inheritance work (if it does)? Does your language support multiple
   inheritance?

   Since Go doesn't have classes, it doesn't technically have inheritance, However,
   you can embed structs within each other which is kind of similar to the concept
   of inheritance. See classes_inheritance.go for an example of embedding.

   Go also does not support traditional multiple inheritance as there is no type
   hierarchy.

   https://go.dev/doc/faq#inheritance

4. If there is inheritance, how does your language deal with overloading method names
   and resolving those calls?

   Go does not support overloading method names, and this was a deliberate design decision
   when designing the language. Go's philosophy is that it's "simpler without it."

   https://go.dev/doc/faq#overloading

5. Is there anything else that’s important to know about objects and inheritance in your
   language?

   Go does somewhat encapsulation through having public fields and methods be capitalized,
   and private fields/methods be lowercase. The complier enforces this, so it's not
   just a naming convention in Go. Capitalization = exported. Lowercase = unexported.
   This is because Go's goal was simplicity in handingling visibility.

   This is not just for the name of the struct but for the names of the fields in struct as well.

   For example:

   ```
   type Mug struct {
      Size int
      color string
   }
   ```

   The mug struct is public and can be exported because it's capitalized. However, you could only
   access the Size field because it is capitalized and the color field is not.

   https://medium.com/@oluwatobiojo2911/why-capitalisation-matters-in-go-structs-its-not-just-style-bf9005fa4347
   https://www.geeksforgeeks.org/go-language/object-oriented-programming-in-golang/
