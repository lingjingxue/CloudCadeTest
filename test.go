package main

import (
	"fmt"
)

func TestProfanityInit() {
	ProfanityInit()
}

func TestProfanityFilter() {
	var old = "dsadas worda dsad"
	var new = ProfanityFilter(old)
	fmt.Println("old:", old, "new:", new)
}
