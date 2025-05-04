package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)

	i1 := i / 100
	i2 := (i / 10) % 10
	i3 := i % 10

	fmt.Println(i1 + i2 + i3)
}
