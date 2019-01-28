package main

import "fmt"

func traverse() {
	var arr = [...]int{9,23,8,47,5,2,3,4,5}

	// 类似于 Python 中的 enumerate()
	for i, v := range arr {
		fmt.Println(i, v)
	}

}


func main() {
	// var arr1 [5]int
	// var arr2 = [3]int {1,2,3}
	// var arr3 = [...]int{2,3,5,2,3,4,5}
	//
	// var grid [2][3]int

	traverse()
}
