/*

The expression T(v) converts the value v to the type T. Some numeric conversions:

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

Or, put more simply:

i := 42
f := float64(i)
u := uint(f)

*/


package main
import "fmt"

func main() {

  var gta int   = 5
  var heros int = 3

  hp := float32(gta * heros)
  mp := uint(hp)

  fmt.Printf("Which version fo GTA sieres do you like most? - GTA %d\n", gta)
  fmt.Printf("How many heroes in it? - %d\n", heros)

  // something un-interesting
  //hp -> 15.000000
  fmt.Printf("hp -> %f\n", hp)

  // mp -> %!f(uint=15)
  fmt.Printf("mp -> %f\n", mp)
}
