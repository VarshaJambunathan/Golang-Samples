package main

import (
"fmt"
)
func main() {
	var sum int=0

	for i:=0;i<=50;i++ {
		/*
		Without logical operators		
		if i %3 ==0 {

			fmt.Println(i)
			sum = sum + i;
		}else{
			if i%5==0{
				fmt.Println(i)
				sum = sum + i;
			}
		}
		*/

		//Using logical operators
		if (i%3 ==0) || (i%5==0) {
			//fmt.Println(i)
			sum = sum + i
		}
		
	}	
	fmt.Println(sum)

}