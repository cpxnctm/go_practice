package main

import (
	"fmt"
	"github.com/cpxnctm/go_practice/001-practice/tempconv"
)

func main(){
	c := (tempconv.CToF(20))
	f := (tempconv.FToC(20))

	fmt.Printf("20 degrees Celsius is %v in Fahrenheit", c)
	fmt.Printf("20 degrees Fahrenheit is %v in Celsius", f)
}