package main

import (
	_ "embed"
	"fmt"
)

//go:embed README.md
var readme []byte

func main() {
	fmt.Println("Hello Fergus DAT!")
	fmt.Printf("README.md: \n%v\n", string(readme))
}
