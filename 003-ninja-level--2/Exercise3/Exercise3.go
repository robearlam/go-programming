package main

import "fmt"

func main() {
	a := 42
	fmt.Printf("%d, %b, %#x", a, a, a)
	b := a << 1
	fmt.Printf("\n%d, %b, %#x", b, b, b)
}
