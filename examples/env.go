package main

import (
	"log"

	"github.com/BenjaminNechicattu/gogenc/env"
)

func envExamples() {
	log.Println("------------------ env -----------------")

	val, err := env.GetEnv("TEST_ENV", "default")
	if err != nil {
		log.Println("Error:", err)
	} else {
		log.Println("TEST_ENV:", val)
	}
	valInt, err := env.GetEnv("TEST_ENV_INT", 42)
	if err != nil {
		log.Println("Error:", err)
	} else {
		log.Println("TEST_ENV_INT:", valInt)
	}
	valBool, err := env.GetEnv("TEST_ENV_BOOL", false)
	if err != nil {
		log.Println("Error:", err)
	} else {
		log.Println("TEST_ENV_BOOL:", valBool)
	}
	valFloat, err := env.GetEnv("TEST_ENV_FLOAT", 3.14)
	if err != nil {
		log.Println("Error:", err)
	} else {
		log.Println("TEST_ENV_FLOAT:", valFloat)
	}
	valString, err := env.GetEnv("TEST_ENV_STRING", "default")
	if err != nil {
		log.Println("Error:", err)
	} else {
		log.Println("TEST_ENV_STRING:", valString)
	}

}
