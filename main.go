package main

import (
	"fmt"
	"sync"
)

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func main() {
	counter := 0
	var wg sync.WaitGroup

	for i := range 100 {
		wg.Go(func() {
			fmt.Println(i)
			counter = counter + i
		})
	}

	wg.Wait()
	fmt.Println(counter)
}
