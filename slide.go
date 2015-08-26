package main

import (
	"flag"
	"fmt"
)

var max int
var str string

func init() {
	flag.IntVar(&max, "size", 2, "window size")
	flag.StringVar(&str, "str", "", "param string")
	flag.Parse()
}

func slide(n int, str string) []string {
	buf := make([]string, 0)
	for i := 0; i < len(str)-(n-1); i++ {
		buf = append(buf, str[i:i+n])
	}
	return buf
}

func main() {
	for _, s := range slide(max, str) {
		fmt.Println(s)
	}
}
