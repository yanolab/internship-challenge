package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var file string

func init() {
	flag.StringVar(&file, "file", "", "digits file")
	flag.Parse()
}

func slide(n int, str string) []string {
	buf := make([]string, 0)
	for i := 0; i < len(str)-(n-1); i++ {
		buf = append(buf, str[i:i+n])
	}
	return buf
}

func productString(str string) int {
	n := 1
	for _, ch := range str {
		v, err := strconv.Atoi(string(ch))
		if err != nil {
			panic(err)
		}
		n = n * v
	}
	return n
}

func main() {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	max := 0
	for _, s := range slide(5, strings.TrimRight(string(buf), "\n")) {
		i := productString(s)
		if max < i {
			max = i
		}
	}
	fmt.Println(max)
}
