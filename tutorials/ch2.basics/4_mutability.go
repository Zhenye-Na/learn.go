/*

In Go, only constants are immutable.

However because arguments are passed by value, a function receiving a value
argument and mutating it, won’t mutate the original value.

*/

package main
import "fmt"

// struct
type Artist struct {
	Name, Genre string
	Songs       int
}


func newRelease(artist Artist) int {
	artist.Songs++
	return artist.Songs
}


func main() {
	me := Artist{Name: "Matt", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs\n", me.Name, me.Songs)
}
