package main
import "fmt"

// inline if-else
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
  // basic if-else statement
  if 7 % 2 == 0 {
    fmt.Println("7 is even")
  } else {
    fmt.Println("7 is odd")
  }


  // You can have an if statement without an else.
  if 8 % 4 == 0 {
    fmt.Println("8 is divisible by 4")
  }


  // A statement can precede conditionals; any variables declared in this statement are available in all branches.
  if num := 9; num < 0 {
    fmt.Println(num, "is negative")
  } else if num < 10 {
    fmt.Println(num, "has 1 digit")
  } else {
    fmt.Println(num, "has multiple digits")
  }


  // inline if-else
  fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
