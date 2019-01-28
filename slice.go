package main

import "fmt"

func update(s []int) {
	s[0] = 2333
}

// slice 扩展
func reslice() {
	var arr = [...]int{1,2,3,4,5,6,7,8,9,0}

	s1 := arr[2:6]
	s2 := s1[3:5]

	// slice 可以向后扩展 但是不能向前扩展
	// s[i] 不可以超过 len(s), 向后扩展不可以超过底层数组 cap(s)
	fmt.Println("s1 = ", s1)
	fmt.Println("s2 = ", s2)

}


func makeSlice() {
	// 已知 slice 长度 但是内容未知
	s1 := make([]int, 16)     // len = 16
	s2 := make([]int, 16, 32) // len = 16, cap = 32
}


func main() {
	var arr = [...]int{1,2,3,4,5,6,7,8,9,0}

	// 以下 arr[*:*] 都是 slice 是对 array 的一个 "view"
	fmt.Println("arr[2:6] : ", arr[2:6])
	fmt.Println("arr[2:] : ", arr[2:])
	fmt.Println("arr[:6] : ", arr[:6])
	fmt.Println("arr[:] : ", arr[:])

	s1 := arr[3:5]
	fmt.Println(s1)
	update(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	reslice()
}
