package main

import (
	"flag"
	"fmt"
)

var max int

func init() {
	flag.IntVar(&max, "upper", 100, "upper number")
	flag.Parse()
}

func main() {
	for i := 1; i <= max; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("fizzbuzz")
		case i%5 == 0:
			fmt.Println("buzz")
		case i%3 == 0:
			fmt.Println("fizz")
		default:
			fmt.Println(i)
		}
	}
}
