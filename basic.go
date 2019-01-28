package main

import "fmt"

// 函数外不可以直接用":"定义变量，每一行开头都要是"var"或者是"func"
var globalVar = 66

var (
	a = 1
	b = "66666"
	c = true
)


func variableDefault() {
	var i int
	var s string
	// fmt.Println(i, s)
	fmt.Printf("%d, %q\n", i, s)
}

func variableInitialVal() {
	var i, j int = 666, 233
	var s string = "mdzz"
	fmt.Printf("%d, %d, %q\n", i, j, s)
}

func variableTypeDeduction() {
	var i, j, a, b = 34, 45, "wertwert", true
	fmt.Println(i, j, a, b)
}

func variableSimplify() {
	// 第一次使用要用":"来定义，之后不要再次使用，否则相当于定义了两次
	i, j, a, b := 34, 45, "wertwert", true
	fmt.Println(i, j, a, b)
}

func consts() {
	// 常量尽量不要大写，因为 Go 语言中，变量首字母大写代表"public"
	const c = 10086
	const word = "Harry Potter"
	fmt.Println(c, word)
}

func enums() {

	/*

	`iota` identifier is used in const declarations to simplify
	definitions of incrementing numbers.

	The values of iota is reset to `0` whenever the reserved word
	`const` appears in the source (i.e. each const block) and increments
	by one after each ConstSpec e.g. each Line.

	*/

	const (
		c = iota
		cpp
		java
		python
		matlab
	)

	fmt.Println(c, cpp, java, python, matlab)


	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(b, kb, mb, gb, tb, pb)

}

func main() {
	fmt.Println("Hello, World!")

	variableDefault()
	variableInitialVal()
	variableTypeDeduction()
	variableSimplify()

	fmt.Println(a, b, c, globalVar)

	consts()
	enums()

}
