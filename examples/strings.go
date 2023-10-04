package main

import (
	"log"

	"github.com/BenjaminNechicattu/gogenc/strings"
)

func stringsExamples() {

	log.Println("------------------ strings -----------------")

	testString := "benjamin"
	expected := "nimajneb"
	resp := strings.Reverse(testString)
	log.Printf("testString: %v\n", testString)
	log.Printf("expected: %v\n", expected)
	log.Printf("resp: %v\n", resp)
}
