package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBinary(n int) string {
	result := ""

	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}


func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(convertToBinary(5))  // 101
	fmt.Println(convertToBinary(13)) // 1101

	printFile("file.txt")
}
