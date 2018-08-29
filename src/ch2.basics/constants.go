/*
Go supports constants of character, string, boolean, and numeric values.

Constants can only be character, string, boolean, or numeric values and
cannot be declared using the := syntax. An untyped constant takes the type
needed by its context.
*/

package main

import (
    "fmt"
    "math"
)

const Pi = 3.14
const (
        StatusOK                   = 200
        StatusCreated              = 201
        StatusAccepted             = 202
        StatusNonAuthoritativeInfo = 203
        StatusNoContent            = 204
        StatusResetContent         = 205
        StatusPartialContent       = 206
)

const s string = "constant"

func main() {
    fmt.Println(s)

    // A const statement can appear anywhere a var statement can.
    const n = 500000000


    // Constant expressions perform arithmetic with arbitrary precision.
    const d = 3e20 / n
    fmt.Println(d)


    // A numeric constant has no type until it’s given one, such as by an explicit cast.
    fmt.Println(int64(d))


    /* A number can be given a type by using it in a context that requires one,
       such as a variable assignment or function call. For example,
       here math.Sin expects a float64. */
    fmt.Println(math.Sin(n))
}
