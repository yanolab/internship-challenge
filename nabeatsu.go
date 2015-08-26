package main

import (
	"flag"
	"fmt"
	"strconv"
)

var max int

func init() {
	flag.IntVar(&max, "upper", 100, "upper number")
	flag.Parse()
}

func isNabeatsuNumber(n int) bool {
	if n%3 == 0 {
		return true
	}
	for _, ch := range strconv.Itoa(n) {
		if ch == '3' {
			return true
		}
	}
	return false
}

func main() {
	for i := 1; i <= max; i++ {
		if isNabeatsuNumber(i) {
			fmt.Println("Aho")
		} else {
			fmt.Println(i)
		}
	}
}
