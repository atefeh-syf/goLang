package main

import "sync"

func main() {
	BadSimple()
	GoodSimple()
}

func BadSimple() {
	wg := sync.WaitGroup{}
	counter := 0

	wg.Add(1000000)
	for i := 0; i < 1_000_000; i++ {
		go func() {
			defer wg.Done()
			counter++
		}()
	}
	wg.Wait()
	println(counter)
}

func GoodSimple() {
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}
	counter := 0

	wg.Add(1000000)
	for i := 0; i < 1_000_000; i++ {
		go func() {
			defer wg.Done()
			mx.Lock()
			counter++
			mx.Unlock()
		}()
	}
	wg.Wait()
	println(counter)
}
