package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num := os.Args[1]

	val, err := strconv.Atoi(num)

	if err != nil {
		fmt.Println("Can't Convert!!")
		return
	}

	if val%2 == 0 {
		fmt.Println("Even number :", num)
		return
	}
	fmt.Println("Odd number :", num)
}
