package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

func main() {
	RunConcurrency()
}

func goChannels()  {
	// Channels
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(2)

	go func( ch <-chan int, wg *sync.WaitGroup) {
		/*
			msg, ok := <-ch
			if ok {
				fmt.Println(msg)
			}
		*/
		for num := range ch {
			fmt.Println(num)
		}
		wg.Done()
	}(ch, wg)

	go func( ch chan<- int, wg *sync.WaitGroup) {
		for i := 1; i <= 10; i++ {
			ch <- int(math.Sqrt(float64(i + 40)) + math.Phi * math.Abs(math.Pi))  + rand.Intn(23)
		}
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}