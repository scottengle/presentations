package main

import (
	"encoding/json"
	"fmt"
)

// start main
func main() {
	a := sample{Prop: "hello world!"} // HL

	val, _ := json.Marshal(a)

	fmt.Println(string(val))
}

// end main

// show a highlight
type sample struct {
	Prop string `json:"property"` // HLprop
}

// end a highlight
