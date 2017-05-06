package main

import (
	"fmt"
)

func main() {              

	//Default values are zero
	var age int
	var tall bool
	var name, place string
	fmt.Printf("%#v, %#v, %#v, %#v\n", age, tall, name, place)

	//Automatic variable type assignment
	i := 42               // int
    s, b := "Baiju", true // string, bool
    fmt.Printf("%#v, %#v, %#v\n", i, s, b)
}


/* OUTPUT:
   
   0, false, "", ""

*/