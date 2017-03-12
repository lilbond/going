/*
A multi-line comment
*/
package main

import (
  "fmt"
  "strconv"
)

var (
  anInt int         //declare an int
  aString string    //declare a string
)

// A single line comment
func main()  {
  anInt = 1
  aString = "Hello, World"

  fmt.Println(anInt)
  fmt.Println(aString)

  // Programming is all about decision making, is not it?
  if anInt == 1 {
    fmt.Println(aString)
  }

  // A decision without a negative case is not so useful
  if anInt == 2 {
    fmt.Println(aString)
  } else {
    fmt.Println("Damn it was not true!!!")
  }

  // Ah! that was nice but how can I take more than one decisions
  if anInt == 2 {
    fmt.Println("It is 2 indeed")
  } else if anInt == 1 {
    fmt.Println("It is 1 indeed")
  } else {
    fmt.Println("I seriously have not idea, what it is")
  }
  /*
    Notice: else and else if have to be on the same line as the closing braces.
  */

  executor()
}

/*
Working with functions. We have already seen main, lets go beyond it.

A callable, reusable and self contained unit of code. Provides a
logical grouping and helps in organizing snippets to perform unit of work.

Disclaimer: I am NOT good at definitions and this one is purely self cooked :-)
*/

func greetAwesomePeople() {
  fmt.Println("Hello Awesome People. I am a dumb " +
      "function but teaches a very powerful thing." +
      "\nGuess what?") // Guess what?
}

// Same goes for me, guess guess :-)
func iAmBitSmarter(message string) {
  fmt.Println(message)
}

// And same goes for me
func iAmBitMoreSmarter(a, b int) int{
  return a + b
}

// will be called from within main
func executor() {
  greetAwesomePeople()
  iAmBitSmarter("Custom Message>> Sum of 10 and 2 is : " +
    strconv.Itoa(iAmBitMoreSmarter(10, 2)))
}

/*
Assignment
----------
Write the smartest calculator which:
- Works only with integer
- Handles add, subtract, mul and divide
client should be able to use your calculator like:
add(10,2), subtract(11, 3) etc.
*/
