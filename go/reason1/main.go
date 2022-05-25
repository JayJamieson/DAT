package main

import "fmt"

var environment string

func main() {
	fmt.Printf("Build commit: %v\n", environment)
	fmt.Println("Hello Fergus DAT!")
}
