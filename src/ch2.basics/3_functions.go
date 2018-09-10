package main
import "fmt"

// a function that takes two ints and returns their sum as an int
func add(x int, y int) int {
  // Go won't automatically return the value of the last expression.
  return x + y
}


// When you have multiple consecutive parameters of the same type, you may omit
// the type name for the like-typed parameters up to the final parameter that
// declares the type.
func mul(a, b int) int {
  return a * b
}

func main() {

  fmt.Println("42 + 13 = ", add(42, 13))

  result := mul(4, 5)
  fmt.Println("4 x 5 = ", result)

}
