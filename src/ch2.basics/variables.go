package main
import "fmt"

func main() {

    // var declares 1 or more variables.
    var a = "initial"
    fmt.Println(a)


    // declare multiple variables at once.
    var b, c int = 1, 2
    fmt.Println(b, c)


    // declare multiple variables like this
    var (
    	name     string = "Harry Potter"
    	age      int    =  14
    	location string = "Hogwarts School of Witchcraft and Wizardry"
    )
    fmt.Printf("%s is a %d years old wizard at %s", name, age, location)


    // infer the type of initialized variables.
    var d = true
    fmt.Println(d)

    // Variables declared without a corresponding initialization are zero-valued.
    // For example, the zero value for an int is 0.
    var e int
    fmt.Println(e)

    // The := syntax is shorthand for declaring and initializing a variable,
    // e.g. for var f string = "short" in this case.
    f := "short"
    fmt.Println(f)
}
