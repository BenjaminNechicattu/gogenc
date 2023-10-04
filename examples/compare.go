package main

import (
	"log"

	"github.com/BenjaminNechicattu/gogenc/compare"
)

func compareExamples() {

	log.Println("------------------ compare -----------------")

	// int comparison
	var a int = 1
	var b int = 2
	maxInt := compare.Max(a, b)
	log.Println("maxInt : ", maxInt)
	minInt := compare.Min(a, b)
	log.Println("minInt : ", minInt)

	// float comparison
	var c float32 = 1.8
	var d float32 = 2.8
	maxFloat32 := compare.Max(c, d)
	log.Println("maxFloat32 : ", maxFloat32)
	minFloat32 := compare.Min(c, d)
	log.Println("minFloat32 : ", minFloat32)

	// string comparison
	var e string = "abc"
	var f string = "abd"
	maxString := compare.Max(e, f)
	log.Println("maxString : ", maxString)
	minString := compare.Min(e, f)
	log.Println("minString : ", minString)

}
