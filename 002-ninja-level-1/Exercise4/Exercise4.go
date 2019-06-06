package main

import "fmt"

type cantona int

var x cantona

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println(x)
}
