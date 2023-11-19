package main

import (
	"fmt"
	"net/http"
)

type anycode int

func (m anycode) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "print any code on ther screen")
}

func main() {
	var d anycode

	http.ListenAndServe(":8080", d)
}
