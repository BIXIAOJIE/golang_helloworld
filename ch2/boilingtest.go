package main

import "fmt"

func write() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("%gf or %g c", f, c)
}
