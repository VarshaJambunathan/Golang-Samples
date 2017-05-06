package main

import (
 "fmt"
 )
func main() {
	//Using Scan function
	var num int
	fmt.Println("Enter a number:")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("Even number :",num)
		return
	}
	fmt.Println("Odd number :",num)
}