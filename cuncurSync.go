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

	cacheCh := make(chan data.Book)
	dbCh := make(chan data.Book)

	for i:= 0; i < 34; i++ {
		id := rnd.Intn(6) + 1
		wg.Add(2)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- data.Book){
			if b, ok := queryCache(id, m); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, m,  cacheCh)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- data.Book) {
			if b, ok := queryDatabase(id); ok {
				m.Lock()
				cache[id] = b
				m.Unlock()
				ch <- b
			}
			wg.Done()
		}(id, wg, m, dbCh)

		// create one goroutine per query to handle response
		go func(cacheCh, dbCh <-chan data.Book) {
			select {
				case b := <-cacheCh:
					fmt.Println("from cache")
					fmt.Println(b)
					<-dbCh
				case b := <-dbCh:
					fmt.Println("from database")
					fmt.Println(b)
			}
		}( cacheCh, dbCh)
		time.Sleep(120*time.Microsecond)
	}
	wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (data.Book, bool) {
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int) (data.Book, bool) {
	time.Sleep(100 * time.Microsecond) // artificial simulation
	for _, b := range data.Books {
		if b.ID == id {
			return b, true
		}
	}
	return data.Book{}, false
}

