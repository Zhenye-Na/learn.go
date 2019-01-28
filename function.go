package main

import (
	"fmt"
)

func div(a, b int) (q, r int) {
	// q, r 已经定义好了，可以直接用来赋值，不需要自己定义
	q = a / b
	r = a % b
	return
}


//func apply(op func(int, int) int, a, b int) int {
//	p := reflect.ValueOf(op).Pointer()
//	opName := runtime.FuncForPC(p).Name()
//	fmt.Print("Calling function %s, with args: (%d, %d)\n", opName, a, b)
//}


func sumNumbers(numbers ...int) int {
	result := 0
	for i := range numbers {
		result += numbers[i]
	}
	return result
}


func main() {
	q, r := div(23, 3)
	fmt.Println(q, r)

	fmt.Println(sumNumbers(72,93,84,75,923,745,72,945,73,27))
}
