package test

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func testmain() {
	var pair []string
	pair = append(pair, "BTC-USD", "ETH-USD", "ADA-USD", "XRP-USD", "SOL-USD", "AUDIO-USD")

	var wg sync.WaitGroup
	min := 1
	max := len(pair)

	wg.Add(max)
	for _, row := range pair {
		randNumber := rand.Intn(max-min) + min
		go DoSomething(row, randNumber, &wg)

	}

	wg.Wait()
	fmt.Println("DONDE!")
}

func DoSomething(pair string, t int, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		time.Sleep(time.Duration(t) * time.Second)
		fmt.Println("Pair:", pair, "Sleep for ", t, " seconds!")
	}()
}
