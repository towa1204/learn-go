package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("q1.txt")
	if err != nil {
		panic(err)
	}
	name := "aiueo"
	num := 100
	pi := 3.14
	fmt.Fprintf(file, "name = %s\n", name)
	fmt.Fprintf(file, "num = %d\n", num)
	fmt.Fprintf(file, "pi = %fq\n", pi)
}
