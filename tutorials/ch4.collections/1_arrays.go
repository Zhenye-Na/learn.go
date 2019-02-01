/*

The type [n]T is an array of n values of type T.

The expression:

```
var a [10]int
```

declares a variable a as an array of ten integers.
By default an array is zero-valued, which for ints means 0s.

An array’s length is part of its type, so arrays cannot be resized.

*/



package main
import "fmt"

func main() {

  // initialize a two emply slots array
  var a [2]string
  a[0] = "Hello"
  a[1] = "World"
  fmt.Println(a[0], a[1])
  fmt.Println(a)


  // set the array entries as declaring the array
  b := [2]string{"hello", "world!"}
  fmt.Println(b)


  // use ellipsis to use an implicit length when you pass the values
  c := [...]string{"hello", ",", "world", "!"}
  fmt.Println(c)
  fmt.Println("lenth of array c is: ", len(c))

  // multidimensional array
  var d [2][3]int
  for i := 0; i < 2; i++ {
    for j := 0; j < 3; j++ {
      d[i][j] = (i + 1) * (j + 1)
    }
  }
  fmt.Println(d)
}

/*

Output:

Hello World
[Hello World]
[hello world!]
[hello , world !]
lenth of array c is:  4
[[1 2 3] [2 4 6]]

*/
