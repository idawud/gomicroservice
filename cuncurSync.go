package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/idawud/gomicroservice/data"
)

var cache = make(map[int]data.Book)
var rnd  =  rand.New(rand.NewSource(time.Now().UnixNano()))

func RunConcurrency() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}

	for i:= 0; i < 6; i++ {
		id := rnd.Intn(6) + 1
		wg.Add(2)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex){
			if b, ok := queryCache(id, m); ok {
				fmt.Println("From Cache: " )
				fmt.Println(b.String())
			}
			wg.Done()
		}(id, wg, m)

		go func(id int,  wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryDatabase(id, m); ok {
				fmt.Println("From Database: ")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)

	}
	wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (data.Book, bool) {
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (data.Book, bool) {
	time.Sleep(100 * time.Microsecond) // artificial simulation
	for _, b := range data.Books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}
	return data.Book{}, false
}

