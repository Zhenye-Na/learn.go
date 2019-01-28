package main

import (
	"fmt"
	"io/ioutil"
)

func readFile() {
	const filename = "file.txt"

	if contents, err := ioutil.ReadFile(filename); err != nil {
		// err != nil ==> error is found
		fmt.Println(err)
	} else {
		fmt.Println(contents)
	}

}

func calGrade(score int) string {
	grade := ""

	switch {
	case score < 0 || score > 100:
		panic("Grade out of bounds!")
	case score < 60:
		grade = "D"
	case score < 70:
		grade = "C"
	case score < 88:
		grade = "B"
	case score <= 100:
		grade = "A"

	}

	return grade
}


func main() {
	readFile()

	fmt.Println(calGrade(66))
	fmt.Println(calGrade(777))
	fmt.Println(calGrade(88))
	fmt.Println(calGrade(99))
	fmt.Println(calGrade(9))
}
