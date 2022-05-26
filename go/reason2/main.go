package main

import (
	"embed"
	"fmt"
)

//go:embed README.md
var readme []byte

//go:embed README.md
var f embed.FS

func main() {
	fmt.Println("Hello Fergus DAT!")
	fmt.Printf("README.md embeded byte array: \n%v\n", string(readme))

	data1, _ := f.ReadFile("README.md")
	fmt.Printf("README.md embeded readonly filesystem: \n%v\n", string(data1))

	_, err := f.ReadFile("non_existing_file.md")

	fmt.Printf("non_existing_file error: \n%v\n", err.Error())
}
