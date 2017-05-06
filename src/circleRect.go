package main

import (
"fmt"
"os"
"strconv"
)

type circle struct{
	r float64
}

type rectangle struct{
	l float64
	b float64
}

func (r rectangle) area() float64{
	return r.l*r.b
}

func (c circle) area() float64{
	var pi float64=3.141
	return pi*c.r*c.r
}

func (r rectangle) perimeter() float64{
	return 2*(r.l+r.b)
}

func (c circle) perimeter() float64{
	var pi float64=3.141
	return 2*pi*c.r
}

func main() {
	shape := os.Args[1]

	if shape == "circle" {
		radius := os.Args[2]
		rad,_ := strconv.ParseFloat(radius,64) 
		cir := circle{rad}

		fmt.Println("Area",cir.area())
		fmt.Println("Perimeter",cir.perimeter())
	}else{
		if shape == "rectangle" {
			wid := os.Args[2]
			hgt := os.Args[3]

			w,_ :=strconv.ParseFloat(wid,64) 
			h,_ :=strconv.ParseFloat(hgt,64) 

			rect := rectangle{w,h}

			fmt.Println("Area",rect.area())
			fmt.Println("Perimeter",rect.perimeter())
		}
	}
	
}