package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"net/http"
	_ "net/http/pprof"
)

// Some function that does work
func hardWork(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Start: %v\n", time.Now())

	// Memory
	a := make([]string, 0, 0)
	for i := 0; i < 500000; i++ {
		a = append(a, "aaaa")
	}
	fmt.Println(strings.Join(a, ","))

	// Blocking
	time.Sleep(10 * time.Second)
	fmt.Printf("End: %v\n", time.Now())
}

func main() {
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go hardWork(&wg)
	time.Sleep(time.Second)
	go hardWork(&wg)
	time.Sleep(time.Second)
	go hardWork(&wg)
	wg.Wait()
}
