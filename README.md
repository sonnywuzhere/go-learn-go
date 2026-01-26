# Go Learn Go
If you've ever wanted to learn Go, now's your time to go for it! If you want the official tutorial from the Go documentation, 
head on over to https://go.dev/learn/

But if you're looking for a fun alternative, then you're in the right place. 

## Step 0

In order to dive into the exciting world of Go, we first need to do all the boring (but important) stuff...

### Installation 
Go to https://go.dev/doc/install and click the big blue button that says "Download". Make sure you have the corrent operating 
system selected (Linux, Mac, Windows), otherwise you'll have a sad day. 

### Programming Environment 
Pick your favorite IDE (integrated development environment), since Go doesn't need anything super fancy. 

Top options: 
- VS Code (free) https://code.visualstudio.com/
- GoLand (paid) https://www.jetbrains.com/go/
- Vim (free) https://www.vim.org/

And if you need a coding conditional statment to aid in your decision making, look no further:

```
if (cheap && easy) {
	pick(VS Code)
	// https://code.visualstudio.com/
} else if (!cheap && fancy) {
	pick(GoLand)
	// https://www.jetbrains.com/go/
} else if (nerdy && uglyWebsite) {
	pick(Vim)
	// https://www.vim.org/
} else {
  try {
    askAI()
  } catch {
    askGoogle()
  }
}
```

### Hello, world! program (a classic) 

Now to get to the good stuff: actually running code! 

Once you have Go downloaded successfully, create a file with a .go extension. Open it in your IDE of choice. 
It's time to write your first Hello, world! program in Go!

(the following code is from the Go documentation: https://go.dev/doc/tutorial/getting-started)

package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	fmt.Println("This program is written in Go")
}


### Running programs 

Great, you have your Go program sitting on your computer doing absolutely nothing. 
Let's change that. 

Open a terminal (either in your IDE or separately) and enter the following: 

go run fileName.go

This should print 

Hello, world!
This program is written in Go

If you get an error, retrace the steps to double check your work and try again. 


### Writing comments 
https://www.w3schools.com/go/go_comments.php

Comments are written in Go with two forward slashes for single line comments

// this would be a comment in Go 

/* this is also a comment in Go 
and it is a multi-line comment */
