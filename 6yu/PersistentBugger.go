// Write a function, persistence, that takes in a positive parameter num and returns its multiplicative persistence, which is the number of times you must multiply the digits in num until you reach a single digit.

// 39 --> 3 (because 3*9 = 27, 2*7 = 14, 1*4 = 4 and 4 has only one digit)
// 999 --> 4 (because 9*9*9 = 729, 7*2*9 = 126, 1*2*6 = 12, and finally 1*2 = 2)
// 4 --> 0 (because 4 is already a one-digit number)
// 25 --> 2 (because 2*5 = 10, 1*0 = 0 and 0 has only one digit)

package main

import (
	"strconv"
)

func MultiplyDigits(n int) int {
  res := 1 
  nStr := strconv.Itoa(n)

  for _, char := range nStr {
    digit, err := strconv.Atoi(string(char))
    
    if err == nil {
      res = res * digit
    }
  }

  return res 
}

func Persistence(n int) int {
  counter := 0

  for n >= 10 {
    n = MultiplyDigits(n)
    counter += 1
  }

  return counter
}

func main() {
  println(Persistence(39))
  println(Persistence(999))
  println(Persistence(4))
  println(Persistence(25))
}
