/*

Go has pointers, but no pointer arithmetic. Struct fields can be accessed through
a struct pointer. The indirection through the pointer is transparent (you can
directly call fields and methods on a pointer).

Note that by default Go passes arguments by value (copying the arguments),
if you want to pass the arguments by reference, you need to pass pointers (or use
a structure using reference values like slices and maps.

To get the pointer of a value, use the `&` symbol in front of the value;
To dereference a pointer, use the `*` symbol.

Methods are often defined on pointers and not values (although they can be defined on both),
so you will often store a pointer in a variable as in the example below:

*/

package main

import "fmt"

func zeroval(ival int) {
  ival = 0
}

// zeroptr in contrast has an *int parameter, meaning that it takes an int pointer.
// The *iptr code in the function body then dereferences the pointer from its
// memory address to the current value at that address. Assigning a value to a
// dereferenced pointer changes the value at the referenced address.

func zeroptr(iptr *int) {
  *iptr = 0
}


func main() {
  i := 1
  fmt.Println("initial:", i)
  zeroval(i)
  fmt.Println("zeroval:", i)

  //The &i syntax gives the memory address of i, i.e. a pointer to i.
  zeroptr(&i)
  fmt.Println("zeroptr:", i)

  //Pointers can be printed too.
  fmt.Println("pointer:", &i)
}
