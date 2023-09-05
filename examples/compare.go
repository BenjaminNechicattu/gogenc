package main

import (
	"fmt"

	"github.com/BenjaminNechicattu/gogenc/compare"
)

func main() {

	// int comparison
	var a int = 1
	var b int = 2
	maxInt := compare.Max(a, b)
	fmt.Println("maxInt : ", maxInt)
	minInt := compare.Min(a, b)
	fmt.Println("minInt : ", minInt)

	// float comparison
	var c float32 = 1.8
	var d float32 = 2.8
	maxFloat32 := compare.Max(c, d)
	fmt.Println("maxFloat32 : ", maxFloat32)
	minFloat32 := compare.Min(c, d)
	fmt.Println("minFloat32 : ", minFloat32)

	// string comparison
	var e string = "abc"
	var f string = "abd"
	maxString := compare.Max(e, f)
	fmt.Println("maxString : ", maxString)
	minString := compare.Min(e, f)
	fmt.Println("minString : ", minString)

}
