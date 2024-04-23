package main

import (
	"fmt"
	"sync"
)

// func main() {

// 	// var ch chan int - unbuffered channel
// 	ch := make(chan int, 3)
// 	go func(ch chan int) {
// 		ch <- 1
// 		ch <- 2
// 	}(ch)
// 	ch <- <-ch + <-ch
// 	fmt.Println(<-ch)
// }

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(x int) {
			fmt.Println(x)
			wg.Done()
		}(i)
	}
	wg.Wait()
}