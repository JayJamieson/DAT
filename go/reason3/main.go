package main

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed README.md
var readme []byte

// http handlers need to implement handler method signature to
// work as request handlers.
// requirements are w http.ResponseWriter, r *http.Request as parameters
func readmeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(readme)
}
func main() {
	fmt.Printf("README.md: \n%v\n", string(readme))
	fmt.Println("Hello Fergus DAT!")

	// register a url patern and handler to default global ServerMux
	http.HandleFunc("/api/readme", readmeHandler)

	fmt.Println("go to http://localhost:8080/api/readme")

	// start a blocking server and bind to localhost:8080
	http.ListenAndServe(":8080", nil)
}
