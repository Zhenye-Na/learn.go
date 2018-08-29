/*

Attention: For loop is the only loop in Go, i.e. there is no `while ()` in Go

General:

For loop example

  sum := 0
  for i := 0; i < 10; i++ {
      sum += i
  }

For loop without pre/post statements

  sum := 1
  for ; sum < 1000; {
      sum += sum
  }

For loop as a while loop

  sum := 1
  for sum < 1000 {
      sum += sum
  }

*/


package main
import "fmt"


func main() {

  i := 1
  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }


  // A classic initial/condition/after for loop.
  for j := 7; j <= 9; j++ {
    fmt.Println(j)
  }


  // for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
  for {
    fmt.Println("loop")
    break
  }


  // You can also continue to the next iteration of the loop.
  for n := 0; n <= 5; n++ {
    if n % 2 == 0 {
      continue
    }
    fmt.Println(n)
  }
}

/*

Output:

1
2
3
7
8
9
loop
1
3
5

*/
