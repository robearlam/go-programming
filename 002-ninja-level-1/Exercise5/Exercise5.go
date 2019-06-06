package main

import "fmt"

type cantona int

var x cantona

var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 1234

	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
}
