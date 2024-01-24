package main

import (
	"strconv"
	"strings"
)

// Write a function that takes an integer as input, and returns the number of bits that are equal to one in the binary representation of that number. You can guarantee that input is non-negative.
// Example: The binary representation of 1234 is 10011010010, so the function should return 5 in this case

func CountBits(value uint) int {
  binaryValue := strconv.FormatInt(int64(value), 2)

  return strings.Count(binaryValue, "1")
}

func main() {
  println(CountBits(1234))
}
