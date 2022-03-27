package main

import (
	"fmt"
	"time"
)

func collatz(n int) int {
	count := 1
	m := n
	for m > 1 {
		if m%2 == 0 {
			m /= 2
		} else {
			m = m*3 + 1
		}
		count++
	}
	return count
}

func main() {
	now := time.Now()
	const limit int = 100000000
	max := 0
	key := 0
	for i := 2; i < limit; i++ {
		rc := collatz(i)
		if rc > max {
			max = rc
			key = i
		}
	}
	fmt.Printf("%d(%d)\n", key, max)
	fmt.Printf("%v ms\n", time.Since(now).Milliseconds())
}
