package main

import (
	"encoding/json"
	"fmt"
)

// start important OMIT
// start main
func main() {
	a := sample{Prop: "hello world!"}
	val, _ := json.Marshal(a)
	fmt.Println(string(val))
}

// end main

// show a struct OMIT
type sample struct {
	Prop string `json:"property"`
}

// end a struct OMIT
// end important OMIT
