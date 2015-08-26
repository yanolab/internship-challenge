package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
)

var pos int

func init() {
	flag.IntVar(&pos, "position", 0, "position")
	flag.Parse()
}

func primes() <-chan int {
	c := make(chan int, 1)
	cache := make([]int, 0)

	isPrime := func(n int) bool {
		for _, x := range cache[0 : sort.SearchInts(cache, int(math.Sqrt(float64(n))))+1] {
			if n%x == 0 {
				return false
			}
		}
		return true
	}

	go func() {
		c <- 2
		cache = append(cache, 2)
		for i := 3; ; i += 2 {
			if isPrime(i) {
				c <- i
				cache = append(cache, i)
			}
		}
	}()

	return c
}

func main() {
	counter := 0
	for prime := range primes() {
		counter++
		if counter == pos {
			fmt.Println(prime)
			break
		}
	}
}
