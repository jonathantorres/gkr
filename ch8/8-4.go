package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("8-4.go")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	f.Close()
}
