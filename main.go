package main

import (
	"fmt"
)

func printData(i interface{}) {
	switch i.(type) {
	case name:
		fmt.Printf("(%v, %T)\n", i, i)
	case age:
		fmt.Printf("(%v, %T)\n", i, i)
	case weight:
		fmt.Printf("(%v, %T)\n", i, i)
	default:
		fmt.Println("Data not found")
	}

}

type name string
type age int
type weight float32

func main() {
	var n name = "Sandra"
	var a age = 19
	var w weight = 55.5

	printData(n)
	printData(a)
	printData(w)
}
