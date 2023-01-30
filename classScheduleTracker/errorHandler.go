package main

import (
	"fmt"
	"os"
)

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handleValidationError(err string) {
	fmt.Println(err)
	os.Exit(1)
}
